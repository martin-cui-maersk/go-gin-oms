<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="getTitle" @ok="handleSubmit">
    <BasicForm @register="registerForm">
      <template #store="{ model, field }">
        <BasicTree
          v-model:value="model[field]"
          :treeData="treeData"
          :fieldNames="{ title: 'name', key: 'id' }"
          checkable
          :selectable="false"
          search
          toolbar
          :title="t('system.user.assignStore')"
          v-model:checkedKeys="checkedKeys"
        />
      </template>
    </BasicForm>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, computed, unref } from 'vue';
  import { BasicModal, useModalInner } from '@/components/Modal';
  import { BasicForm, useForm } from '@/components/Form';
  import { accountFormSchema } from './user.data';
  import { addUser, updateUser } from '@/api/system/user';
  import { BasicTree, TreeItem } from '@/components/Tree';
  import { storeTree } from '@/api/system/store';
  import { useI18n } from '@/hooks/web/useI18n';

  const { t } = useI18n();
  defineOptions({ name: 'UserModal' });

  const emit = defineEmits(['success', 'register']);

  const isUpdate = ref(true);
  const rowId = ref('');
  const checkedKeys = ref([]);
  const treeData = ref<TreeItem[]>([]);

  const [registerForm, { setFieldsValue, updateSchema, resetFields, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 24 },
    schemas: accountFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 23,
    },
  });

  const [registerModal, { setModalProps, closeModal, changeOkLoading }] = useModalInner(
    async (data) => {
      resetFields();
      setModalProps({ loading: false, loadingTip: 'Loading...' });
      isUpdate.value = !!data?.isUpdate;

      // 需要在setFieldsValue之前先填充treeData，否则Tree组件可能会报key not exist警告
      if (unref(treeData).length === 0) {
        treeData.value = (await storeTree()) as any as TreeItem[];
      }

      if (unref(isUpdate)) {
        rowId.value = data.record.id;
        setFieldsValue({
          ...data.record,
        });
      }

      updateSchema([
        {
          field: 'pwd',
          show: !unref(isUpdate),
        },
      ]);
    },
  );

  const getTitle = computed(() => (!unref(isUpdate) ? t('system.user.addUser') : t('system.user.editUser') ));

  async function handleSubmit() {
    try {
      const values = await validate();
      await changeOkLoading(true);
      await setModalProps({ loading: true });
      // TODO custom api
      console.log(values);
      if (unref(isUpdate)) {
        updateUser(values).then(async (res) => {
          console.log('updateUser', res);
          await changeOkLoading(false);
          await setModalProps({ loading: false });
          if (res) {
            closeModal();
            emit('success', { isUpdate: unref(isUpdate), values: { ...values, id: rowId.value } });
          }
        });
      } else {
        addUser(values).then(async (res) => {
          console.log('updateUser', res);
          await changeOkLoading(false);
          await setModalProps({ loading: false });
          if (res) {
            closeModal();
            emit('success', { isUpdate: unref(isUpdate), values: { ...values, id: rowId.value } });
          }
        });
      }
    } finally {
      // setModalProps({ loading: false });
    }
  }
</script>
