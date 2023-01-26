import {
  Box,
  Button,
  Center,
  Heading,
  Input,
  InputGroup,
  InputRightElement,
  Text,
  useClipboard,
} from '@chakra-ui/react';
import useLanguage from '../useLanguage';

export const Copyright = () => {
  const now = new Date();

  const {convertLang} = useLanguage();
  const copycopyright = useClipboard(`© ${now.getFullYear()} cateiru`);
  const copycopyright1 = useClipboard(
    `Copyright (C) ${now.getFullYear()} cateiru`
  );

  return (
    <Box mt="15vh">
      <Heading textAlign="center" mb="1rem" id="copyright">
        {convertLang({ja: 'コピーライト表記', en: 'Copyright Statement'})}
      </Heading>
      <Text textAlign="center" mb="1.5rem" mx=".5rem">
        {convertLang({
          ja: 'Cateiruが作成したソフトウェアをライセンス下のもとに使用する場合などにご使用ください。',
          en: 'Please use this when using software created by Cateiru.',
        })}
      </Text>
      <Center mb=".5rem">
        <InputGroup size="md" w={{base: '96%', sm: '400px'}}>
          <Input
            pr="4.5rem"
            value={copycopyright.value}
            onFocus={e => e.target.select()}
            onChange={() => {}}
            border="0px"
          />
          <InputRightElement width="4.5rem">
            <Button h="1.75rem" size="sm" onClick={copycopyright.onCopy}>
              {copycopyright.hasCopied
                ? convertLang({ja: 'OK', en: 'Copied'})
                : convertLang({ja: 'コピー', en: 'Copy'})}
            </Button>
          </InputRightElement>
        </InputGroup>
      </Center>
      <Center>
        <InputGroup size="md" w={{base: '96%', sm: '400px'}}>
          <Input
            pr="4.5rem"
            value={copycopyright1.value}
            onFocus={e => e.target.select()}
            onChange={() => {}}
            border="0px"
          />
          <InputRightElement width="4.5rem">
            <Button h="1.75rem" size="sm" onClick={copycopyright1.onCopy}>
              {copycopyright1.hasCopied
                ? convertLang({ja: 'OK', en: 'Copied'})
                : convertLang({ja: 'コピー', en: 'Copy'})}
            </Button>
          </InputRightElement>
        </InputGroup>
      </Center>
    </Box>
  );
};
