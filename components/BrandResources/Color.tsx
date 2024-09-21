import {
  Box,
  Button,
  Center,
  Heading,
  SimpleGrid,
  Text,
  useClipboard,
  useColorMode,
} from "@chakra-ui/react";
import React from "react";
import { TbCheck } from "react-icons/tb";
import { MultiLang } from "../../utils/config/lang";
import { textColor } from "../../utils/theme/color";
import { colorTheme } from "../../utils/theme/theme";
import useLanguage from "../useLanguage";
import { ColorCircle } from "./CorloCircle";

export const Color = () => {
  const { convertLang } = useLanguage();

  return (
    <Box>
      <Center>
        <Heading textAlign="center" mb="1rem" id="colors">
          {convertLang({ ja: "カラーパレット", en: "Color Palettes" })}
        </Heading>
      </Center>
      <Center>
        <Box w={{ base: "96%", md: "700px" }}>
          <ColorCircle color1="#2bc4cf" color2="#572bcf" color3="#cf2ba1" />
          <ColorCircle color1="#17c9c9" color2="#336cff" />
          <ColorCircle color1="#e23d3d" color2="#ec44bd" />

          <SimpleGrid
            columns={{ base: 1, sm: 2, md: 3 }}
            spacing="3rem"
            mt="3rem"
            mx={{ base: "10%", md: 0 }}
          >
            <ColorTile
              color="#572bcf"
              title={{ ja: "プライマリ", en: "Primary" }}
            />
            <ColorTile
              color="#2bc4cf"
              title={{ ja: "セカンダリ", en: "Secondary" }}
            />
            <ColorTile
              color="#cf2ba1"
              title={{ ja: "アクセント", en: "Accent" }}
            />
            <ColorTile
              color={colorTheme.lightBackground}
              title={{ ja: "背景", en: "Background" }}
            />
            <ColorTile
              color={colorTheme.darkBackground}
              title={{ ja: "ダークテーマ背景", en: "DarkMode Background" }}
            />
            <ColorTile
              color={colorTheme.lightText}
              title={{ ja: "文字", en: "Text" }}
            />
            <ColorTile
              color={colorTheme.darkText}
              title={{ ja: "ダークテーマ文字", en: "Text" }}
            />
          </SimpleGrid>
        </Box>
      </Center>
    </Box>
  );
};

const ColorTile: React.FC<{
  color: string;
  title: MultiLang;
}> = (props) => {
  const { convertLang } = useLanguage();
  const { colorMode } = useColorMode();
  const { onCopy, hasCopied } = useClipboard(props.color);
  const [hasCopiedSlow, setHasCopiedSlow] = React.useState(false);

  React.useEffect(() => {
    if (hasCopied) {
      setHasCopiedSlow(true);
    }

    setTimeout(() => {
      setHasCopiedSlow(false);
    }, 1000);
  }, [hasCopied]);

  return (
    <Box>
      <Button
        w="100%"
        h="250px"
        fontSize="1.4rem"
        color={hasCopiedSlow ? textColor(props.color) : props.color}
        bgColor={props.color}
        _hover={{
          bgColor: props.color,
          color: textColor(props.color),
        }}
        boxShadow={
          colorMode === "light"
            ? "0px 0px 17px -5px #a0acc0"
            : "0px 0px 17px -5px #000"
        }
        borderRadius="10px"
        onClick={onCopy}
      >
        {hasCopied ? <TbCheck size="35px" /> : props.color}
      </Button>
      <Text textAlign="center" mt=".5rem" fontSize="1.2rem">
        {convertLang(props.title)}
      </Text>
    </Box>
  );
};
