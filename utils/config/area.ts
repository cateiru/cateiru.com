import type {MultiLang} from './lang';

export interface Area {
  country: MultiLang;
  prefecture: MultiLang;
}

export const tokyo: Area = {
  country: {
    en: 'japan',
    ja: '日本',
  },
  prefecture: {
    en: 'tokyo',
    ja: '東京都',
  },
};

export const kyoto: Area = {
  country: {
    en: 'japan',
    ja: '日本',
  },
  prefecture: {
    en: 'kyoto',
    ja: '京都府',
  },
};

export const ibaraki: Area = {
  country: {
    en: 'japan',
    ja: '日本',
  },
  prefecture: {
    en: 'ibaraki',
    ja: '茨城県',
  },
};

export const saitama: Area = {
  country: {
    en: 'japan',
    ja: '日本',
  },
  prefecture: {
    en: 'saitama',
    ja: '埼玉県',
  },
};
