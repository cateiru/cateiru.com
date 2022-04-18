import {Center, Divider, Box, Text} from '@chakra-ui/react';

const Footer = () => {
  return (
    <>
      <Center>
        <Box width="95%" margin="1rem 0 1rem 0">
          <Divider />
        </Box>
      </Center>
      <Text textAlign="center" mb="1.5rem">
        &copy; {new Date().getFullYear()} cateiru
      </Text>
    </>
  );
};

export default Footer;
