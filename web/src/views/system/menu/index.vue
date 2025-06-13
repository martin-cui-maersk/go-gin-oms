<template>
  <div>
    <BasicTable @register="registerTable" @fetch-success="onFetchSuccess">
      <template #toolbar>
        <Authority :value="'AddMenu'">
          <a-button type="primary" @click="handleCreate">
            {{ t('system.menu.addMenu') }}
          </a-button>
        </Authority>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'clarity:note-edit-line',
                onClick: handleEdit.bind(null, record),
                auth: 'UpdateMenu', // 更新菜单按钮
              },
              {
                icon: 'ant-design:delete-outlined',
                color: 'error',
                auth: 'DeleteMenu', // 删除菜单按钮
                popConfirm: {
                  title: 'Confirm delete?',
                  placement: 'left',
                  confirm: handleDelete.bind(null, record),
                },
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <MenuDrawer @register="registerDrawer" @success="handleSuccess" />
  </div>
</template>
<script lang="ts" setup>
  import { Authority } from '@/components/Authority';
  import { useI18n } from '@/hooks/web/useI18n';
  import { nextTick } from 'vue';

  import { BasicTable, useTable, TableAction } from '@/components/Table';
  import { getMenuList, deleteMenu } from '@/api/system/menu';

  import { useDrawer } from '@/components/Drawer';
  import MenuDrawer from './MenuDrawer.vue';

  import { columns, searchFormSchema } from './menu.data';
  import { useMessage } from '@/hooks/web/useMessage';

  const { t } = useI18n();

  const { createMessage, createInfoModal, createErrorModal, createSuccessModal } = useMessage();

  defineOptions({ name: 'MenuManagement' });

  const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerTable, { reload, expandAll, setLoading }] = useTable({
    title: t('system.menu.menuList'),
    api: getMenuList,
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
    },
    isTreeTable: true,
    pagination: false,
    striped: true,
    useSearchForm: false, // 关闭搜索栏
    showTableSetting: true,
    bordered: true,
    showIndexColumn: false,
    canResize: false,
    actionColumn: {
      width: 100,
      title: t('global.operation'),
      dataIndex: 'action',
      // slots: { customRender: 'action' },
      // fixed: undefined,
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
    console.log('handleDelete Menu', record.id);
    setLoading(true);
    deleteMenu(record.id).then(() => {
      setLoading(false);
      setTimeout(() => {
        reload();
      }, 1000);
    });
  }

  function handleSuccess() {
    reload();
  }

  function onFetchSuccess() {
    // 演示默认展开所有表项
    nextTick(expandAll);
  }
</script>
