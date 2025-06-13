import { platformList, platformMap, merchantSelect } from '@/api/system/merchant';
// import { isStoreNameExist } from '@/api/system/store';
import { BasicColumn, FormSchema } from '@/components/Table';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import { useI18n } from '@/hooks/web/useI18n';

const { t } = useI18n();

export const columns: BasicColumn[] = [
  {
    title: t('system.store.storeName'),
    dataIndex: 'storeName',
    width: 200,
  },
  {
    title: t('system.store.storeCode'),
    dataIndex: 'storeCode',
    width: 150,
  },
  {
    title: t('system.store.shopId'),
    dataIndex: 'shopId',
    width: 200,
  },
  {
    title: t('system.store.platform'),
    dataIndex: 'platform',
    width: 100,
    customRender: ({ text }) => {
      return platformMap()[text] ?? '';
    },
  },
  {
    title: t('system.store.merchant'),
    dataIndex: 'merchantName',
    width: 180,
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
    field: 'storeName',
    label: t('system.store.storeName'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'storeCode',
    label: t('system.store.storeCode'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'shopId',
    label: t('system.store.shopId'),
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'platform',
    label: t('system.store.platform'),
    component: 'Select',
    componentProps: {
      options: platformList,
    },
    colProps: { span: 8 },
  },
  {
    field: 'merchant',
    label: t('system.store.merchant'),
    component: 'ApiSelect',
    componentProps: {
      api: merchantSelect,
      labelField: 'label',
      valueField: 'value',
    },
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

export const storeFormSchema: FormSchema[] = [
  {
    field: 'id',
    label: 'ID',
    component: 'InputNumber',
    required: false,
    show: false,
  },
  {
    field: 'storeName',
    label: t('system.store.storeName'),
    component: 'Input',
    required: true,
    // helpMessage: ['本字段异步验证', '店铺名不能重复'],
    // rules: [
    //   {
    //     required: true,
    //     message: '请输入店铺名',
    //   },
    //   {
    //     trigger: 'blur',
    //     validator(_, value) {
    //       return new Promise((resolve, reject) => {
    //         if (!value) return resolve();
    //         isStoreNameExist(value)
    //           .then(resolve)
    //           .catch((err) => {
    //             reject(err.message || '验证失败');
    //           });
    //       });
    //     },
    //   },
    // ],
  },
  {
    field: 'storeCode',
    label: t('system.store.storeCode'),
    component: 'Input',
    required: true,
  },
  {
    field: 'shopId',
    label: t('system.store.shopId'),
    component: 'Input',
    required: true,
  },
  {
    field: 'platform',
    label: 'Platform',
    component: 'Select',
    componentProps: {
      options: platformList,
    },
    required: true,
  },
  {
    field: 'merchant',
    label: t('system.store.merchant'),
    component: 'ApiSelect',
    componentProps: {
      api: merchantSelect,
      labelField: 'label',
      valueField: 'value',
    },
    required: true,
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
];
