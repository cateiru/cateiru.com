import {Box, Button, Center, Heading} from '@chakra-ui/react';
import NextLink from 'next/link';
import useLanguage from '../useLanguage';

export const Link = () => {
  const {convertLang} = useLanguage();

  return (
    <Box mt="15vh" mb="10vh">
      <Heading textAlign="center" mb="1rem">
        {convertLang({ja: 'リンク', en: 'Links'})}
      </Heading>
      <Center>
        <NextLink passHref href="/">
          <Button>
            {convertLang({ja: 'トップへ戻る', en: 'Back to top'})}
          </Button>
        </NextLink>
      </Center>
    </Box>
  );
};
