import {
  Box,
  Button,
  Center,
  Flex,
  Heading,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  useDisclosure,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import useSWR from 'swr';
import {api} from '../../utils/api';
import {fetcher, SWRError} from '../../utils/swr';
import {Contact} from '../../utils/types';
import {Back} from '../Back';
import useLanguage from '../useLanguage';

export const ContactEdit = () => {
  const {convertLang} = useLanguage();
  const {data, error} = useSWR<Contact[], SWRError>('/user/contact', fetcher);
  const {isOpen, onClose, onOpen} = useDisclosure();
  const toast = useToast();

  const [select, setSelect] = React.useState<Contact>();

  const onSelect = (d: Contact) => {
    setSelect(d);
    onOpen();
  };

  const onDelete = () => {
    const f = async () => {
      if (!select) {
        return;
      }

      const res = await fetch(api(`/user/contact?contact_id=${select?.id}`), {
        method: 'DELETE',
        credentials: 'include',
        mode: 'cors',
      });

      if (res.ok) {
        toast({
          status: 'success',
          title: convertLang({ja: '削除しました', en: 'Success deleted'}),
        });
      } else {
        toast({
          status: 'error',
          title: (await res.json()).message,
        });
      }
    };
    f();
  };

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ja: 'リンク編集', en: 'Link Edit'})}
      </Heading>
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
        <Box
          mx={{base: '.5rem', sm: '1.5rem', md: '0'}}
          display={{base: 'block', md: 'flex'}}
          alignItems="center"
          flexDirection="column"
        >
          <Box width={{base: 'auto', md: '700px'}}>
            <Back />
            <TableContainer>
              <Table variant="simple">
                <Thead>
                  <Tr>
                    <Th>{convertLang({ja: '件名', en: 'Subject'})}</Th>
                    <Th>{convertLang({ja: '詳細', en: 'Detail'})}</Th>
                    <Th>
                      {convertLang({ja: 'メールアドレス', en: 'Email Address'})}
                    </Th>
                    <Th></Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map(v => {
                    return (
                      <Tr key={v.id}>
                        <Td>{v.title}</Td>
                        <Td>{v.detail}</Td>
                        <Td>{v.mail}</Td>
                        <Td>
                          <Button size="sm" onClick={() => onSelect(v)}>
                            {convertLang({ja: '詳細', en: 'Detail'})}
                          </Button>
                        </Td>
                      </Tr>
                    );
                  })}
                </Tbody>
              </Table>
            </TableContainer>
          </Box>
        </Box>
      )}
      <Modal isOpen={isOpen} onClose={onClose} size="xl">
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>
            {convertLang({ja: 'お問い合わせ詳細', en: 'Contact Detail'})}
          </ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody mb="1rem">
            <Flex justifyContent="right">
              <Button
                colorScheme="red"
                variant="outline"
                size="sm"
                onClick={onDelete}
              >
                {convertLang({ja: '削除', en: 'Delete'})}
              </Button>
            </Flex>
            <TableContainer>
              <Heading>
                {convertLang({
                  ja: 'お問い合わせ内容',
                  en: 'Contents of inquiry',
                })}
              </Heading>
              <Table variant="simple">
                <Tbody>
                  <Tr>
                    <Td fontWeight="bold">
                      {convertLang({ja: '件名', en: 'Subject'})}
                    </Td>
                    <Td>{select?.title}</Td>
                  </Tr>
                </Tbody>
              </Table>
            </TableContainer>
          </ModalBody>
        </ModalContent>
      </Modal>
    </Box>
  );
};
