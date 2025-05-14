<template>
  <template v-if="getShow">
    <LoginFormTitle class="enter-x" />
    <Form class="p-4 enter-x" :model="formData" :rules="getFormRules" ref="formRef">
      <p>
        <span>{{ t('basic.login.resetPasswordTip1') }}</span>
        &nbsp;-->&nbsp;
        <span style="color: #0960bd; font-weight: bold">
          {{ t('basic.login.resetPasswordTip2') }}
        </span>
      </p>
      <FormItem name="password" class="enter-x">
        <StrengthMeter
          size="large"
          v-model:value="formData.password"
          :placeholder="t('basic.login.password')"
        />
      </FormItem>
      <FormItem name="confirmPassword" class="enter-x">
        <InputPassword
          size="large"
          visibilityToggle
          v-model:value="formData.confirmPassword"
          :placeholder="t('basic.login.confirmPassword')"
        />
      </FormItem>

      <FormItem class="enter-x">
        <Button type="primary" size="large" block @click="handleSubmit" :loading="loading">
          {{ t('global.submitText') }}
        </Button>
        <Button size="large" block class="mt-4" @click="handleBackLogin">
          {{ t('basic.login.backSignIn') }}
        </Button>
      </FormItem>
    </Form>
  </template>
</template>
<script lang="ts" setup>
  import { reactive, ref, computed, unref } from 'vue';
  import LoginFormTitle from './LoginFormTitle.vue';
  import { Form, Input, Button } from 'ant-design-vue';
  import { StrengthMeter } from '@/components/StrengthMeter';
  import { useI18n } from '@/hooks/web/useI18n';
  import { useLoginState, useFormRules, LoginStateEnum, useFormValid } from './useLogin';
  import { hashStringWithSHA256, submitResetPassword } from '@/api/basic/user';

  const FormItem = Form.Item;
  const InputPassword = Input.Password;
  const { t } = useI18n();
  const { handleBackLogin, getLoginState, setLoginState } = useLoginState();

  const formRef = ref();
  const loading = ref(false);

  const formData = reactive({
    password: '',
    confirmPassword: '',
    code: '',
    email: '',
  });
  const { getFormRules } = useFormRules(formData);
  const { validForm } = useFormValid(formRef);

  const getShow = computed(() => unref(getLoginState) === LoginStateEnum.RESET_PASSWORD_SUBMIT);

  async function handleSubmit() {
    const data = await validForm();
    if (!data) return;
    console.log(data);
    const password = await hashStringWithSHA256(formData.password);
    const confirmPassword = await hashStringWithSHA256(formData.confirmPassword);
    console.log('handleStart', formData);
    loading.value = true;
    submitResetPassword({
      password: password,
      confirmPassword: confirmPassword,
      code: sessionStorage.getItem('RESET_CODE'),
      email: sessionStorage.getItem('RESET_EMAIL'),
    }).then((res) => {
      loading.value = false;
      console.log('submitResetPassword', res);
      setLoginState(LoginStateEnum.LOGIN);
      return res;
    });
  }
</script>
