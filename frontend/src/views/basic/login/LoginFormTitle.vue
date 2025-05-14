<template>
  <h2 class="mb-3 text-2xl font-bold text-center xl:text-3xl enter-x xl:text-left">
    {{ getFormTitle }}
  </h2>
</template>
<script lang="ts" setup>
  import { computed, unref } from 'vue';
  import { useI18n } from '@/hooks/web/useI18n';
  import { LoginStateEnum, useLoginState } from './useLogin';

  const { t } = useI18n();

  const { getLoginState } = useLoginState();

  const getFormTitle = computed(() => {
    const titleObj = {
      [LoginStateEnum.RESET_PASSWORD]: t('basic.login.getVerificationCodeFormTitle'),
      [LoginStateEnum.LOGIN]: t('basic.login.signInFormTitle'),
      [LoginStateEnum.RESET_PASSWORD_SUBMIT]: t('basic.login.forgetFormTitle'),
    };
    return titleObj[unref(getLoginState)];
  });
</script>
