import {
  Box,
  Heading,
  ListItem,
  Skeleton,
  Text,
  UnorderedList,
  useColorMode,
} from '@chakra-ui/react';
import useSWR from 'swr';
import {fetcher, SWRError} from '../../utils/swr';
import {UserData} from '../../utils/types';
import useLanguage from '../useLanguage';

export const PreviewUserData = () => {
  const {colorMode} = useColorMode();
  const {convertLang} = useLanguage();
  const {data, error, isLoading} = useSWR<UserData, SWRError>(
    '/contact/preview',
    fetcher
  );

  if (isLoading) {
    return (
      <Skeleton
        w="100%"
        py="1rem"
        px=".5rem"
        mb="1.5rem"
        height="150px"
        borderRadius="5px"
      ></Skeleton>
    );
  }

  if (error) {
    return (
      <Box
        w="100%"
        py="1rem"
        px=".5rem"
        bgColor={colorMode === 'dark' ? '#5c2428' : 'red.100'}
        borderRadius="5px"
        borderLeftWidth="4px"
        borderLeftColor={colorMode === 'dark' ? 'red.400' : 'red.400'}
        mb="1.5rem"
        whiteSpace="pre-wrap"
        fontSize="1rem"
      >
        <Text fontSize="1rem" fontWeight="bold">
          {convertLang({
            ja: 'うまく送信できない場合があるかもしれません。',
            en: 'You may not be able to send the message properly.',
          })}
        </Text>
        <Text mt=".5rem">
          {convertLang({
            ja: 'サーバーとの接続を試みましたが、何らかの問題が発生したため失敗しました。もう一度ページをロードしてみてください。',
            en: 'Attempts to connect to the server have failed due to some problem. Please try loading the page again.',
          })}
        </Text>
        <Text mt="1rem">{error.message}</Text>
      </Box>
    );
  }

  return (
    <Box
      w="100%"
      py="1rem"
      px=".5rem"
      bgColor={colorMode === 'dark' ? 'gray.700' : 'gray.100'}
      borderRadius="5px"
      borderLeftWidth="4px"
      borderLeftColor={colorMode === 'dark' ? 'gray.500' : 'gray.400'}
      mb="1.5rem"
      whiteSpace="pre-wrap"
      fontSize="1rem"
    >
      <Text fontSize="1rem" fontWeight="bold">
        {convertLang({
          ja: '以下のユーザ情報も一緒に送信されます。',
          en: 'The following user information is also sent together.',
        })}
      </Text>
      <UnorderedList mt=".5rem">
        <ListItem>{`${convertLang({
          ja: 'ブラウザ',
          en: 'Browser',
        })}: ${data?.browser}`}</ListItem>
        <ListItem>{`${convertLang({
          ja: 'OS',
          en: 'OS',
        })}: ${data?.os}`}</ListItem>
        {data?.device !== '' && (
          <ListItem>{`${convertLang({
            ja: 'デバイス',
            en: 'Device',
          })}: ${data?.device}`}</ListItem>
        )}
        <ListItem>{`${convertLang({
          ja: 'モバイル',
          en: 'Mobile',
        })}: ${data?.is_mobile}`}</ListItem>
      </UnorderedList>
    </Box>
  );
};
