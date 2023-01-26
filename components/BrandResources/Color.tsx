import {
  Box,
  Button,
  Center,
  Heading,
  SimpleGrid,
  useClipboard,
} from '@chakra-ui/react';
import React from 'react';
import {TbCheck} from 'react-icons/tb';
import useLanguage from '../useLanguage';

export const Color = () => {
  const {convertLang} = useLanguage();

  return (
    <Box mt="15vh">
      <Heading textAlign="center" mb="1rem" id="colors">
        {convertLang({ja: '色', en: 'Colors'})}
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
