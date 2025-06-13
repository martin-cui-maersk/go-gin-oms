<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="registerDrawer"
    showFooter
    :title="getTitle"
    width="50%"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm">
      <template #menu="{ model, field }">
        <BasicTree
          v-model:value="model[field]"
          :treeData="treeData"
          :fieldNames="{ title: 'menuName', key: 'id' }"
          checkable
          :selectable="false"
          search
          toolbar
          :title="t('system.role.assignMenu')"
          v-model:checkedKeys="checkedKeys"
        />
      </template>
    </BasicForm>
  </BasicDrawer>
</template>
<script lang="ts" setup>
  import { ref, computed, unref } from 'vue';
  import { BasicForm, useForm } from '@/components/Form';
  import { formSchema } from './role.data';
  import { BasicDrawer, useDrawerInner } from '@/components/Drawer';
  import { BasicTree, TreeItem } from '@/components/Tree';

  import { addRole, updateRole } from '@/api/system/role';
  import { getActiveMenuList } from '@/api/system/menu';
  import { useI18n } from '@/hooks/web/useI18n';

  const { t } = useI18n();
  const checkedKeys = ref([]);
  const emit = defineEmits(['success', 'register']);
  const isUpdate = ref(true);
  const treeData = ref<TreeItem[]>([]);

  const [registerForm, { resetFields, setFieldsValue, validate }] = useForm({
    labelWidth: 90,
    baseColProps: { span: 24 },
    schemas: formSchema,
    showActionButtonGroup: false,
  });

  const [registerDrawer, { setDrawerProps, closeDrawer, changeOkLoading }] = useDrawerInner(async (data) => {
    resetFields();
    setDrawerProps({ loading: false, loadingTip: 'Loading...' });
    // 需要在setFieldsValue之前先填充treeData，否则Tree组件可能会报key not exist警告
    if (unref(treeData).length === 0) {
      treeData.value = (await getActiveMenuList()) as any as TreeItem[];
    }
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      setFieldsValue({
        ...data.record,
      });
    }
  });

  const getTitle = computed(() =>
    !unref(isUpdate) ? t('system.role.addRole') : t('system.role.editRole'),
  );

  async function handleSubmit() {
    try {
      const values = await validate();
      console.log('role checkedKeys', checkedKeys);
      setDrawerProps({ loading: true });
      // TODO custom api
      console.log(values);
      if (!unref(isUpdate)) {
        addRole(values).then(async (res) => {
          await changeOkLoading(false);
          await setDrawerProps({ loading: false });
          if (res) {
            closeDrawer();
            emit('success');
          }
        });
      } else {
        updateRole(values).then(async (res) => {
          await changeOkLoading(false);
          await setDrawerProps({ loading: false });
          if (res) {
            closeDrawer();
            emit('success');
          }
        });
      }
    } finally {
      // setDrawerProps({ loading: false });
    }
  }
</script>
