import { Flex, IconButton, useColorMode } from "@chakra-ui/react";
import React from "react";
import { IoMoon, IoSunny } from "react-icons/io5";
import { useRecoilState } from "recoil";
import { langState } from "../../../utils/state/atoms";

const Header = React.memo(() => {
  const { colorMode, toggleColorMode } = useColorMode();
  const [lang, setLang] = useRecoilState(langState);

  React.useEffect(() => {
    const language =
      (window.navigator.languages && window.navigator.languages[0]) ||
      window.navigator.language;

    if (!/^ja\b/.test(language)) {
      // change english
      document.querySelector("html")?.setAttribute("lang", "en");
      setLang("en");
    }
  }, []);

  const toggleLang = () => {
    const htmlElement = document.querySelector("html");

    if (lang === "ja") {
      htmlElement?.setAttribute("lang", "en");
      setLang("en");
    } else {
      htmlElement?.setAttribute("lang", "ja");
      setLang("ja");
    }
  };

  return (
    <Flex
      width="100%"
      justifyContent="end"
      position="fixed"
      top={0}
      left={0}
      zIndex="1000"
    >
      <IconButton
        aria-label="switch color mode"
        icon={
          colorMode === "dark" ? (
            <IoSunny size="20px" />
          ) : (
            <IoMoon size="20px" />
          )
        }
        onClick={toggleColorMode}
        variant="ghost"
        my=".5rem"
        mr=".25rem"
        borderRadius="25"
      />
      <IconButton
        aria-label="switch color mode"
        icon={<>{lang === "ja" ? "🇺🇸" : "🇯🇵"}</>}
        onClick={toggleLang}
        variant="ghost"
        my=".5rem"
        mr=".5rem"
        borderRadius="25"
      />
    </Flex>
  );
});

Header.displayName = "Header";

export default Header;
