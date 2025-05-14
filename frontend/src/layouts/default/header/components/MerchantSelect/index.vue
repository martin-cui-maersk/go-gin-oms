<template>
  <div
    @click="goSelectMerchant"
    :title="
      userInfo.merchantIds.length > 1 ? t('routes.account.selectMerchant') : userInfo.merchantName
    "
  >
    <span class="cursor-pointer flex items-center" style="margin-left: -5px; color: #1a1a1a">
      <Icon icon="ant-design:shop-twotone" />
      <span style="margin-left: -5px; font-size: 13px">{{ userInfo.merchantName }}</span>
    </span>
  </div>
</template>
<script lang="ts" setup>
  import Icon from '@/components/Icon/Icon.vue';
  import { useUserStore } from '@/store/modules/user';
  import { router } from '@/router';
  import { useI18n } from '@/hooks/web/useI18n';

  defineOptions({ name: 'MerchantSelect' });
  const { t } = useI18n();

  const userStore = useUserStore();
  const userInfo = userStore.getUserInfo;
  console.log('用户可切换的商户列表', userInfo.merchantIds);

  function goSelectMerchant() {
    if (userInfo.merchantIds && userInfo.merchantIds.length > 1) {
      router.replace({
        path: '/my-account/select-merchant',
      });
    } else {
      console.log('no merchant select');
    }
  }
</script>
