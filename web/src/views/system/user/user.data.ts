import { roleSelect } from '@/api/system/role';
import { BasicColumn, FormSchema } from '@/components/Table';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import { useI18n } from '@/hooks/web/useI18n';

const { t } = useI18n();

export const columns: BasicColumn[] = [
  {
    title: t('system.user.userName'),
    dataIndex: 'name',
    width: 200,
  },
  {
    title: t('system.user.email'),
    dataIndex: 'email',
    width: 300,
  },
  {
    title: t('system.user.role'),
    dataIndex: 'roleName',
    width: 100,
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
  {
    title: t('global.remark'),
    dataIndex: 'remarks',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'name',
    label: t('system.user.userName'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'email',
    label: t('system.user.email'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'roleId',
    label: t('system.user.role'),
    component: 'ApiSelect',
    componentProps: {
      api: roleSelect,
      labelField: 'label',
      valueField: 'value',
    },
    // required: true,
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

export const accountFormSchema: FormSchema[] = [
  {
    field: 'id',
    label: 'ID',
    component: 'InputNumber',
    required: false,
    show: false,
  },
  {
    field: 'name',
    label: t('system.user.userName'),
    component: 'Input',
    required: true,
    // helpMessage: ['本字段演示异步验证', '不能输入带有admin的用户名'],
    // rules: [
    //   {
    //     required: true,
    //     message: '请输入用户名',
    //   },
    //   {
    //     trigger: 'blur',
    //     validator(_, value) {
    //       return new Promise((resolve, reject) => {
    //         if (!value) return resolve();
    //         isAccountExist(value)
    //           .then(resolve)
    //           .catch((err) => {
    //             reject(err.message || '验证失败');
    //           });
    //       });
    //     },
    //   },
    // ],
  },
  // {
  //   field: 'pwd',
  //   label: '密码',
  //   component: 'InputPassword',
  //   required: true,
  //   ifShow: false,
  // },
  {
    label: t('system.user.role'),
    field: 'roleId',
    component: 'ApiSelect',
    componentProps: {
      api: roleSelect,
      labelField: 'label',
      valueField: 'value',
    },
    required: true,
  },
  {
    label: t('system.user.email'),
    field: 'email',
    component: 'Input',
    required: true,
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
    required: true,
  },
  {
    label: t('global.remark'),
    field: 'remarks',
    component: 'InputTextArea',
  },
  {
    label: ' ',
    field: 'store',
    slot: 'store',
    required: true,
  },
];
