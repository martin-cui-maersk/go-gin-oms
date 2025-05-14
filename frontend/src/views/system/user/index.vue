<template>
  <PageWrapper dense contentFullHeight fixedHeight contentClass="flex">
    <BasicTable @register="registerTable" :searchInfo="searchInfo">
      <template #toolbar>
        <Authority :value="'AddUser'">
          <a-button type="primary" @click="handleCreate">{{ t('system.user.addUser') }}</a-button>
        </Authority>
        <Authority :value="'ExportUser'">
          <a-button type="primary" @click="handleExport">{{ t('system.user.exportUser') }}</a-button>
        </Authority>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'clarity:info-standard-line',
                tooltip: '查看用户详情',
                onClick: handleView.bind(null, record),
                auth: 'ShowUserDetail'
              },
              {
                icon: 'clarity:note-edit-line',
                tooltip: '编辑用户资料',
                onClick: handleEdit.bind(null, record),
                auth: 'UpdateUser'
              },
              {
                icon: 'ant-design:delete-outlined',
                color: 'error',
                auth: 'DeleteUser',
                tooltip: '删除此账号',
                popConfirm: {
                  title: '是否确认删除',
                  placement: 'left',
                  confirm: handleDelete.bind(null, record),
                },
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <UserModal @register="registerModal" @success="handleSuccess" />
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { Authority } from '@/components/Authority';
  import { reactive } from 'vue';

  import { BasicTable, useTable, TableAction } from '@/components/Table';
  import { getUserList } from '@/api/system/user';
  import { PageWrapper } from '@/components/Page';

  import { useModal } from '@/components/Modal';
  import UserModal from './UserModal.vue';

  import { columns, searchFormSchema } from './user.data';
  import { useGo } from '@/hooks/web/usePage';
  import { useI18n } from '@/hooks/web/useI18n';

  const { t } = useI18n();

  defineOptions({ name: 'AccountManagement' });

  const go = useGo();
  const [registerModal, { openModal }] = useModal();
  const searchInfo = reactive<Recordable>({});
  const [registerTable, { reload, updateTableDataRecord, getSearchInfo }] = useTable({
    title: t('system.user.userList'),
    api: getUserList,
    rowKey: 'id',
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
      autoSubmitOnEnter: true,
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
      width: 120,
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
      console.log('handleSuccess', result, values);
    } else {
      reload();
    }
  }

  function handleView(record: Recordable) {
    go('/system/user/user-detail/' + record.id);
  }
</script>
