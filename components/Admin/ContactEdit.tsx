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
  Text,
  Th,
  Thead,
  Tooltip,
  Tr,
  useClipboard,
  useColorMode,
  useDisclosure,
  useToast,
} from "@chakra-ui/react";
import React from "react";
import { TbCheck } from "react-icons/tb";
import useSWR from "swr";
import { api } from "../../utils/api";
import type { MultiLang } from "../../utils/config/lang";
import { copyElement, parseAgo, parseDetailDate } from "../../utils/parse";
import { fetcher, type SWRError } from "../../utils/swr";
import type { Contact } from "../../utils/types";
import { Back } from "../Back";
import useLanguage from "../useLanguage";

export const ContactEdit = () => {
  const { convertLang, lang } = useLanguage();
  const { data, error, mutate } = useSWR<Contact[], SWRError>(
    "/user/contact",
    fetcher,
  );
  const { isOpen, onClose, onOpen } = useDisclosure();
  const toast = useToast();
  const { colorMode } = useColorMode();

  const [select, setSelect] = React.useState<Contact>();

  const onSelect = (d: Contact) => {
    setSelect(d);
    onOpen();
  };

  const onDelete = () => {
    const f = async () => {
      if (!select || !data) {
        return;
      }

      const res = await fetch(api(`/user/contact?contact_id=${select.id}`), {
        method: "DELETE",
        credentials: "include",
        mode: "cors",
      });

      if (res.ok) {
        toast({
          status: "success",
          title: convertLang({ ja: "削除しました", en: "Success deleted" }),
        });
        const d = [...data];
        const i = d.findIndex((v) => v.id === select.id);
        mutate(d.slice(i, i + 1));
        onClose();
      } else {
        toast({
          status: "error",
          title: (await res.json()).message,
        });
      }
    };
    f();
  };

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ ja: "リンク編集", en: "Link Edit" })}
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
          mx={{ base: ".5rem", sm: "1.5rem", md: "0" }}
          display={{ base: "block", md: "flex" }}
          alignItems="center"
          flexDirection="column"
        >
          <Box width={{ base: "auto", md: "700px" }}>
            <Back href="/admin" />
            <TableContainer>
              <Table variant="simple">
                <Thead>
                  <Tr>
                    <Th>{convertLang({ ja: "日時", en: "Date" })}</Th>
                    <Th>{convertLang({ ja: "件名", en: "Subject" })}</Th>
                    <Th>
                      {convertLang({
                        ja: "メールアドレス",
                        en: "Email Address",
                      })}
                    </Th>
                    <Th></Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map((v) => {
                    return (
                      <Tr key={v.id}>
                        <Td>
                          <Tooltip
                            color={colorMode === "dark" ? "black" : "white"}
                            label={parseDetailDate(v.created)}
                            placement="top"
                            hasArrow
                            borderRadius="25"
                          >
                            {parseAgo(v.created, lang)}
                          </Tooltip>
                        </Td>
                        <Td>{v.title}</Td>
                        <Td>{v.mail}</Td>
                        <Td>
                          <Button size="sm" onClick={() => onSelect(v)}>
                            {convertLang({ ja: "詳細", en: "Detail" })}
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
            {convertLang({ ja: "お問い合わせ詳細", en: "Contact Detail" })}
          </ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody mb="1rem">
            {select && (
              <ContactElement
                convertLang={convertLang}
                onDelete={onDelete}
                select={select}
              />
            )}
          </ModalBody>
        </ModalContent>
      </Modal>
    </Box>
  );
};

const ContactElement: React.FC<{
  convertLang: (e: MultiLang) => string;
  onDelete: () => void;
  select: Contact;
}> = ({ convertLang, onDelete, select }) => {
  const { hasCopied, onCopy } = useClipboard(copyElement(select));

  return (
    <>
      <Flex justifyContent="right">
        <Button
          colorScheme="red"
          variant="outline"
          size="sm"
          onClick={onDelete}
        >
          {convertLang({ ja: "削除", en: "Delete" })}
        </Button>
      </Flex>
      <TableContainer>
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold">
                {convertLang({ ja: "件名", en: "Subject" })}
              </Td>
              <Td>{select?.title}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">
                {convertLang({ ja: "日時", en: "Date" })}
              </Td>
              <Td>{select?.created ? parseDetailDate(select?.created) : ""}</Td>
            </Tr>
            {select?.url && (
              <Tr>
                <Td fontWeight="bold">
                  {convertLang({ ja: "URL", en: "URL" })}
                </Td>
                <Td>{select?.url}</Td>
              </Tr>
            )}
            {select?.custom_title && (
              <Tr>
                <Td fontWeight="bold">{select?.custom_title}</Td>
                <Td>{select?.custom_value}</Td>
              </Tr>
            )}
          </Tbody>
        </Table>
      </TableContainer>
      <Text
        as="pre"
        fontFamily="'Kosugi Maru', sans-serif"
        fontSize="1rem"
        minH="120px"
        my="1rem"
        mx=".5rem"
        whiteSpace="break-spaces"
      >
        {select?.detail}
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold">
                {convertLang({ ja: "名前", en: "Name" })}
              </Td>
              <Td>{select?.name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">
                {convertLang({ ja: "メールアドレス", en: "Email" })}
              </Td>
              <Td>{select?.mail}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">{convertLang({ ja: "IP", en: "IP" })}</Td>
              <Td>{select?.ip}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">
                {convertLang({ ja: "言語", en: "Language" })}
              </Td>
              <Td>{select?.lang}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">
                {convertLang({ ja: "デバイス", en: "Device" })}
              </Td>
              <Td>{select?.device_name ?? "-"}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">{convertLang({ ja: "OS", en: "OS" })}</Td>
              <Td>{select?.os ?? "-"}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">
                {convertLang({ ja: "ブラウザ", en: "Browser" })}
              </Td>
              <Td>{select?.browser_name ?? "-"}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">
                {convertLang({ ja: "モバイル端末", en: "Mobile" })}
              </Td>
              <Td>{select?.is_mobile ? "true" : "false"}</Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
      <Button
        mt="1rem"
        onClick={onCopy}
        leftIcon={hasCopied ? <TbCheck size="20" /> : undefined}
        w="100%"
        size="sm"
      >
        {convertLang({ ja: "コピー", en: "Copy" })}
      </Button>
    </>
  );
};
