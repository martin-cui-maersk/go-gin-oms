import type { FormInstance } from 'ant-design-vue/lib/form/Form';
import type {
  RuleObject,
  NamePath,
  Rule as ValidationRule,
} from 'ant-design-vue/lib/form/interface';
import { ref, computed, unref, Ref } from 'vue';
import { useI18n } from '@/hooks/web/useI18n';

export enum LoginStateEnum {
  LOGIN,
  RESET_PASSWORD,
  RESET_PASSWORD_SUBMIT,
}

const currentState = ref(LoginStateEnum.LOGIN);

// 这里也可以优化
// import { createGlobalState } from '@vueuse/core'

export function useLoginState() {
  function setLoginState(state: LoginStateEnum) {
    currentState.value = state;
  }

  const getLoginState = computed(() => currentState.value);

  function handleBackLogin() {
    setLoginState(LoginStateEnum.LOGIN);
  }

  return { setLoginState, getLoginState, handleBackLogin };
}

export function useFormValid<T extends Object = any>(formRef: Ref<FormInstance>) {
  const validate = computed(() => {
    const form = unref(formRef);
    return form?.validate ?? ((_nameList?: NamePath) => Promise.resolve());
  });

  async function validForm() {
    const form = unref(formRef);
    if (!form) return;
    const data = await form.validate();
    return data as T;
  }

  return { validate, validForm };
}

export function useFormRules(formData?: Recordable) {
  const { t } = useI18n();

  const getAccountFormRule = computed(() => createRule(t('basic.login.accountPlaceholder')));
  const getPasswordFormRule = computed(() => createRule(t('basic.login.passwordPlaceholder')));
  const getEmailCodeFormRule = computed(() => createRule(t('basic.login.emailCodePlaceholder')));
  const getEmailFormRule = computed(() => createRule(t('basic.login.emailPlaceholder')));
  const getEmailFormatFormRule = computed(() => createEmailRule(t('basic.login.emailFormatError')));

  const validateConfirmPassword = (password: string) => {
    return async (_: RuleObject, value: string) => {
      if (!value) {
        return Promise.reject(t('basic.login.passwordPlaceholder'));
      }
      if (value !== password) {
        return Promise.reject(t('basic.login.diffPwd'));
      }
      return Promise.resolve();
    };
  };

  const getFormRules = computed((): { [k: string]: ValidationRule | ValidationRule[] } => {
    const accountFormRule = unref(getAccountFormRule);
    const passwordFormRule = unref(getPasswordFormRule);
    const emailFormRule = unref(getEmailFormRule);
    const emailCodeFormRule = unref(getEmailCodeFormRule);
    const emailFormatFormRule = unref(getEmailFormatFormRule);

    const EmailRule = {
      // email: emailFormRule,
      code: emailCodeFormRule,
    };
    console.log('currentState', currentState);
    switch (unref(currentState)) {
      // submit reset password
      case LoginStateEnum.RESET_PASSWORD_SUBMIT:
        return {
          password: passwordFormRule,
          confirmPassword: [
            { validator: validateConfirmPassword(formData?.password), trigger: 'change' },
          ],
        };
      // reset password form rules
      case LoginStateEnum.RESET_PASSWORD:
        return {
          // email: accountFormRule,
          ...EmailRule,
          email: emailFormatFormRule,
        };

      // login form rules
      default:
        return {
          account: accountFormRule,
          password: passwordFormRule,
        };
    }
  });
  return { getFormRules };
}

function createRule(message: string): ValidationRule[] {
  return [
    {
      required: true,
      message,
      trigger: 'change',
    },
  ];
}

function createEmailRule(message: string): ValidationRule[] {
  return [
    {
      required: true,
      message,
      trigger: 'change',
      pattern: /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
    },
  ];
}
