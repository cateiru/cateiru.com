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
} from "@chakra-ui/react";
import NextLink from "next/link";
import useSWR from "swr";
import { fetcher, type SWRError } from "../../utils/swr";
import type { ContactDefault } from "../../utils/types";
import useLanguage from "../useLanguage";
import { CardFrame } from "./CardFrame";

export const ContactDefaultCard = () => {
  const { convertLang } = useLanguage();
  const { data, error } = useSWR<ContactDefault[], SWRError>(
    "/user/contact/default",
    fetcher,
  );

  return (
    <CardFrame>
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ ja: "お問い合わせカスタム", en: "Custom Contact" })}
      </Heading>
      <Center my=".5rem">
        <NextLink passHref href="/admin/contact_default">
          <Button as="a" ml=".5rem">
            {convertLang({ ja: "全て", en: "All" })}
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
                <Th>{convertLang({ ja: "ID", en: "ID" })}</Th>
                <Th>{convertLang({ ja: "名前", en: "Name" })}</Th>
                <Th>{convertLang({ ja: "メール", en: "Email" })}</Th>
                <Th>{convertLang({ ja: "URL", en: "URL" })}</Th>
                <Th>{convertLang({ ja: "カテゴリ", en: "Category" })}</Th>
                <Th>
                  {convertLang({ ja: "カスタムタイトル", en: "Custom Title" })}
                </Th>
              </Tr>
            </Thead>
            <Tbody>
              {data &&
                data.map((v) => {
                  return (
                    <Tr key={`bio-${v.id}`}>
                      <Td>{v.id}</Td>
                      <Td>{v.name}</Td>
                      <Td>{v.email}</Td>
                      <Td>{v.url}</Td>
                      <Td>{v.category}</Td>
                      <Td>{v.custom_title}</Td>
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
