import { BasicPageParams, BasicFetchResult, BasicResponseResult } from '@/api/model/baseModel';

export type UserParams = BasicPageParams & {
  name?: string;
  email?: string;
  [key: string]: any;
};

export type RoleParams = {
  roleName?: string;
  status?: string;
};

export type RolePageParams = BasicPageParams & RoleParams;

export type MenuParams = {
  menuName?: string;
  status?: string;
};

export interface AccountListItem {
  id: string;
  account: string;
  email: string;
  userName: string;
  role: number;
  createTime: string;
  remark: string;
  status: number;
}

export interface DeptListItem {
  id: string;
  orderNo: string;
  createTime: string;
  remark: string;
  status: number;
}

export interface MenuListItem {
  id: string;
  orderNo: string;
  createTime: string;
  status: number;
  icon: string;
  component: string;
  permission: string;
}

export interface RoleListItem {
  id: string;
  roleName: string;
  roleValue: string;
  status: number;
  orderNo: string;
  createTime: string;
}

/**
 * @description: Request list return value
 */
export type UserListGetResultModel = BasicFetchResult<AccountListItem>;

export type MenuListGetResultModel = BasicFetchResult<MenuListItem>;

export type RolePageListGetResultModel = BasicFetchResult<RoleListItem>;

export type StoreListGetResultModel = BasicFetchResult<BasicResponseResult>;

export type RoleListGetResultModel = RoleListItem[];
