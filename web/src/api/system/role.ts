import { RolePageParams, RolePageListGetResultModel } from './model/systemModel';
import { defHttp } from '@/utils/http/axios';

enum Api {
  RoleStatusSet = '/system/set-role-status',
  RolePageList = '/system/role-list',
  RoleAdd = '/system/add-role',
  RoleUpdate = '/system/update-role',
  RoleSelect = '/system/role-ids',
}

export const roleSelect = () => defHttp.get({ url: Api.RoleSelect });

export const addRole = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.RoleAdd, params }, { successMessageMode: 'message' });

export const updateRole = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.RoleUpdate, params }, { successMessageMode: 'message' });

export const getRoleList = (params?: RolePageParams) =>
  defHttp.get<RolePageListGetResultModel>({ url: Api.RolePageList, params });

export const setRoleStatus = (id: number, status: number) =>
  defHttp.post({ url: Api.RoleStatusSet, params: { id, status } });
