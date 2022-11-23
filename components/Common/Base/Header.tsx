import {Flex, IconButton, useColorMode, Text} from '@chakra-ui/react';
import React from 'react';
import {IoMoon, IoSunny} from 'react-icons/io5';
import {useRecoilState} from 'recoil';
import {toUnicode} from '../../../utils/parse';
import {langState} from '../../../utils/state/atoms';

const Header = React.memo(() => {
  const {colorMode, toggleColorMode} = useColorMode();
  const [lang, setLang] = useRecoilState(langState);

  React.useEffect(() => {
    const language =
      (window.navigator.languages && window.navigator.languages[0]) ||
      window.navigator.language;

    if (/^en\b/.test(language)) {
      // change english
      setLang('en');
    }
  }, []);

  const toggleLang = () => {
    if (lang === 'ja') {
      setLang('en');
    } else {
      setLang('ja');
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
          colorMode === 'dark' ? (
            <IoSunny size="20px" />
          ) : (
            <IoMoon size="20px" />
          )
        }
        onClick={toggleColorMode}
        variant="ghost"
        my=".5rem"
        mr=".25rem"
      />
      <IconButton
        aria-label="switch color mode"
        icon={<>{lang === 'ja' ? '🇺🇸' : '🇯🇵'}</>}
        onClick={toggleLang}
        variant="ghost"
        my=".5rem"
        mr=".5rem"
      />
    </Flex>
  );
});

Header.displayName = 'Header';

export default Header;
