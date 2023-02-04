import {
  Box,
  Button,
  Center,
  Heading,
  SimpleGrid,
  Text,
  useClipboard,
  useColorMode,
} from '@chakra-ui/react';
import React from 'react';
import {TbCheck} from 'react-icons/tb';
import {MultiLang} from '../../utils/config/lang';
import {textColor} from '../../utils/theme/color';
import useLanguage from '../useLanguage';

export const Color = () => {
  const {convertLang} = useLanguage();

  return (
    <Box>
      <Heading textAlign="center" mb="1rem" id="colors">
        {convertLang({ja: 'カラーパレット', en: 'Color Palettes'})}
      </Heading>
      <Center>
        <Box w={{base: '96%', md: '700px'}}>
          <Box
            h="20px"
            background="linear-gradient(90deg, #2bc4cf, #572bcf, #cf2ba1)"
            borderRadius="25px"
          ></Box>
          <SimpleGrid columns={3} spacing={10} mt=".5rem" fontSize="1.2rem">
            <Box textAlign="left">
              <ColorCode>#2bc4cf</ColorCode>
            </Box>
            <Box textAlign="center">
              <ColorCode>#572bcf</ColorCode>
            </Box>
            <Box textAlign="right">
              <ColorCode>#cf2ba1</ColorCode>
            </Box>
          </SimpleGrid>
          <SimpleGrid
            columns={{base: 1, sm: 2, md: 3}}
            spacing="3rem"
            mt="3rem"
            mx={{base: '10%', md: 0}}
          >
            <ColorTile
              color="#2bc4cf"
              title={{ja: 'プライマリ', en: 'Primary'}}
            />
            <ColorTile
              color="#cf2ba1"
              title={{ja: 'セカンダリ', en: 'Secondary'}}
            />
            <ColorTile color="#ffffff" title={{ja: '背景', en: 'Background'}} />
            <ColorTile
              color="#1a202c"
              title={{ja: 'ダークモード背景', en: 'DarkMode Background'}}
            />
            <ColorTile color="#0f0f0f" title={{ja: '文字', en: 'Text'}} />
          </SimpleGrid>
        </Box>
      </Center>
    </Box>
  );
};

const ColorCode: React.FC<{children: string}> = props => {
  const {onCopy, hasCopied} = useClipboard(props.children);

  return (
    <Button variant="ghost" color={props.children} onClick={onCopy}>
      {hasCopied ? (
        <TbCheck color={props.children} size="25px" />
      ) : (
        props.children
      )}
    </Button>
  );
};

const ColorTile: React.FC<{
  color: string;
  title: MultiLang;
}> = props => {
  const {convertLang} = useLanguage();
  const {colorMode} = useColorMode();
  const {onCopy, hasCopied} = useClipboard(props.color);
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
          colorMode === 'light'
            ? '0px 0px 17px -5px #a0acc0'
            : '0px 0px 17px -5px #000'
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
