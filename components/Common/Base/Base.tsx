import { Flex, Box } from "@chakra-ui/react";
import type React from "react";
import Footer from "./Footer";
import Header from "./Header";

const Base: React.FC<{ children: React.ReactNode }> = (props) => {
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
