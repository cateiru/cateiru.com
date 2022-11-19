import {
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
import {fetcher, SWRError} from '../../utils/swr';
import {Link} from '../../utils/types';
import useLanguage from '../useLanguage';
import {CardFrame} from './CardFrame';

export const LinkCard = () => {
  const [lang, convertLang] = useLanguage();

  const {data, error} = useSWR<Link[], SWRError>('/user/link', fetcher);

  return (
    <CardFrame>
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ja: 'リンク', en: 'Links'})}
      </Heading>
      <Center my=".5rem">
        <NextLink passHref href="/admin/link">
          <Button as="a" ml=".5rem">
            {convertLang({ja: '全て', en: 'All'})}
          </Button>
        </NextLink>
        <NextLink passHref href="/admin/category">
          <Button as="a" ml=".5rem">
            {convertLang({ja: 'カテゴリ', en: 'Categories'})}
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
                <Th>{convertLang({ja: '名前', en: 'Location'})}</Th>
              </Tr>
            </Thead>
            <Tbody>
              {data &&
                data.map(v => {
                  return (
                    <Tr key={v.id}>
                      <Td></Td>
                    </Tr>
                  );
                })}
            </Tbody>
          </Table>
        </TableContainer>
      )}
    </CardFrame>
  );
};
