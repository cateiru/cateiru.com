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
import { parseAgo } from "../../utils/parse";
import { fetcher, type SWRError } from "../../utils/swr";
import type { Contact } from "../../utils/types";
import useLanguage from "../useLanguage";
import { CardFrame } from "./CardFrame";

export const ContactCard = () => {
  const { convertLang, lang } = useLanguage();

  const { data, error } = useSWR<Contact[], SWRError>("/user/contact", fetcher);

  return (
    <CardFrame>
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ ja: "お問い合わせ", en: "Contacts" })}
      </Heading>
      <Center my=".5rem">
        <NextLink passHref href="/admin/contact">
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
                <Th>{convertLang({ ja: "日時", en: "Time" })}</Th>
                <Th>{convertLang({ ja: "件名", en: "Subject" })}</Th>
              </Tr>
            </Thead>
            <Tbody>
              {data?.slice(0, 8).map((v) => {
                  return (
                    <Tr key={`link-${v.id}`}>
                      <Td>{parseAgo(v.created, lang)}</Td>
                      <Td>{convertLang({ ja: v.title, en: v.title })}</Td>
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
