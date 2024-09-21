import {
  Box,
  Button,
  Flex,
  Heading,
  Spacer,
  Text,
  useToast,
} from "@chakra-ui/react";
import NextLink from "next/link";
import { useRouter } from "next/router";
import React from "react";
import { api } from "../../utils/api";
import useLanguage from "../useLanguage";
import { CardFrame } from "./CardFrame";

export const ControllerCard = () => {
  const { convertLang } = useLanguage();
  const router = useRouter();
  const toast = useToast();

  const [load, setLoad] = React.useState(false);

  const logoutHandler = () => {
    const f = async () => {
      setLoad(true);
      const res = await fetch(api("/logout"), {
        credentials: "include",
        mode: "cors",
      });

      if (res.ok) {
        router.replace("/");
      } else {
        toast({
          status: "error",
          title: (await res.json()).message,
        });
      }
      setLoad(false);
    };
    f();
  };

  return (
    <CardFrame>
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ ja: "コントローラ", en: "Controller" })}
      </Heading>
      <Box mt="2rem" mx="1rem">
        <Flex alignItems="center">
          <Text>
            {convertLang({ ja: "ページを見に行く", en: "Go to Top Page" })}
          </Text>
          <Spacer />
          <NextLink passHref href="/">
            <Button as="a">{convertLang({ ja: "トップ", en: "Top" })}</Button>
          </NextLink>
        </Flex>
        <Flex alignItems="center" mt="1rem">
          <Text>
            {convertLang({ ja: "お問い合わせビルダー", en: "Contact Builder" })}
          </Text>
          <Spacer />
          <NextLink passHref href="/admin/contact_builder">
            <Button as="a">
              {convertLang({ ja: "ビルダー", en: "Builder" })}
            </Button>
          </NextLink>
        </Flex>
        <Flex alignItems="center" mt="1rem">
          <Text>
            {convertLang({ ja: "ブランドリソース", en: "Brand Resources" })}
          </Text>
          <Spacer />
          <NextLink passHref href="/brand_resources">
            <Button as="a">
              {convertLang({ ja: "ブランド", en: "Brand" })}
            </Button>
          </NextLink>
        </Flex>
        <Flex alignItems="center" mt="1rem">
          <Text>{convertLang({ ja: "ログアウト", en: "Logout" })}</Text>
          <Spacer />
          <Button onClick={logoutHandler} isLoading={load} colorScheme="red">
            {convertLang({ ja: "ログアウト", en: "Logout" })}
          </Button>
        </Flex>
      </Box>
    </CardFrame>
  );
};
