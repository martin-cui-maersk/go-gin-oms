<template>
  <PageWrapper dense contentFullHeight fixedHeight contentClass="flex">
    <BasicTable @register="registerTable" :searchInfo="searchInfo">
      <template #toolbar>
        <Authority :value="'AddMerchant'">
          <a-button type="primary" @click="handleCreate">{{ t('system.merchant.addMerchant') }}</a-button>
        </Authority>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'clarity:note-edit-line',
                tooltip: t('system.merchant.editMerchant'),
                onClick: handleEdit.bind(null, record),
                auth: 'UpdateMerchant',
              }
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <MerchantModal @register="registerModal" @success="handleSuccess" />
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { reactive } from 'vue';
  import { Authority } from '@/components/Authority';
  import { BasicTable, useTable, TableAction } from '@/components/Table';
  import { getMerchantList } from '@/api/system/merchant';
  import { PageWrapper } from '@/components/Page';

  import { useModal } from '@/components/Modal';
  import MerchantModal from './MerchantModal.vue';

  import { columns, searchFormSchema } from './merchant.data';
  import { useGo } from '@/hooks/web/usePage';
  import { useI18n } from '@/hooks/web/useI18n';

  const { t } = useI18n();
  defineOptions({ name: 'MerchantManagement' });

  const go = useGo();
  const [registerModal, { openModal }] = useModal();
  const searchInfo = reactive<Recordable>({});
  const [registerTable, { reload, updateTableDataRecord, getSearchInfo }] = useTable({
    title: t('system.merchant.merchantList'),
    api: getMerchantList,
    rowKey: 'id',
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
      autoSubmitOnEnter: true,
      showAdvancedButton: true
    },
    size: 'small',
    striped: true,
    // ellipsis: false,
    useSearchForm: true,
    showTableSetting: true,
    bordered: true,
    showIndexColumn: false,
    isTreeTable: false,
    canResize: true,
    handleSearchInfoFn(info) {
      console.log('handleSearchInfoFn', info);
      return info;
    },
    actionColumn: {
      width: 100,
      title: t('global.operation'),
      dataIndex: 'action',
      // slots: { customRender: 'action' },
    },
  });

  function handleCreate() {
    openModal(true, {
      isUpdate: false,
    });
  }

  function handleEdit(record: Recordable) {
    console.log(record);
    openModal(true, {
      record,
      isUpdate: true,
    });
  }

  function handleDelete(record: Recordable) {
    console.log(record);
  }

  function handleExport() {
    console.log(getSearchInfo());
  }

  function handleSuccess({ isUpdate, values }) {
    if (isUpdate) {
      // 演示不刷新表格直接更新内部数据。
      // 注意：updateTableDataRecord要求表格的rowKey属性为string并且存在于每一行的record的keys中
      const result = updateTableDataRecord(values.id, values);
      console.log(result);
    } else {
      reload();
    }
  }
</script>
