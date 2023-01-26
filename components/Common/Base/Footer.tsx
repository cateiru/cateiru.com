import {Center, Divider, Box, Text, Link} from '@chakra-ui/react';
import NextLink from 'next/link';

const Footer = () => {
  return (
    <>
      <Center>
        <Box width="95%" margin="1rem 0 1rem 0">
          <Divider />
        </Box>
      </Center>
      <Text textAlign="center" mb="1.5rem">
        &copy; {new Date().getFullYear()}{' '}
        <NextLink href="/">
          <Text as="span" _hover={{borderBottom: '1px'}}>
            cateiru
          </Text>
        </NextLink>{' '}
        -{' '}
        <Link href="https://github.com/cateiru/cateiru.com" isExternal>
          GitHub
        </Link>
      </Text>
    </>
  );
};

export default Footer;
