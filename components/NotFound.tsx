import {Center, Heading} from '@chakra-ui/react';
import useLanguage from './useLanguage';

export const NotFound = () => {
  const {convertLang} = useLanguage();

  return (
    <Center h="100vh">
      <Heading
        backgroundImage="linear-gradient(135deg, #2bc4cf, #572bcf, #cf2ba1)"
        backgroundClip="text"
        textAlign="center"
      >
        {convertLang({
          ja: '404 | お探しのページは見つかりません',
          en: '404 | Not Found',
        })}
      </Heading>
    </Center>
  );
};
