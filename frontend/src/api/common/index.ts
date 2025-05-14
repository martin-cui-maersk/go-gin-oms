import { useUserStore } from '@/store/modules/user';

/**
 * 用户店铺列表
 */
export const getUserStoreList = () => {
  const userStore = useUserStore();
  const userInfo = userStore.getUserInfo;
  return userInfo.storeList;
};
