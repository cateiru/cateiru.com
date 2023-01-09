import {
  extendTheme,
  type ThemeConfig,
  defineStyleConfig,
} from '@chakra-ui/react';

const config: ThemeConfig = {
  useSystemColorMode: false,
};

const theme = extendTheme({
  fonts: {
    heading: "'Kosugi Maru', sans-serif",
    body: "'Kosugi Maru', sans-serif",
  },
  components: {
    CloseButton: defineStyleConfig({
      baseStyle: {
        _focus: {
          boxShadow: 'none',
        },
      },
    }),
    Text: defineStyleConfig({
      baseStyle: {
        fontWeight: 'normal',
      },
    }),
    Heading: defineStyleConfig({
      baseStyle: {
        fontWeight: '800',
      },
    }),
  },
  styles: {
    global: (props: {colorMode: string}) => ({
      // Chrome
      '&::-webkit-scrollbar': {
        width: '5px',
        height: '5px',
      },
      '&::-webkit-scrollbar-thumb': {
        backgroundColor: props.colorMode === 'dark' ? 'gray.600' : 'gray.300',
        borderRadius: '100px',
      },
      // FireFox
      html: {
        scrollbarWidth: 'thin',
        scrollbarColor: props.colorMode === 'dark' ? 'gray.600' : 'gray.300',
      },
    }),
  },
  config: config,
});

export default theme;
