import { BasicColumn, FormSchema } from '@/components/Table';
import { h } from 'vue';
import { Switch } from 'ant-design-vue';
import { setRoleStatus } from '@/api/system/role';
import { useMessage } from '@/hooks/web/useMessage';
import { useI18n } from '@/hooks/web/useI18n';

const { t } = useI18n();

type CheckedType = boolean | string | number;
export const columns: BasicColumn[] = [
  {
    title: t('system.role.roleName'),
    dataIndex: 'roleName',
    width: 200,
  },
  {
    title: t('system.role.roleCode'),
    dataIndex: 'roleCode',
    width: 150,
  },
  {
    title: t('global.status'),
    dataIndex: 'status',
    width: 120,
    auth: 'SetRoleStatus', // 同时根据权限控制是否显示
    ifShow: (_column) => {
      return true; // 根据业务控制是否显示
    },
    customRender: ({ record }) => {
      if (!Reflect.has(record, 'pendingStatus')) {
        record.pendingStatus = false;
      }
      return h(Switch, {
        checked: record.status === 1,
        checkedChildren: t('global.statusDisable'),
        unCheckedChildren: t('global.statusEnable'),
        loading: record.pendingStatus,
        onChange(checked: CheckedType) {
          record.pendingStatus = true;
          const newStatus = checked ? 1 : 2;
          const { createMessage } = useMessage();
          setRoleStatus(record.id, newStatus)
            .then((res) => {
              console.log('setRoleStatus => ', res);
              if (res) {
                record.status = newStatus;
                createMessage.success(`已成功修改角色状态`);
              }
            })
            .catch(() => {
              createMessage.error('修改角色状态失败');
            })
            .finally(() => {
              record.pendingStatus = false;
            });
        },
      });
    },
  },
  {
    title: t('global.creationTime'),
    dataIndex: 'createTime',
    width: 160,
  },
  {
    title: t('global.updateTime'),
    dataIndex: 'updateTime',
    width: 160,
  },
  {
    title: t('global.remark'),
    dataIndex: 'remarks',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'roleNme',
    label: t('system.role.roleName'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'roleCode',
    label: t('system.role.roleCode'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'status',
    label: t('global.status'),
    component: 'Select',
    componentProps: {
      options: [
        { label: t('global.statusEnable'), value: 1 },
        { label: t('global.statusDisable'), value: 2 },
      ],
    },
    colProps: { span: 8 },
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'id',
    label: 'ID',
    required: false,
    component: 'InputNumber',
    show: false,
  },
  {
    field: 'roleName',
    label: t('system.role.roleName'),
    required: true,
    component: 'Input',
  },
  {
    field: 'roleCode',
    label: t('system.role.roleCode'),
    required: true,
    component: 'Input',
  },
  {
    field: 'status',
    label: t('global.status'),
    component: 'RadioButtonGroup',
    defaultValue: 1,
    componentProps: {
      options: [
        { label: t('global.statusEnable'), value: 1 },
        { label: t('global.statusDisable'), value: 2 },
      ],
    },
  },
  {
    label: t('global.remark'),
    field: 'remarks',
    component: 'InputTextArea',
  },
  {
    label: ' ',
    field: 'menu',
    slot: 'menu',
  },
];
