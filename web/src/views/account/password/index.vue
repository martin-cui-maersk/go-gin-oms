<template>
  <PageWrapper :title="t('account.changePassword')" :content="t('account.changePasswordDesc')">
    <div class="py-8 bg-white flex flex-col justify-center items-center">
      <BasicForm @register="register" />
      <div class="flex justify-center">
        <a-button @click="resetFields"> {{ t('account.resetPassword') }} </a-button>
        <a-button class="!ml-4" type="primary" @click="handleSubmit">
          {{ t('account.submitPassword') }}
        </a-button>
      </div>
    </div>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { PageWrapper } from '@/components/Page';
  import { BasicForm, useForm } from '@/components/Form';
  import { useI18n } from '@/hooks/web/useI18n';
  import { formSchema } from './pwd.data';
  import { hashStringWithSHA256, changePassword } from '@/api/basic/user';
  // import { useRouter } from 'vue-router';
  // import { PageEnum } from '@/enums/pageEnum';
  import { useUserStoreWithOut } from '@/store/modules/user';

  defineOptions({ name: 'ChangePassword' });

  // const { router } = useRouter();
  const { t } = useI18n();
  const [register, { validate, resetFields }] = useForm({
    size: 'large',
    baseColProps: { span: 24 },
    labelWidth: 250,
    showActionButtonGroup: false,
    schemas: formSchema,
  });

  async function handleSubmit() {
    try {
      const values = await validate();
      const { passwordOld, passwordNew } = values;
      console.log(passwordOld, passwordNew);
      values.passwordOld = await hashStringWithSHA256(values.passwordOld);
      values.passwordNew = await hashStringWithSHA256(values.passwordNew);
      values.confirmPassword = await hashStringWithSHA256(values.confirmPassword);
      console.log(values);
      // TODO custom api
      changePassword(values).then((res) => {
        if (res) {
          // const { router } = useRouter();
          // router.push(PageEnum.BASE_LOGIN);
          setTimeout(() => {
            const userStore = useUserStoreWithOut();
            userStore.redirectLogin();
          }, 1000);
        }
      });
    } catch (error) {
      console.error(error);
    }
  }
</script>
