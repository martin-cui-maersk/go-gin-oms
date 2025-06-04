// axios配置  可自行根据项目进行更改，只需更改该文件即可，其他文件可以不动
// The axios configuration can be changed according to the project, just change the file, other files can be left unchanged

import type { AxiosInstance, AxiosResponse } from 'axios';
import { clone } from 'lodash-es';
import type { RequestOptions, Result } from '#/axios';
import type { AxiosTransform, CreateAxiosOptions } from './axiosTransform';
import { VAxios } from './Axios';
import { checkStatus } from './checkStatus';
import { useGlobSetting } from '@/hooks/setting';
import { useMessage } from '@/hooks/web/useMessage';
import { RequestEnum, ResultEnum, ContentTypeEnum } from '@/enums/httpEnum';
import { isString, isUndefined, isNull, isEmpty } from '@/utils/is';
import { getToken } from '@/utils/auth';
import { setObjToUrlParams, deepMerge } from '@/utils';
import { useErrorLogStoreWithOut } from '@/store/modules/errorLog';
import { useI18n } from '@/hooks/web/useI18n';
import { joinTimestamp, formatRequestDate } from './helper';
import { useUserStoreWithOut } from '@/store/modules/user';
import { AxiosRetry } from '@/utils/http/axios/axiosRetry';
import axios from 'axios';
import { h } from 'vue';
import { useLocaleStoreWithOut } from '@/store/modules/locale';
import {LoginStateEnum, useLoginState} from "@/views/basic/login/useLogin";

const { createConfirm } = useMessage();

const globSetting = useGlobSetting();
const urlPrefix = globSetting.urlPrefix;
const { createMessage, createErrorModal, createSuccessModal } = useMessage();

/**
 * @description: 数据处理，方便区分多种处理方式
 */
const transform: AxiosTransform = {
  /**
   * @description: 处理响应数据。如果数据不是预期格式，可直接抛出错误
   */
  transformResponseHook: (res: AxiosResponse<Result>, options: RequestOptions) => {
    const { t } = useI18n();
    const { isTransformResponse, isReturnNativeResponse } = options;
    // 是否返回原生响应头 比如：需要获取响应头时使用该属性
    if (isReturnNativeResponse) {
      return res;
    }
    // 不进行任何处理，直接返回
    // 用于页面代码可能需要直接获取code，data，message这些信息时开启
    if (!isTransformResponse) {
      return res.data;
    }
    // 错误的时候返回
    // console.log('typeof response', typeof response);
    // const { response } = res;
    const response = res.data;
    console.log('transformResponseHook Response =>', response);
    if (!response) {
      createMessage.error(t('basic.api.apiRequestFailed'));
      // return '[HTTP] Request has no return value';
      throw new Error(t('basic.api.apiRequestFailed'));
    }
    //  这里 code，result，message为 后台统一的字段，需要在 types.ts内修改为项目自己的接口返回格式
    const { code, data, msg } = response;

    // 这里逻辑可以根据项目进行修改
    const hasSuccess = response && Reflect.has(response, 'code') && code === ResultEnum.SUCCESS;
    console.log('hasSuccess => ', hasSuccess);
    if (hasSuccess) {
      let successMsg = msg;

      if (isNull(successMsg) || isUndefined(successMsg) || isEmpty(successMsg)) {
        successMsg = t(`sys.api.operationSuccess`);
      }

      if (options.successMessageMode === 'modal') {
        createSuccessModal({ title: t('basic.api.successTip'), content: successMsg });
      } else if (options.successMessageMode === 'message') {
        createMessage.success(successMsg);
      }
      // console.log('transformResponseHook hasSuccess Return => ', data);
      return Object.keys(data).length === 0 && data.constructor === Object ? true : data;
    }

    // 在此处根据自己项目的实际情况对不同的code执行不同的操作
    // 如果不希望中断当前请求，请return数据，否则直接抛出异常即可
    let timeoutMsg = '';
    switch (code) {
      case ResultEnum.TIMEOUT:
        timeoutMsg = t('basic.api.timeoutMessage');
        const userStore = useUserStoreWithOut();
        // 被动登出，带redirect地址
        userStore.logout(false);
        break;
      case ResultEnum.LOGIN_FAILED: // 登录失败
        createMessage.error(t('basic.login.loginFailed'));
        break;
      case ResultEnum.SYS_ERROR: // code = 500 系统异常
        const error500 = sessionStorage.getItem('error500');
        if (error500 === null) {
          // createErrorModal({ title: t('basic.api.errorTip'), content: t('basic.api.errMsg500') });
          sessionStorage.setItem('error500', String(1));
          setTimeout(() => {
            sessionStorage.removeItem('error500');
          }, 3000);
          createConfirm({
            iconType: 'warning',
            okCancel: false,
            title: () => h('span', t('basic.api.errorTip')),
            content: () => h('span', t('basic.api.errMsg500')),
            onOk: async () => {
              sessionStorage.removeItem('error500');
            },
          });
        } else {
          setTimeout(() => {
            sessionStorage.removeItem('error500');
          }, 3000);
        }
        break;
      case ResultEnum.ERROR: // code = 503 业务接口异常
        if (options.errorMessageMode === 'modal') {
          createErrorModal({ title: t('basic.api.errorTip'), content: msg });
        } else if (options.errorMessageMode === 'message') {
          createMessage.error(msg);
        }
        break;
      case ResultEnum.API_NO_PERMISSION: // code = 401 接口无权限
        const error401 = sessionStorage.getItem('error401');
        if (error401 === null) {
          createErrorModal({ title: t('basic.api.errorTip'), content: t('basic.api.errMsg401') });
          sessionStorage.setItem('error401', String(1));
          setTimeout(() => {
            sessionStorage.removeItem('error401');
          }, 2000);
        }
        break;
      case ResultEnum.INVALID_MERCHANT: // code = 601 绑定的商户异常
        createErrorModal({ title: t('basic.api.errorTip'), content: msg });
        break;
      case ResultEnum.INVALID_TOKEN: // code = 600 token失效
        const loginExpired = sessionStorage.getItem('loginExpired');
        if (loginExpired === null) {
          sessionStorage.setItem('loginExpired', String(1));
          createConfirm({
            iconType: 'warning',
            okCancel: false,
            title: () => h('span', t('basic.app.logoutTip')),
            content: () => h('span', t('basic.login.expired')),
            onOk: async () => {
              sessionStorage.removeItem('loginExpired');
              const userStore = useUserStoreWithOut();
              userStore.redirectLogin();
            },
          });
        } else {
          const userStore = useUserStoreWithOut();
          userStore.redirectLogin();
        }
        break;
      case ResultEnum.LOGIN_FAILED_WITH_INITIAL_PASSWORD: // code = 604 初始密码登录，强制跳转修改密码
        createConfirm({
          iconType: 'warning',
          okCancel: false,
          title: () => h('span', t('global.reminder')),
          content: () => h('span', t('basic.login.initialPasswordError')),
          onOk: async () => {
            // TODO::
            const { setLoginState } = useLoginState();
            setLoginState(LoginStateEnum.CHANGE_INITIAL_PASSWORD);
          },
        });
        break;
      case ResultEnum.LOGIN_FAILED_WITH_EXPIRED_PASSWORD: // code = 605 密码过期，强制跳转修改密码
        createConfirm({
          iconType: 'warning',
          okCancel: false,
          title: () => h('span', t('global.reminder')),
          content: () => h('span', t('basic.login.expiredPasswordError')),
          onOk: async () => {
            // TODO::
            const { setLoginState } = useLoginState();
            setLoginState(LoginStateEnum.CHANGE_INITIAL_PASSWORD);
          },
        });
        break;
      case ResultEnum.LOCKED_ACCOUNT: // code = 606 账号被锁定
        createConfirm({
          iconType: 'warning',
          okCancel: false,
          title: () => h('span', t('global.reminder')),
          content: () => h('span', t('basic.login.lockedError')),
          onOk: async () => {
            // TODO::
          },
        });
        break;
      default:
        if (options.errorMessageMode === 'modal') {
          createErrorModal({ title: t('basic.api.errorTip'), content: msg });
        } else if (options.errorMessageMode === 'message') {
          createMessage.error(msg);
        }
    }
    console.log('transformResponseHook Error Return => ', data);
    return data.length > 0 ? data : false;
    console.log('timeoutMsg', timeoutMsg);
    // errorMessageMode='modal'的时候会显示modal错误弹窗，而不是消息提示，用于一些比较重要的错误
    // errorMessageMode='none' 一般是调用时明确表示不希望自动弹出错误提示
    if (options.errorMessageMode === 'modal') {
      createErrorModal({ title: t('basic.api.errorTip'), content: timeoutMsg });
    } else if (options.errorMessageMode === 'message') {
      createMessage.error(timeoutMsg);
    }

    throw new Error(timeoutMsg || t('basic.api.apiRequestFailed'));
  },

  // 请求之前处理config
  beforeRequestHook: (config, options) => {
    const { apiUrl, joinPrefix, joinParamsToUrl, formatDate, joinTime = true, urlPrefix } = options;

    // 防止二次token失效弹窗
    const loginExpired = sessionStorage.getItem('loginExpired');
    console.log(config, loginExpired);
    if (
      config.url !== '/user/login' &&
      config.url !== '/user/reset-verification-code' &&
      config.url !== '/user/verify-reset-verification-code' &&
      config.url !== '/user/submit-reset-password' &&
      config.url !== '/user/login-verify' &&
      loginExpired !== null
    ) {
      const userStore = useUserStoreWithOut();
      userStore.redirectLogin();
      console.log('stop request as loginExpired not null');
      throw new Error('stop request as loginExpired not null.');
    }

    if (joinPrefix) {
      config.url = `${urlPrefix}${config.url}`;
    }

    if (apiUrl && isString(apiUrl)) {
      config.url = `${apiUrl}${config.url}`;
    }
    const params = config.params || {};
    const data = config.data || false;
    formatDate && data && !isString(data) && formatRequestDate(data);
    if (config.method?.toUpperCase() === RequestEnum.GET) {
      if (!isString(params)) {
        // 给 get 请求加上时间戳参数，避免从缓存中拿数据。
        config.params = Object.assign(params || {}, joinTimestamp(joinTime, false));
      } else {
        // 兼容restful风格
        config.url = config.url + params + `${joinTimestamp(joinTime, true)}`;
        config.params = undefined;
      }
    } else {
      if (!isString(params)) {
        formatDate && formatRequestDate(params);
        if (
          Reflect.has(config, 'data') &&
          config.data &&
          (Object.keys(config.data).length > 0 || config.data instanceof FormData)
        ) {
          config.data = data;
          config.params = params;
        } else {
          // 非GET请求如果没有提供data，则将params视为data
          config.data = params;
          config.params = undefined;
        }
        if (joinParamsToUrl) {
          config.url = setObjToUrlParams(
            config.url as string,
            Object.assign({}, config.params, config.data),
          );
        }
      } else {
        // 兼容restful风格
        config.url = config.url + params;
        config.params = undefined;
      }
    }
    return config;
  },

  /**
   * @description: 请求拦截器处理
   */
  requestInterceptors: (config, options) => {
    // 请求之前处理config
    const token = 'Bearer ' + getToken();
    // console.log('requestInterceptors token', token);
    if (token && (config as Recordable)?.requestOptions?.withToken !== false) {
      // jwt token
      (config as Recordable).headers.Authorization = options.authenticationScheme
        ? `${options.authenticationScheme} ${token}`
        : token;
    }

    // header 头添加时区
    const localeStore = useLocaleStoreWithOut();
    // console.log(localeStore.getTimeZone);
    config.headers.TimeZone = localeStore.getTimeZone;
    config.headers.Lang = localeStore.getLocale;

    return config;
  },

  /**
   * @description: 响应拦截器处理
   */
  responseInterceptors: (res: AxiosResponse<any>) => {
    let responseToken = res.headers.authorization;
    console.log('responseInterceptors responseToken => ', responseToken);
    if (responseToken) {
      // set new token from response header
      responseToken = responseToken.replace('Bearer ', '');
      const userStore = useUserStoreWithOut();
      userStore.setNewToken(responseToken);
    }
    return res;
  },

  /**
   * @description: 响应错误处理
   */
  responseInterceptorsCatch: (axiosInstance: AxiosInstance, error: any) => {
    const { t } = useI18n();
    const errorLogStore = useErrorLogStoreWithOut();
    errorLogStore.addAjaxErrorInfo(error);
    const { response, code, message, config } = error || {};
    const errorMessageMode = config?.requestOptions?.errorMessageMode || 'none';
    const msg: string = response?.data?.error?.message ?? '';
    const err: string = error?.toString?.() ?? '';
    let errMessage = '';

    if (axios.isCancel(error)) {
      return Promise.reject(error);
    }

    try {
      if (code === 'ECONNABORTED' && message.indexOf('timeout') !== -1) {
        errMessage = t('basic.api.apiTimeoutMessage');
      }
      if (err?.includes('Network Error')) {
        errMessage = t('basic.api.networkExceptionMsg');
      }

      if (errMessage) {
        if (errorMessageMode === 'modal') {
          createErrorModal({ title: t('basic.api.errorTip'), content: errMessage });
        } else if (errorMessageMode === 'message') {
          createMessage.error(errMessage);
        }
        return Promise.reject(error);
      }
    } catch (error) {
      throw new Error(error as unknown as string);
    }

    checkStatus(error?.response?.status, msg, errorMessageMode);

    // 添加自动重试机制 保险起见 只针对GET请求
    const retryRequest = new AxiosRetry();
    const { isOpenRetry } = config.requestOptions.retryRequest;
    config.method?.toUpperCase() === RequestEnum.GET &&
      isOpenRetry &&
      error?.response?.status !== 401 &&
      // @ts-ignore
      retryRequest.retry(axiosInstance, error);
    return Promise.reject(error);
  },
};

function createAxios(opt?: Partial<CreateAxiosOptions>) {
  return new VAxios(
    // 深度合并
    deepMerge(
      {
        // See https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication#authentication_schemes
        // authentication schemes，e.g: Bearer
        // authenticationScheme: 'Bearer',
        authenticationScheme: '',
        timeout: 30 * 1000,
        // 基础接口地址
        // baseURL: globSetting.apiUrl,

        headers: { 'Content-Type': ContentTypeEnum.JSON },
        // 如果是form-data格式
        // headers: { 'Content-Type': ContentTypeEnum.FORM_URLENCODED },
        // 数据处理方式
        transform: clone(transform),
        // 配置项，下面的选项都可以在独立的接口请求中覆盖
        requestOptions: {
          // 默认将prefix 添加到url
          joinPrefix: true,
          // 是否返回原生响应头 比如：需要获取响应头时使用该属性
          isReturnNativeResponse: false,
          // 需要对返回数据进行处理
          isTransformResponse: true,
          // post请求的时候添加参数到url
          joinParamsToUrl: false,
          // 格式化提交参数时间
          formatDate: true,
          // 消息提示类型
          errorMessageMode: 'message',
          // 接口地址
          apiUrl: globSetting.apiUrl,
          // 接口拼接地址
          urlPrefix: urlPrefix,
          //  是否加入时间戳
          joinTime: true,
          // 忽略重复请求
          ignoreCancelToken: true,
          // 是否携带token
          withToken: true,
          retryRequest: {
            isOpenRetry: true,
            count: 0, // 接口重试次数
            waitTime: 500,
          },
        },
      },
      opt || {},
    ),
  );
}
export const defHttp = createAxios();

// other api url
// export const otherHttp = createAxios({
//   requestOptions: {
//     apiUrl: 'xxx',
//     urlPrefix: 'xxx',
//   },
// });
