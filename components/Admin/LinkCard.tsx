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
import { fetcher, SWRError } from "../../utils/swr";
import { LinkCategory } from "../../utils/types";
import useLanguage from "../useLanguage";
import { CardFrame } from "./CardFrame";

export const LinkCard = () => {
  const { convertLang } = useLanguage();

  const { data, error } = useSWR<LinkCategory[], SWRError>(
    "/user/link",
    fetcher,
  );

  return (
    <CardFrame>
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ ja: "リンク", en: "Links" })}
      </Heading>
      <Center my=".5rem">
        <NextLink passHref href="/admin/link">
          <Button as="a" ml=".5rem">
            {convertLang({ ja: "全て", en: "All" })}
          </Button>
        </NextLink>
        <NextLink passHref href="/admin/category">
          <Button as="a" ml=".5rem">
            {convertLang({ ja: "カテゴリ", en: "Categories" })}
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
                <Th>{convertLang({ ja: "名前", en: "Name" })}</Th>
                <Th>{convertLang({ ja: "カテゴリ", en: "Category" })}</Th>
              </Tr>
            </Thead>
            <Tbody>
              {data &&
                data.slice(0, 5).map((v) => {
                  return (
                    <Tr key={`link-${v.link.id}`}>
                      <Td>
                        {convertLang({ ja: v.link.name_ja, en: v.link.name })}
                      </Td>
                      <Td>
                        {convertLang({
                          ja: v.category.name_ja,
                          en: v.category.name,
                        })}
                      </Td>
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
