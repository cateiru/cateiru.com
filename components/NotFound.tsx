import {Box, Center, Heading, Text} from '@chakra-ui/react';
import useLanguage from './useLanguage';

export const NotFound = () => {
  const {convertLang} = useLanguage();

  return (
    <Center h="100vh">
      <Box
        backgroundImage="linear-gradient(135deg, #2bc4cf, #572bcf, #cf2ba1)"
        backgroundClip="text"
      >
        <Center mb="1rem">
          <Text fontSize="1.5rem" fontWeight="bold">
            ﾊﾞﾝﾊﾞﾝﾊﾞﾝﾊﾞﾝﾊﾞﾝﾊﾞﾝﾊﾞﾝﾞﾝﾊﾞﾝ
            <br />
            ﾊﾞﾝ&nbsp;&nbsp;&nbsp;∧＿∧&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ﾊﾞﾝﾊﾞﾝﾊﾞﾝ
            <br />
            ﾊﾞﾝ&nbsp;&nbsp;(∩`･ω･)&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ﾊﾞﾝﾊﾞﾝ
            <br />
            &nbsp;&nbsp;&nbsp;＿/_ﾐつ/￣￣￣/
            <br />
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;＼/＿＿＿/￣￣
            <br />
          </Text>
        </Center>
        <Heading textAlign="center">
          {convertLang({
            ja: '404 | お探しのページは見つかりません',
            en: '404 | Not Found',
          })}
        </Heading>
      </Box>
    </Center>
  );
};
