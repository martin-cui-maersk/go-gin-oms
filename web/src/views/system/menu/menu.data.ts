import { BasicColumn, FormSchema } from '@/components/Table';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import Icon from '@/components/Icon/Icon.vue';
import { useI18n } from '@/hooks/web/useI18n';

const { t } = useI18n();
export const columns: BasicColumn[] = [
  {
    title: t('system.menu.menuName'),
    dataIndex: 'menuName',
    width: 250,
    align: 'left',
    fixed: 'left',
  },
  {
    title: t('system.menu.icon'),
    dataIndex: 'icon',
    width: 70,
    customRender: ({ record }) => {
      return h(Icon, { icon: record.icon });
    },
    align: 'center',
  },
  {
    title: t('system.menu.type'),
    dataIndex: 'type',
    width: 80,
    customRender: ({ record }) => {
      return menuType(record.type);
    },
    align: 'center',
  },
  {
    title: t('system.menu.permissionCode'),
    dataIndex: 'permission',
    width: 140,
    align: 'center',
  },
  {
    title: t('system.menu.routePath'),
    dataIndex: 'path',
    width: 200,
    align: 'center',
  },
  {
    title: t('system.menu.apiPath'),
    dataIndex: 'apiPath',
    width: 200,
    align: 'center',
  },
  {
    title: t('system.menu.component'),
    dataIndex: 'component',
    width: 200,
  },
  {
    title: t('system.menu.sort'),
    dataIndex: 'sort',
    width: 50,
  },
  {
    title: t('global.status'),
    dataIndex: 'status',
    width: 80,
    customRender: ({ record }) => {
      const status = record.status;
      const enable = status === 1;
      const color = enable ? 'green' : 'red';
      const text = enable ? t('global.statusEnable') : t('global.statusDisable');
      return h(Tag, { color: color }, () => text);
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
];

// const isDir = (type) => type === 0;
const isMenu = (type) => type === 1;
const isButton = (type) => type === 2;

const menuType = function (type) {
  let color = '';
  let text = '';
  switch (type) {
    case 0:
      color = 'orange';
      text = t('system.menu.typeDir');
      break;
    case 1:
      color = 'green';
      text = t('system.menu.typeMenu');
      break;
    case 2:
      color = 'blue';
      text = t('system.menu.typeButton');
      break;
  }
  return h(Tag, { color: color }, () => text);
};

export const searchFormSchema: FormSchema[] = [
  {
    field: 'menuName',
    label: t('system.menu.menuName'),
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
    component: 'InputNumber',
    required: false,
    show: false,
  },
  {
    field: 'type',
    label: t('system.menu.type'),
    component: 'RadioButtonGroup',
    defaultValue: 0,
    componentProps: {
      options: [
        { label: t('system.menu.typeDir'), value: 0 },
        { label: t('system.menu.typeMenu'), value: 1 },
        { label: t('system.menu.typeButton'), value: 2 },
      ],
    },
    colProps: { lg: 24, md: 24 },
  },
  {
    field: 'menuName',
    label: () => t('system.menu.menuName'),
    component: 'Input',
    required: true,
  },
  {
    field: 'menuTitle',
    label: t('system.menu.langPack'),
    component: 'Input',
    required: ({ values }) => !isButton(values.type),
  },
  {
    field: 'parentId',
    label: t('system.menu.parentMenu'),
    component: 'TreeSelect',
    componentProps: {
      fieldNames: {
        label: 'menuName',
        value: 'id',
      },
      getPopupContainer: () => document.body,
    },
  },

  {
    field: 'sort',
    label: t('system.menu.sort'),
    component: 'InputNumber',
    required: true,
  },
  {
    field: 'icon',
    label: t('system.menu.icon'),
    component: 'IconPicker',
    required: false,
    ifShow: ({ values }) => !isButton(values.type),
  },
  {
    field: 'path',
    label: t('system.menu.routePath'),
    component: 'Input',
    // required: true,
    ifShow: ({ values }) => !isButton(values.type),
  },
  {
    field: 'apiPath',
    label: t('system.menu.apiPath'),
    component: 'Input',
    // required: true,
    ifShow: ({ values }) => isButton(values.type),
  },
  {
    field: 'component',
    label: t('system.menu.component'),
    component: 'Input',
    ifShow: ({ values }) => isMenu(values.type),
  },
  {
    field: 'permission',
    label: t('system.menu.permissionCode'),
    component: 'Input',
    ifShow: ({ values }) => isButton(values.type),
  },
  {
    field: 'currentActiveMenu',
    label: t('system.menu.currentActiveMenu'),
    component: 'Input',
    required: false,
    ifShow: ({ values }) => !isButton(values.type),
  },
  {
    field: 'redirect',
    label: t('system.menu.redirect'),
    component: 'Input',
    required: false,
    ifShow: ({ values }) => !isButton(values.type),
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
    field: 'keepalive',
    label: '是否缓存',
    component: 'RadioButtonGroup',
    defaultValue: 1,
    componentProps: {
      options: [
        { label: '否', value: 0 },
        { label: '是', value: 1 },
      ],
    },
    ifShow: ({ values }) => isMenu(values.type),
  },

  {
    field: 'show',
    label: '是否显示',
    component: 'RadioButtonGroup',
    defaultValue: 1,
    componentProps: {
      options: [
        { label: '是', value: 1 },
        { label: '否', value: 0 },
      ],
    },
    ifShow: ({ values }) => !isButton(values.type),
  },
];
