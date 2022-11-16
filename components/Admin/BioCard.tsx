import {
  Box,
  Button,
  Center,
  Heading,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import useSWR from 'swr';
import {parseDate} from '../../utils/parse';
import {fetcher, SWRError} from '../../utils/swr';
import {BioLocArray} from '../../utils/types';
import useLanguage from '../useLanguage';

export const BioCard = () => {
  const [lang, convertLang] = useLanguage();
  const {data, error} = useSWR<BioLocArray, SWRError>('/user/bio', fetcher);

  return (
    <Box
      w="100%"
      minH="350px"
      boxShadow="0px 1px 26px -3px #a0acc0"
      borderRadius="56px"
      p="1rem"
    >
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ja: '略歴', en: 'Biography'})}
      </Heading>
      <Center my=".5rem">
        <NextLink passHref href="/admin/bio">
          <Button as="a" ml=".5rem">
            {convertLang({ja: '全て', en: 'All'})}
          </Button>
        </NextLink>
        <NextLink passHref href="/admin/location">
          <Button as="a" ml=".5rem">
            {convertLang({ja: '場所', en: 'Locations'})}
          </Button>
        </NextLink>
      </Center>
      {error ? (
        <Center
          fontSize="1.5rem"
          fontWeight="bold"
          color="red.400"
          textAlign="center"
        >
          {error.status} : {error.message}
        </Center>
      ) : (
        <TableContainer>
          <Table variant="simple">
            <Thead>
              <Tr>
                <Th>{convertLang({ja: '場所', en: 'Location'})}</Th>
                <Th>{convertLang({ja: 'ポジション', en: 'Position'})}</Th>
                <Th>{convertLang({ja: '参加', en: 'Join'})}</Th>
                <Th>{convertLang({ja: '退場', en: 'Leave'})}</Th>
              </Tr>
            </Thead>
            <Tbody>
              {data &&
                data.map(v => {
                  return (
                    <Tr key={v.biography.id}>
                      <Td fontWeight="bold">
                        {convertLang({
                          ja: v.location.name_ja,
                          en: v.location.name,
                        })}
                      </Td>
                      <Td>
                        {convertLang({
                          ja: v.biography.position_ja,
                          en: v.biography.position,
                        })}
                      </Td>
                      <Td>{parseDate(v.biography.join, lang)}</Td>
                      <Td>{parseDate(v.biography.leave, lang)}</Td>
                    </Tr>
                  );
                })}
            </Tbody>
          </Table>
        </TableContainer>
      )}
    </Box>
  );
};
