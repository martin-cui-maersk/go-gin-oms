import type { DropMenu } from '../components/Dropdown';
import type { LocaleSetting, LocaleType } from '#/config';

export const LOCALE: { [key: string]: LocaleType } = {
  ZH_CN: 'zh_CN',
  EN_US: 'en',
};

export const localeSetting: LocaleSetting = {
  showPicker: true,
  // Locale
  locale: LOCALE.ZH_CN,
  // Default locale
  fallback: LOCALE.ZH_CN,
  // available Locales
  availableLocales: [LOCALE.ZH_CN, LOCALE.EN_US],
  // default time zone
  timeZone: 'SHANGHAI',
};

// locale list
export const localeList: DropMenu[] = [
  {
    text: '简体中文',
    event: LOCALE.ZH_CN,
  },
  {
    text: 'English',
    event: LOCALE.EN_US,
  },
];

// time zone list
export const timeZoneList: DropMenu[] = [
  {
    text: 'Asia/Shanghai',
    event: 'SHANGHAI',
  },
  {
    text: 'Asia/Tokyo',
    event: 'TOKYO',
  },
  {
    text: 'Asia/Jakarta',
    event: 'JAKARTA',
  },
  {
    text: 'Asia/Bangkok',
    event: 'BANGKOK',
  },
];
