<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="getTitle"
    @ok="handleSubmit"
    width="50%"
  >
    <BasicForm @register="registerForm" />
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, computed, unref } from 'vue';
  import { BasicModal, useModalInner } from '@/components/Modal';
  import { BasicForm, useForm } from '@/components/Form';
  import { storeFormSchema } from './merchant.data';
  import { addMerchant, updateMerchant } from '@/api/system/merchant';
  import { useI18n } from '@/hooks/web/useI18n';

  const { t } = useI18n();

  defineOptions({ name: 'MerchantModal' });

  const emit = defineEmits(['success', 'register']);

  const isUpdate = ref(true);
  const rowId = ref('');

  const [registerForm, { setFieldsValue, updateSchema, resetFields, validate }] = useForm({
    labelWidth: 150,
    baseColProps: { span: 24 },
    schemas: storeFormSchema,
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

  const getTitle = computed(() =>
    !unref(isUpdate) ? t('system.merchant.addMerchant') : t('system.merchant.editMerchant'),
  );

  async function handleSubmit() {
    try {
      const values = await validate();
      await changeOkLoading(true);
      await setModalProps({ loading: true });
      // TODO custom api
      console.log(values);
      if (unref(isUpdate) && values.id !== undefined && values.id > 0) {
        updateMerchant(values).then(async (res) => {
          console.log('updateMerchant', res);
          await changeOkLoading(false);
          await setModalProps({ loading: false });
          if (res) {
            closeModal();
            emit('success', { isUpdate: unref(isUpdate), values: { ...values, id: rowId.value } });
          }
        });
      } else {
        addMerchant(values).then(async (res) => {
          await changeOkLoading(false);
          await setModalProps({ loading: false });
          if (res) {
            closeModal();
            emit('success', { isUpdate: unref(isUpdate), values: { ...values, id: rowId.value } });
          }
        });
      }
    } finally {
      // await changeOkLoading(false);
      // await setModalProps({ loading: false });
    }
  }
</script>
