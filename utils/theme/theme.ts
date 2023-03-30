import {
  extendTheme,
  type ThemeConfig,
  defineStyleConfig,
} from '@chakra-ui/react';
import {ColorThemes} from '../types';

const config: ThemeConfig = {
  useSystemColorMode: false,
};

export const colorTheme: ColorThemes = {
  darkBackground: '#242838',
  lightBackground: '#ffffff',
  darkText: '#e8e8e8',
  lightText: '#1f1f1f',
};

const theme = extendTheme({
  colors: {
    brand: {
      200: '#E2E8F0',
      300: '#CBD5E0',
      500: '#404663',
      600: '#343952',
    },
  },

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
        width: '7px',
        height: '5px',
      },
      '&::-webkit-scrollbar-thumb': {
        backgroundColor: props.colorMode === 'dark' ? 'brand.600' : 'gray.400',
        borderRadius: '100px',
        ':hover': {
          backgroundColor:
            props.colorMode === 'dark' ? 'brand.500' : 'brand.500',
        },
      },
      '::-webkit-scrollbar-track': {
        backgroundColor: 'rgba(0,0,0,0)',
      },
      // FireFox
      html: {
        scrollbarWidth: 'thin',
        scrollbarColor: props.colorMode === 'dark' ? 'brand.600' : 'gray.400',
        overflow: 'overlay',
      },
      body: {
        width: '100vw',
        background:
          props.colorMode === 'dark'
            ? colorTheme.darkBackground
            : colorTheme.lightBackground,
        color:
          props.colorMode === 'dark'
            ? colorTheme.darkText
            : colorTheme.lightText,
      },
      ':root': {
        '--background-color':
          props.colorMode === 'dark'
            ? colorTheme.darkBackground
            : colorTheme.lightBackground,
        '--text-color':
          props.colorMode === 'dark'
            ? colorTheme.darkText
            : colorTheme.lightText,
      },
    }),
  },
  config: config,
});

export default theme;
