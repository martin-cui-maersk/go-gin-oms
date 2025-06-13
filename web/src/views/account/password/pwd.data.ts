import { FormSchema } from '@/components/Form';
import { useI18n } from '@/hooks/web/useI18n';

const { t } = useI18n();

export const formSchema: FormSchema[] = [
  {
    field: 'passwordOld',
    label: t('account.oldPassword'),
    component: 'InputPassword',
    componentProps: {
      placeholder: t('account.oldPassword'),
    },
    rules: [
      {
        required: true,
        message: t('account.oldPasswordTip'),
      },
    ],
  },
  {
    field: 'passwordNew',
    label: t('account.newPassword'),
    component: 'StrengthMeter',
    componentProps: {
      placeholder: t('account.newPassword'),
    },
    rules: [
      {
        required: true,
        message: t('account.newPasswordTip'),
      },
    ],
  },
  {
    field: 'confirmPassword',
    label: t('account.newPasswordConfirm'),
    component: 'InputPassword',

    dynamicRules: ({ values }) => {
      return [
        {
          required: true,
          validator: (_, value) => {
            if (!value) {
              return Promise.reject(t('account.passwordRequired'));
            }
            if (value !== values.passwordNew) {
              return Promise.reject(t('account.passwordNotMatch'));
            }
            return Promise.resolve();
          },
        },
      ];
    },
  },
];
