import {Flex, Box} from '@chakra-ui/react';
import React from 'react';
import Footer from './Footer';
import Header from './Header';

const Base: React.FC = props => {
  return (
    <Flex flexDirection="column" minHeight="100vh">
      <Box>
        <Header />
        {props.children}
      </Box>
      <Box marginTop="auto" as="footer">
        <Footer />
      </Box>
    </Flex>
  );
};

export default Base;
