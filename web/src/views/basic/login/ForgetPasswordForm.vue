<template>
  <template v-if="getShow">
    <LoginFormTitle class="enter-x" />
    <Form class="p-4 enter-x" :model="formData" :rules="getFormRules" ref="formRef">
      <p>
        <span style="color: #0960bd; font-weight: bold">
          {{ t('basic.login.resetPasswordTip1') }}
        </span>
        &nbsp;-->&nbsp;
        <span>{{ t('basic.login.resetPasswordTip2') }}</span>
      </p>
      <FormItem name="email" class="enter-x">
        <Input
          size="large"
          v-model:value="formData.email"
          :placeholder="t('basic.login.email')"
          @blur="blurEmailError"
          @focus="blurEmailError"
        />
        <span style="color: #ed6f6f" v-if="errors.email">{{ errors.email }}</span>
      </FormItem>
      <FormItem name="code" class="enter-x">
        <CountdownInput
          size="large"
          v-model:value="formData.code"
          :placeholder="t('basic.login.emailCode')"
          :sendCodeApi="handleStart"
        />
      </FormItem>

      <FormItem class="enter-x">
        <Button type="primary" size="large" block @click="handleNext" :loading="loading">
          {{ t('global.nextText') }}
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
  import { CountdownInput } from '@/components/CountDown';
  import { useI18n } from '@/hooks/web/useI18n';
  import { useLoginState, useFormRules, LoginStateEnum, useFormValid } from './useLogin';
  import {
    getVerificationCodeForResetPassword,
    verifyVerificationCodeForResetPassword,
  } from '@/api/basic/user';

  const FormItem = Form.Item;
  const { t } = useI18n();
  const { handleBackLogin, getLoginState, setLoginState } = useLoginState();
  const { getFormRules } = useFormRules();

  const formRef = ref();
  const loading = ref(false);

  const formData = reactive({
    email: '',
    code: '',
  });

  const errors = reactive({
    email: '',
  });

  const { validForm } = useFormValid(formRef);

  const getShow = computed(() => unref(getLoginState) === LoginStateEnum.RESET_PASSWORD);

  function validateEmail(email) {
    const regExp = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return regExp.test(email);
  }

  function blurEmailError() {
    const form = unref(formRef);
    console.log('handleStart', formData.email, form.clearValidate());
    if (formData.email.trim() === '') {
      errors.email = 'Please input Email';
    } else if (!validateEmail(formData.email)) {
      errors.email = 'The email format is incorrect';
    } else {
      errors.email = '';
    }
  }

  async function handleNext() {
    errors.email = '';
    const form = unref(formRef);
    if (!form) return;
    const data = await validForm();
    if (!data) return;
    loading.value = true;
    verifyVerificationCodeForResetPassword(data).then(async (res) => {
      loading.value = false;
      console.log('verifyVerificationCodeForResetPassword', res);
      if (res) {
        await form.resetFields();
        sessionStorage.setItem('RESET_EMAIL', data.email);
        sessionStorage.setItem('RESET_CODE', data.code);
        setLoginState(LoginStateEnum.RESET_PASSWORD_SUBMIT);
      }
    });
  }

  async function handleStart() {
    const form = unref(formRef);
    console.log('handleStart', formData.email, form.clearValidate());
    if (formData.email.trim() === '') {
      errors.email = 'Please input Email';
    } else if (!validateEmail(formData.email)) {
      errors.email = 'The email format is incorrect';
    } else {
      errors.email = '';
    }
    if (errors.email) return false;
    return new Promise((resolve, reject) => {
      getVerificationCodeForResetPassword(formData).then((res) => {
        console.log('getVerificationCodeForResetPassword', res);
        if (res) {
          resolve(true);
        } else {
          reject(false);
        }
      });
    });
  }
</script>
