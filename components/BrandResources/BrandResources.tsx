import {
  Box,
  Center,
  Divider,
  Heading,
  Text,
  useColorMode,
} from "@chakra-ui/react";
import { Color } from "./Color";
import { Copyright } from "./Copyright";
import { Icon } from "./Icon";
import { IconGenerator } from "./IconGenerator";
import { Link } from "./Link";

export const BrandResources = () => {
  const { colorMode } = useColorMode();

  return (
    <Box w="100%" h="100%">
      <Center
        h={{ base: "100vh", md: "400px" }}
        bgColor={colorMode === "dark" ? "gray.700" : "gray.100"}
        mb="10rem"
      >
        <Heading
          fontWeight="bold"
          fontSize={{ base: "3rem", md: "4rem" }}
          marginX={{ base: "3rem", sm: "5rem" }}
        >
          Cateiru&apos;s&nbsp;
          <Text
            as="span"
            fontWeight="bold"
            fontSize={{ base: "3rem", md: "4rem" }}
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
      <D />
      <Color />
      <D />
      <Copyright />
      <D />
      <Link />
    </Box>
  );
};

const D = () => {
  return (
    <Center>
      <Divider my="7rem" w={{ base: "80%", md: "700px" }} />
    </Center>
  );
};
