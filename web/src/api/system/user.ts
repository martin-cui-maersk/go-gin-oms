import { UserParams, UserListGetResultModel } from './model/systemModel';
import { defHttp } from '@/utils/http/axios';

enum Api {
  UserList = '/system/user-list',
  UserAdd = '/system/add-user',
  UserUpdate = '/system/update-user',
}

export const addUser = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.UserAdd, params }, { successMessageMode: 'message' });

export const updateUser = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.UserUpdate, params }, { successMessageMode: 'message' });

export const getUserList = (params: UserParams) =>
  defHttp.get<UserListGetResultModel>({ url: Api.UserList, params });

export const isAccountExist = (params: UserParams) => defHttp.get({ url: Api.UserList, params });
