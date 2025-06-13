<template>
  <LoginFormTitle v-show="getShow" class="enter-x" />
  <Form
    class="p-4 enter-x"
    :model="formData"
    :rules="getFormRules"
    ref="formRef"
    v-show="getShow"
    @keypress.enter="handleLogin"
  >
    <FormItem name="account" class="enter-x">
      <Input
        size="large"
        v-model:value="formData.account"
        :placeholder="t('basic.login.userName')"
        class="fix-auto-fill"
      />
    </FormItem>
    <FormItem name="password" class="enter-x">
      <InputPassword
        size="large"
        visibilityToggle
        v-model:value="formData.password"
        :placeholder="t('basic.login.password')"
      />
    </FormItem>
    <FormItem class="enter-x">
      <Button type="primary" size="large" block @click="handleLogin" :loading="loading">
        {{ t('basic.login.loginButton') }}
      </Button>
    </FormItem>
    <ARow class="enter-x" :gutter="[16, 16]">
      <ACol :md="7" :xs="24" />
      <ACol :md="10" :xs="24" style="text-align: center">
        <a @click="setLoginState(LoginStateEnum.RESET_PASSWORD)">
          {{ t('basic.login.forgetPassword') }}
        </a>
      </ACol>
      <ACol :md="7" :xs="24" />
    </ARow>
  </Form>
  <BasicModal
    @register="register"
    title="旋转校验"
    width="400px"
    :footer="null"
    :canFullscreen="false"
    :useWrapper="false"
    :destroyOnClose="true"
    :maskStyle="{ backgroundColor: 'rgb(0, 0, 0, 0.55)' }"
    size="small"
  >
    <div>
      <RotateDragVerify :src="verify" :diffDegree="10" ref="el" @success="verifyLogin" />
    </div>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { reactive, ref, unref, computed } from 'vue';

  import { Form, Input, Row, Col, Button } from 'ant-design-vue';
  import LoginFormTitle from './LoginFormTitle.vue';

  import { useI18n } from '@/hooks/web/useI18n';
  import { useMessage } from '@/hooks/web/useMessage';

  import { useUserStore } from '@/store/modules/user';
  import { LoginStateEnum, useLoginState, useFormRules, useFormValid } from './useLogin';
  import { useDesign } from '@/hooks/web/useDesign';
  import { hashStringWithSHA256 } from '@/api/basic/user';
  //import { onKeyStroke } from '@vueuse/core';
  import { useModal, BasicModal } from '@/components/Modal';
  import { RotateDragVerify } from '@/components/Verify';
  import verify from '@/assets/images/verify.jpg';

  const ACol = Col;
  const ARow = Row;
  const FormItem = Form.Item;
  const InputPassword = Input.Password;
  const { t } = useI18n();
  const { notification, createErrorModal } = useMessage();
  const { prefixCls } = useDesign('login');
  const userStore = useUserStore();

  const { setLoginState, getLoginState } = useLoginState();
  const { getFormRules } = useFormRules();

  const formRef = ref();
  const loading = ref(false);
  const [register, { openModal, closeModal }] = useModal();
  // 自动填充密码
  // const formData = reactive({
  //   account: 'vben',
  //   password: '123456',
  // });
  const formData = reactive({});

  const { validForm } = useFormValid(formRef);

  //onKeyStroke('Enter', handleLogin);

  const getShow = computed(() => unref(getLoginState) === LoginStateEnum.LOGIN);

  async function verifyLogin() {
    closeModal();
    await handleLoginRequest();
  }

  async function handleLogin() {
    const data = await validForm();
    console.log('handleLogin', data);
    if (!data) return;
    openModal();
  }

  async function handleLoginRequest() {
    const data = await validForm();
    console.log('handleLogin', data);
    if (!data) return;
    try {
      // hash-256 加密密码
      data.password = await hashStringWithSHA256(data.password);
      loading.value = true;
      const userInfo = await userStore.login({
        password: data.password,
        account: data.account,
        mode: 'message', //不要默认的错误提示
      });
      if (userInfo) {
        notification.success({
          message: t('basic.login.loginSuccessTitle'),
          description: `${t('basic.login.loginSuccessDesc')}: ${userInfo.userName}`,
          duration: 3,
        });
      }
    } catch (error) {
      createErrorModal({
        title: t('basic.api.errorTip'),
        content: (error as unknown as Error).message || t('basic.api.networkExceptionMsg'),
        getContainer: () => document.body.querySelector(`.${prefixCls}`) || document.body,
      });
    } finally {
      loading.value = false;
    }
  }
</script>
