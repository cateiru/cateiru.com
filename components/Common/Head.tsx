import { useColorMode } from "@chakra-ui/react";
import NextHead from "next/head";
import type React from "react";
import type { MultiLang } from "../../utils/config/lang";
import useLanguage from "../useLanguage";

interface Props {
  title: MultiLang;
}

export const Head: React.FC<Props> = (props) => {
  const { convertLang } = useLanguage();
  const { colorMode } = useColorMode();

  return (
    <NextHead>
      <title>{convertLang(props.title)}</title>
      <meta
        name="theme-color"
        content={colorMode === "dark" ? "#0e121c" : "#fff"}
      />
    </NextHead>
  );
};
