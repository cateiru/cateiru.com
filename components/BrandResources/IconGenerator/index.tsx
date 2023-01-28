import {Box, Center, Heading, Text} from '@chakra-ui/react';
import useLanguage from '../../useLanguage';
import {Generator} from './Generator';

export const IconGenerator = () => {
  const {convertLang} = useLanguage();

  return (
    <Center>
      <Box w={{base: '96%', md: '500px'}} mb="5rem">
        <Heading textAlign="center" mt="3rem" mb="2rem">
          {convertLang({ja: 'アイコンジェネレーター', en: 'Icon Generator'})}
        </Heading>
        <Text textAlign="center" mb="1rem">
          {convertLang({
            ja: 'アイコンジェネレーターを使用すると、アイコンのサイズ、背景、円形のマスクなどが可能です。',
            en: 'The icon generator allows for icon size, background, circular masks, etc.',
          })}
        </Text>
        <Generator />
      </Box>
    </Center>
  );
};
