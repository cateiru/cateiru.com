/**********************************************************
 * md text tags
 *
 * @author Yuto Watanabe <yuto.w51942@gmail.com>
 * @version 1.0.0
 *
 * Copyright (C) 2021 hello-slide
 **********************************************************/
import {Text, useColorMode} from '@chakra-ui/react';
import type {CodeComponent} from 'react-markdown/lib/ast-to-react';

export const Code: CodeComponent = ({children}) => {
  const {colorMode} = useColorMode();
  return (
    <Text
      as="span"
      fontWeight="medium"
      fontSize="1rem"
      backgroundColor={colorMode === 'dark' ? 'gray.600' : 'gray.100'}
      borderRadius="3px"
      padding="0rem .4rem 0rem .4rem"
    >
      {children}
    </Text>
  );
};
