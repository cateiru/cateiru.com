import {
  Avatar,
  Box,
  Center,
  Heading,
  Text,
  Table,
  Tbody,
  Tr,
  Td,
  TableContainer,
  Badge,
  Button,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import {useRecoilValue} from 'recoil';
import {parseDate} from '../../utils/parse';
import {UserState} from '../../utils/state/atoms';
import useLanguage from '../useLanguage';

export const UserCard = () => {
  const [lang, convertLang] = useLanguage();
  const user = useRecoilValue(UserState);

  return (
    <Box
      w="100%"
      minH="350px"
      boxShadow="0px 1px 26px -3px #a0acc0"
      borderRadius="56px"
      p="1rem"
      position="relative"
    >
      <Box position="absolute" top="20" right="20">
        <NextLink passHref href="/admin/user">
          <Button size="sm" as="a">
            {convertLang({ja: '編集', en: 'edit'})}
          </Button>
        </NextLink>
      </Box>
      {user?.selected && (
        <Box position="absolute" top="20" left="20">
          <Badge variant="subtle" colorScheme="green">
            {convertLang({ja: '選択済み', en: 'selected'})}
          </Badge>
        </Box>
      )}
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ja: 'ユーザ', en: 'User Card'})}
      </Heading>
      <Box mt="1rem">
        <Center>
          <Avatar
            src={user?.avatar_url}
            size="lg"
            boxShadow="0px 1px 26px -3px #a0acc0"
          />
        </Center>
        <Text textAlign="center" mt=".5rem">
          id: {user?.id}
        </Text>
        <TableContainer>
          <Table variant="simple">
            <Tbody>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: 'ユーザID', en: 'User ID'})}
                </Td>
                <Td>{user?.user_id}</Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: '名字', en: 'Family Name'})}
                </Td>
                <Td>
                  {lang === 'ja' ? user?.family_name_ja : user?.family_name}
                </Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: '名前', en: 'Given Name'})}
                </Td>
                <Td>
                  {lang === 'ja' ? user?.given_name_ja : user?.given_name}
                </Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: 'メールアドレス', en: 'Mail Address'})}
                </Td>
                <Td>{user?.mail}</Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: '誕生日', en: 'Birth Date'})}
                </Td>
                <Td>{parseDate(user?.birth_date ?? '', lang)}</Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: '在住', en: 'Location'})}
                </Td>
                <Td>{lang === 'ja' ? user?.location_ja : user?.location}</Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: '作成日時', en: 'Created'})}
                </Td>
                <Td>{parseDate(user?.created ?? '', lang)}</Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: '編集日時', en: 'Modified'})}
                </Td>
                <Td>{parseDate(user?.modified ?? '', lang)}</Td>
              </Tr>
            </Tbody>
          </Table>
        </TableContainer>
      </Box>
    </Box>
  );
};
