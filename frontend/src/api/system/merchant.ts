import { defHttp } from '@/utils/http/axios';

enum Api {
  MerchantList = '/system/merchant-list',
  MerchantAdd = '/system/add-merchant',
  MerchantUpdate = '/system/update-merchant',
  MerchantSelect = '/system/merchant-ids',
}

export const getMerchantList = (params?: Recordable<any>) =>
  defHttp.get({ url: Api.MerchantList, params });

export const addMerchant = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.MerchantAdd, params }, { successMessageMode: 'message' });

export const updateMerchant = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.MerchantUpdate, params }, { successMessageMode: 'message' });

export const merchantSelect = () =>
  defHttp.get({ url: Api.MerchantSelect }, { errorMessageMode: 'message' });

/**
 * define platform
 */
export const platformList = [
  { label: '-', value: '' },
  { label: 'Amazon', value: 'amazon' },
  { label: 'TikTok', value: 'tiktok' },
  { label: 'Douyin', value: 'douyin' },
  { label: 'Shopify', value: 'shopify' },
  { label: 'WeChat', value: 'wechat' },
  { label: 'HKTVMall', value: 'hktvmall' },
  { label: 'Rakuten', value: 'rakuten' },
  { label: 'Shopee', value: 'shopee' },
  { label: 'Lazada', value: 'lazada' },
  { label: 'OMS', value: 'oms' },
  { label: '天猫', value: 'tmall' },
  { label: '小红书', value: 'red' },
  { label: '拼多多', value: 'pdd' },
  { label: 'POS', value: 'pos' },
];

export const platformMap = () => {
  const map = [];
  platformList.forEach((item) => {
    map[item.value] = item.label;
  });
  return map;
};
