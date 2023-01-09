import {Box, Center, Heading, useColorMode} from '@chakra-ui/react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import {MultiLang} from '../../utils/config/lang';
import useLanguage from '../useLanguage';

export const Section: React.FC<{
  next: () => void;
  r?: React.MutableRefObject<HTMLDivElement>;
  heading?: MultiLang;
  children: React.ReactNode;
}> = props => {
  const {colorMode} = useColorMode();
  const {convertLang} = useLanguage();

  return (
    <Center minHeight="100vh" ref={props.r} mb={{base: '150px', md: '25px'}}>
      <Box width={{base: '90%', md: '600px', lg: '900px'}}>
        {props.heading && (
          <Heading textAlign="center" mb="2rem" id={props.heading.en}>
            {convertLang(props.heading)}
          </Heading>
        )}
        {props.children}
        <Center mt="1.5rem">
          <Center
            w="100px"
            h="2rem"
            borderRadius="25"
            _hover={{
              bgColor: props.heading
                ? colorMode === 'dark'
                  ? 'gray.600'
                  : 'gray.200'
                : 'rgba(255, 255, 255, 0.2)',
            }}
            transition=".2s cubic-bezier(0.45, 0, 0.55, 1)"
            cursor="pointer"
            onClick={props.next}
          >
            <IoCaretDown size="25px" />
          </Center>
        </Center>
      </Box>
    </Center>
  );
};
