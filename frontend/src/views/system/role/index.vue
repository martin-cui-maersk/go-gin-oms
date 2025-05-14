<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <Authority :value="'AddRole'">
          <a-button type="primary" @click="handleCreate"> {{ t('system.role.addRole') }} </a-button>
        </Authority>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'clarity:note-edit-line',
                onClick: handleEdit.bind(null, record),
                auth: 'UpdateRole',
              },
              {
                icon: 'ant-design:delete-outlined',
                color: 'error',
                auth: 'DeleteRole',
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
    <RoleDrawer @register="registerDrawer" @success="handleSuccess" />
  </div>
</template>
<script lang="ts" setup>
  import { Authority } from '@/components/Authority';
  import { BasicTable, useTable, TableAction } from '@/components/Table';
  import { getRoleList } from '@/api/system/role';

  import { useDrawer } from '@/components/Drawer';
  import RoleDrawer from './RoleDrawer.vue';

  import { columns, searchFormSchema } from './role.data';
  import { useI18n } from '@/hooks/web/useI18n';

  const { t } = useI18n();
  defineOptions({ name: 'RoleManagement' });

  const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerTable, { reload }] = useTable({
    title: t('system.role.roleList'),
    api: getRoleList,
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
    },
    useSearchForm: true,
    showTableSetting: true,
    bordered: true,
    showIndexColumn: false,
    actionColumn: {
      width: 80,
      title: t('global.operation'),
      dataIndex: 'action',
      // slots: { customRender: 'action' },
      fixed: 'right',
    },
  });

  function handleCreate() {
    openDrawer(true, {
      isUpdate: false,
    });
  }

  function handleEdit(record: Recordable) {
    openDrawer(true, {
      record,
      isUpdate: true,
    });
  }

  function handleDelete(record: Recordable) {
    console.log(record);
  }

  function handleSuccess() {
    reload();
  }
</script>
