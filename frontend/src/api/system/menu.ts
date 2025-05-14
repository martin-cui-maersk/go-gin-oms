import { MenuParams, MenuListGetResultModel } from './model/systemModel';
import { defHttp } from '@/utils/http/axios';

enum Api {
  MenuList = '/system/menu-list',
  MenuAdd = '/system/add-menu',
  MenuUpdate = '/system/update-menu',
  MenuDelete = '/system/delete-menu',
}

export const addMenu = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.MenuAdd, params }, { successMessageMode: 'message' });

export const updateMenu = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.MenuUpdate, params }, { successMessageMode: 'message' });

export const deleteMenu = (id: number) =>
  defHttp.post({ url: Api.MenuDelete, params: { id } }, { successMessageMode: 'message' });

export const getMenuList = (params?: MenuParams) =>
  defHttp.get<MenuListGetResultModel>({ url: Api.MenuList, params });

export const getActiveMenuList = () =>
  defHttp.get<MenuListGetResultModel>({ url: Api.MenuList, params: { status: 1 } });
