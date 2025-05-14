import { defHttp } from '@/utils/http/axios';
import { LoginParams, LoginResultModel, GetUserInfoModel } from './model/userModel';

import { ErrorMessageMode } from '#/axios';

enum Api {
  Login = '/user/login',
  Logout = '/user/logout',
  GetUserInfo = '/user/info',
  GetUserMerchantList = '/user/my-merchant-list',
  UserChangeMerchant = '/user/change-merchant',
  GetPermCode = '/user/permission-code',
  ChangePassword = '/user/change-password',
  UpdateBasicInfo = '/user/update-basic-info',
  GetResetCode = '/user/reset-verification-code',
  VerifyResetCode = '/user/verify-reset-verification-code',
  SubmitResetPassword = '/user/submit-reset-password',
  TestRetry = '/user/testRetry',
}

/**
 * @description: user login api
 */
export function loginApi(params: LoginParams, mode: ErrorMessageMode = 'modal') {
  return defHttp.post<LoginResultModel>(
    {
      url: Api.Login,
      params,
    },
    {
      errorMessageMode: mode,
    },
  );
}

/**
 * @description: submit reset password
 */
export const submitResetPassword = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.SubmitResetPassword, params }, { successMessageMode: 'message' });

/**
 * @description: get verification code for reset password
 */
export const getVerificationCodeForResetPassword = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.GetResetCode, params }, { successMessageMode: 'message' });

/**
 * @description: get verification code for reset password
 */
export const verifyVerificationCodeForResetPassword = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.VerifyResetCode, params }, { errorMessageMode: 'message' });

/**
 * @description: update basic info
 */
export const updateBasicInfo = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.UpdateBasicInfo, params }, { successMessageMode: 'message' });

/**
 * @description: user change password
 */
export const changePassword = (params?: Recordable<any>) =>
  defHttp.post({ url: Api.ChangePassword, params }, { successMessageMode: 'message' });

/**
 * @description: getUserInfo
 */
export function getUserInfo() {
  return defHttp.get<GetUserInfoModel>({ url: Api.GetUserInfo }, { errorMessageMode: 'modal' });
}

/**
 * @description: 获取用户授权的商户列表
 */
export function getUserMerchantList() {
  return defHttp.get({ url: Api.GetUserMerchantList }, { errorMessageMode: 'modal' });
}

/**
 * @description: 用户切换商户
 */
export function userChangeMerchant(id: number) {
  return defHttp.post(
    { url: Api.UserChangeMerchant, params: { id } },
    { errorMessageMode: 'modal' },
  );
}

export function getPermCode() {
  return defHttp.get<string[]>({ url: Api.GetPermCode });
}

export function doLogout() {
  return defHttp.get({ url: Api.Logout });
}

// Web Crypto API 加密密码
export async function hashStringWithSHA256(input) {
  // 将字符串转换为 ArrayBuffer
  const msgBuffer = new TextEncoder().encode(input);

  // 创建一个哈希算法实例 crypto 需要在https下使用
  const hashBuffer = await window.crypto.subtle.digest('SHA-256', msgBuffer);

  // 将 ArrayBuffer 转换为十六进制字符串
  const hashArray = Array.from(new Uint8Array(hashBuffer));
  return hashArray.map((b) => b.toString(16).padStart(2, '0')).join('');
}

export function testRetry() {
  return defHttp.get(
    { url: Api.TestRetry },
    {
      retryRequest: {
        isOpenRetry: true,
        count: 5,
        waitTime: 1000,
      },
    },
  );
}
