<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="registerDrawer"
    showFooter
    :title="getTitle"
    width="50%"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm" />
  </BasicDrawer>
</template>
<script lang="ts" setup>
  import { ref, computed, unref } from 'vue';
  import { BasicForm, useForm } from '@/components/Form';
  import { formSchema } from './menu.data';
  import { BasicDrawer, useDrawerInner } from '@/components/Drawer';

  import { getMenuList, addMenu, updateMenu } from '@/api/system/menu';
  import { useI18n } from '@/hooks/web/useI18n';

  const { t } = useI18n();
  defineOptions({ name: 'MenuDrawer' });

  const emit = defineEmits(['success', 'register']);

  const isUpdate = ref(true);

  const [registerForm, { resetFields, setFieldsValue, updateSchema, validate }] = useForm({
    labelWidth: 120,
    schemas: formSchema,
    showActionButtonGroup: false,
    baseColProps: { lg: 12, md: 24 },
  });

  const [registerDrawer, { setDrawerProps, closeDrawer, changeOkLoading }] = useDrawerInner(async (data) => {
    resetFields();
    setDrawerProps({ loading: false, loadingTip: 'Loading...' });
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      setFieldsValue({
        ...data.record,
      });
    }
    const treeData = await getMenuList();
    updateSchema({
      field: 'parentId',
      componentProps: { treeData },
    });
  });

  const getTitle = computed(() =>
    !unref(isUpdate) ? t('system.menu.addMenu') : t('system.menu.editMenu'),
  );

  async function handleSubmit() {
    try {
      const values = await validate();
      setDrawerProps({ loading: true });
      // TODO custom api
      console.log(values);
      if (!unref(isUpdate)) {
        addMenu(values).then(async (res) => {
          await changeOkLoading(false);
          await setDrawerProps({ loading: false });
          if (res) {
            closeDrawer();
            emit('success');
          }
        });
      } else {
        updateMenu(values).then(async (res) => {
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
