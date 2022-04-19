import React from 'react';
import {useRecoilValue} from 'recoil';
import type {MultiLang} from '../utils/config/lang';
import {langState} from '../utils/state/atoms';

const useLanguage = (): [string, (e: MultiLang) => string] => {
  const lang = useRecoilValue(langState);

  const convertLang = React.useCallback(
    (e: MultiLang) => {
      if (lang === 'ja') {
        return e.ja;
      }
      return e.en;
    },
    [lang]
  );

  return [lang, convertLang];
};

export default useLanguage;
