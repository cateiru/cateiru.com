import {Box, Center, Heading, Text, useColorMode} from '@chakra-ui/react';
import {Color} from './Color';
import {Copyright} from './Copyright';
import {Icon} from './Icon';
import {Link} from './Link';

export const BrandResources = () => {
  const {colorMode} = useColorMode();

  return (
    <Box w="100%" h="100%">
      <Center
        h="400px"
        bgColor={colorMode === 'dark' ? 'gray.700' : 'gray.100'}
      >
        <Heading
          fontWeight="bold"
          fontSize="4rem"
          marginX={{base: '1rem', sm: '5rem'}}
        >
          Cateiru&apos;s&nbsp;
          <Text
            as="span"
            fontWeight="bold"
            fontSize="4rem"
            display="inline-block"
            background="linear-gradient(124deg, #2bc4cf, #572bcf, #cf2ba1)"
            backgroundClip="text"
          >
            Brand&nbsp;
          </Text>
          Resources
        </Heading>
      </Center>
      <Icon />
      <Color />
      <Copyright />
      <Link />
    </Box>
  );
};
