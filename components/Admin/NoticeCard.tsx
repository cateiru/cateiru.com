import {
  Badge,
  Button,
  Center,
  Heading,
  Table,
  TableContainer,
  Tbody,
  Td,
  Tr,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import useSWR from 'swr';
import {MultiLang} from '../../utils/config/lang';
import {fetcher, SWRError} from '../../utils/swr';
import {Notice} from '../../utils/types';
import useLanguage from '../useLanguage';
import {CardFrame} from './CardFrame';

export const NoticeCard = () => {
  const {convertLang} = useLanguage();

  const {data, error} = useSWR<Notice, SWRError>('/user/notice', fetcher);

  return (
    <CardFrame>
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ja: '通知先', en: 'Notice destination'})}
      </Heading>
      <Center mt=".5rem">
        <NextLink passHref href="/admin/notice">
          <Button as="a" ml=".5rem">
            {convertLang({ja: '設定', en: 'Edit'})}
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
            <Tbody>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: 'Discord', en: 'Discord'})}
                </Td>
                <Td>
                  {typeof data?.discord_webhook === 'string' ? (
                    <OK convertLang={convertLang} />
                  ) : (
                    <NO convertLang={convertLang} />
                  )}
                </Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: 'Slack', en: 'Slack'})}
                </Td>
                <Td>
                  {data?.slack_webhook ? (
                    <OK convertLang={convertLang} />
                  ) : (
                    <NO convertLang={convertLang} />
                  )}
                </Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ja: 'Email', en: 'Email'})}
                </Td>
                <Td>
                  {data?.mail ? (
                    <OK convertLang={convertLang} />
                  ) : (
                    <NO convertLang={convertLang} />
                  )}
                </Td>
              </Tr>
            </Tbody>
          </Table>
        </TableContainer>
      )}
    </CardFrame>
  );
};

const OK: React.FC<{convertLang: (e: MultiLang) => string}> = ({
  convertLang,
}) => {
  return (
    <Badge colorScheme="green">
      {convertLang({ja: '接続済み', en: 'Connected'})}
    </Badge>
  );
};

const NO: React.FC<{convertLang: (e: MultiLang) => string}> = ({
  convertLang,
}) => {
  return (
    <Badge colorScheme="red">
      {convertLang({ja: '未接続', en: 'No Connect'})}
    </Badge>
  );
};
