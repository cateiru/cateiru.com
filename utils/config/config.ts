import {tokyo, kyoto, ibaraki, saitama} from './area';
import type {Area} from './area';
import type {Biography} from './bio';
import type {MultiLang} from './lang';
import type {Product} from './product';
import type {SocialLinks} from './social';

export interface Config {
  // My profile
  name: MultiLang;
  nameKana: string;

  birthday: Date;
  avatar: string;
  birthplace: Area;
  residence: Area;

  bio: Biography;

  socialLinks: SocialLinks;
  products: Product[];
}

const config: Config = {
  name: {
    ja: '渡邊 悠人',
    en: 'Yuto Watanabe',
  },
  nameKana: 'わたなべ ゆうと',

  birthday: new Date(2000, 10, 1),
  avatar:
    'https://storage.googleapis.com/cateiru-sso/avatar/9f997a49b29b8f2971433690cc0ad8', // from CateiruSSO admin
  birthplace: ibaraki,
  residence: saitama,

  bio: {
    university: {
      name: {
        ja: '東京電機大学',
        en: 'Tokyo Denki University',
      },
      area: tokyo,

      admission: new Date(2019, 4),
      graduation: new Date(2023, 3),
    },
    works: [
      {
        name: {
          ja: '株式会社はてな',
          en: 'Hatena Co., Ltd.',
        },
        area: kyoto,

        admission: new Date(2023, 4),
        occupation: {
          en: 'Web Application Engineer',
          ja: 'Webアプリケーションエンジニア',
        },
      },
    ],
  },

  socialLinks: {
    github: 'https://github.com/cateiru',
    gitlab: 'https://gitlab.com/cateiru',
    bitbucket: 'https://bitbucket.org/cateiru',

    twitter: 'https://twitter.com/cateiru',
    facebook: 'https://www.facebook.com/yuto51942',
    instagram: 'https://www.instagram.com/yuto51942',
    discord: 'Cateiru#9825',

    niconico: 'https://www.nicovideo.jp/user/56247011',

    hatenaBlog: 'https://blog.cateiru.com/',
    qiita: 'https://qiita.com/cateiru',
    zenn: 'https://zenn.dev/cateiru',
    note: 'https://note.com/cateiru',

    hatena: 'https://profile.hatena.ne.jp/cateiru/',
  },

  products: [
    {
      name: {
        en: "find!",
        ja: "ふぁいんど！"
      },
      createDate: new Date(2022, 5),
      link: "https://find.cateiru.com",
      description: {
        en: "Web application to find a meeting partner with ease.",
        ja: "待ち合わせ相手が楽に見つかるWebアプリ",
      }
    },
    {
      name: {
        en: 'Cateiru SSO',
        ja: 'Cateiru SSO',
      },
      createDate: new Date(2022, 1),
      link: 'https://sso.cateiru.com',
      description: {
        en: 'My single sign on service.',
        ja: 'cateiruのシングル・サインオンサービス',
      },
    },
    {
      name: {
        en: 'Hello Slide',
        ja: 'Hello Slide',
      },
      createDate: new Date(2021, 8),
      link: 'https://hello-slide.jp',
      description: {
        en: 'Evolved slides that allow participation.',
        ja: '参加ができる進化したスライド',
      },
    },
    {
      name: {
        en: 'Earthquake Alert',
        ja: 'Earthquake Alert',
      },
      createDate: new Date(2019, 6),
      link: 'https://twitter.com/e_alert_',
      description: {
        en: 'Bot that sends earthquake information with images.',
        ja: '地震情報を画像付きで発信するBot',
      },
    },
  ],
};

export default config;
