import { FormSchema } from '@/components/Table';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import { useI18n } from '@/hooks/web/useI18n';

const { t } = useI18n();

export const columns = [
  {
    title: t('system.merchant.merchantName'),
    dataIndex: 'merchantName',
    width: 180,
  },
  {
    title: t('system.merchant.merchantCode'),
    dataIndex: 'merchantCode',
    width: 180,
  },
  {
    title: t('system.merchant.contact'),
    dataIndex: 'contact',
    width: 120,
  },
  {
    title: t('system.merchant.email'),
    dataIndex: 'email',
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
    field: 'merchantName',
    label: t('system.merchant.merchantName'),
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

export const storeFormSchema: FormSchema[] = [
  {
    field: 'id',
    label: 'ID',
    component: 'InputNumber',
    required: false,
    show: false,
  },
  {
    field: 'merchantName',
    label: t('system.merchant.merchantName'),
    component: 'Input',
    required: true,
  },
  {
    field: 'merchantCode',
    label: t('system.merchant.merchantCode'),
    component: 'Input',
    required: true,
  },
  {
    field: 'contact',
    label: t('system.merchant.contact'),
    component: 'Input',
    required: true,
  },
  {
    field: 'email',
    label: t('system.merchant.email'),
    component: 'Input',
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
