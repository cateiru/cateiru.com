import {
  Box,
  Button,
  Flex,
  Heading,
  Spacer,
  Text,
  useToast,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import {useRouter} from 'next/router';
import React from 'react';
import {api} from '../../utils/api';
import useLanguage from '../useLanguage';

export const ControllerCard = () => {
  const [lang, convertLang] = useLanguage();
  const router = useRouter();
  const toast = useToast();

  const [load, setLoad] = React.useState(false);

  const logoutHandler = () => {
    const f = async () => {
      setLoad(true);
      const res = await fetch(api('/logout'), {
        credentials: 'include',
        mode: 'cors',
      });

      if (res.ok) {
        router.replace('/');
      } else {
        toast({
          status: 'error',
          title: (await res.json()).message,
        });
      }
      setLoad(false);
    };
    f();
  };

  return (
    <Box
      w="100%"
      minH="350px"
      boxShadow="0px 1px 26px -3px #a0acc0"
      borderRadius="56px"
      p="1rem"
      position="relative"
    >
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ja: 'コントローラ', en: 'COntroller'})}
      </Heading>
      <Box mt="2rem" mx="1rem">
        <Flex alignItems="center">
          <Text>{convertLang({ja: 'ログアウト', en: 'Logout'})}</Text>
          <Spacer />
          <Button onClick={logoutHandler} isLoading={load}>
            {convertLang({ja: 'ログアウト', en: 'Logout'})}
          </Button>
        </Flex>
        <Flex alignItems="center" mt="1rem">
          <Text>
            {convertLang({ja: 'ページを見に行く', en: 'Go to Top Page'})}
          </Text>
          <Spacer />
          <NextLink passHref href="/">
            <Button as="a">{convertLang({ja: 'トップ', en: 'Top'})}</Button>
          </NextLink>
        </Flex>
      </Box>
    </Box>
  );
};
