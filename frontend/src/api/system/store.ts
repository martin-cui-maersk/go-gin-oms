import { defHttp } from '@/utils/http/axios';

enum Api {
  StoreTree = '/system/store-tree',
  StoreList = '/system/store-list',
  StoreAdd = '/system/add-store',
  StoreUpdate = '/system/update-store',
  StoreNameCheck = '/system/check-store-name',
}

export const isStoreNameExist = (storeName: string) =>
  defHttp.post({ url: Api.StoreNameCheck, params: { storeName } }, { errorMessageMode: 'none' });

export const getStoreList = (params?: Recordable<any>) =>
  defHttp.get({ url: Api.StoreList, params });

export const addStore = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.StoreAdd, params }, { successMessageMode: 'message' });

export const updateStore = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.StoreUpdate, params }, { successMessageMode: 'message' });

export const storeTree = () => defHttp.get({ url: Api.StoreTree }, { errorMessageMode: 'message' });
