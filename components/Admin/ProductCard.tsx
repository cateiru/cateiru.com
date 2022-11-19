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
import {parseDate} from '../../utils/parse';
import {fetcher, SWRError} from '../../utils/swr';
import {ProductArray} from '../../utils/types';
import useLanguage from '../useLanguage';
import {CardFrame} from './CardFrame';

export const ProductCard = () => {
  const [lang, convertLang] = useLanguage();
  const {data, error} = useSWR<ProductArray, SWRError>(
    '/user/product',
    fetcher
  );

  return (
    <CardFrame>
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ja: '制作物', en: 'Products'})}
      </Heading>
      <Center my=".5rem">
        <NextLink passHref href="/admin/products">
          <Button as="a" ml=".5rem">
            {convertLang({ja: '全て', en: 'All'})}
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
          {error.message}
        </Center>
      ) : (
        <TableContainer>
          <Table variant="simple">
            <Thead>
              <Tr>
                <Th>{convertLang({ja: '名前', en: 'Name'})}</Th>
                <Th>{convertLang({ja: '詳細', en: 'Detail'})}</Th>
                <Th>{convertLang({ja: '開発日時', en: 'Dev Time'})}</Th>
                <Th></Th>
              </Tr>
            </Thead>
            <Tbody>
              {data &&
                data.map(v => {
                  return (
                    <Tr key={`prod-${v.id}`}>
                      <Td>{convertLang({ja: v.name_ja, en: v.name})}</Td>
                      <Td
                        whiteSpace="nowrap"
                        textOverflow="ellipsis"
                        overflow="hidden"
                        maxWidth="300px"
                      >
                        {convertLang({ja: v.detail_ja, en: v.detail})}
                      </Td>
                      <Td>{parseDate(v.dev_time, lang)}</Td>
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
