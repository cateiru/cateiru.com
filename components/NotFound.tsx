import {Center, Heading} from '@chakra-ui/react';
import useLanguage from './useLanguage';

export const NotFound = () => {
  const [lang, convertLang] = useLanguage();

  return (
    <Center h="100vh">
      <Heading>
        {convertLang({
          ja: '404 | お探しのページは見つかりません',
          en: '404 | Not Found',
        })}
      </Heading>
    </Center>
  );
};
