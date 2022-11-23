import NextHead from 'next/head';
import React from 'react';
import {MultiLang} from '../../utils/config/lang';
import useLanguage from '../useLanguage';

interface Props {
  title: MultiLang;
}

export const Head: React.FC<Props> = props => {
  const {convertLang} = useLanguage();

  return (
    <NextHead>
      <title>{convertLang(props.title)}</title>
    </NextHead>
  );
};
