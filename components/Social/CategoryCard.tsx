import {
  Box,
  Center,
  Divider,
  Flex,
  Image,
  Link,
  Text,
  useColorMode,
} from "@chakra-ui/react";
import { TbPhoto } from "react-icons/tb";
import type { PublicLink } from "../../utils/types";
import useLanguage from "../useLanguage";

export const CategoryCard: React.FC<{ links: PublicLink }> = ({ links }) => {
  const { colorMode } = useColorMode();
  const { convertLang } = useLanguage();

  return (
    <Box
      w="100%"
      minH="350px"
      boxShadow={
        colorMode === "light"
          ? "0px 1px 26px -3px #a0acc0"
          : "0px 1px 26px -3px #000"
      }
      borderRadius="20"
      p="1rem"
      position="relative"
    >
      <Flex justifyContent="center" mt=".5rem" alignItems="center">
        <Text fontSize="1.5rem">{links.emoji}</Text>
        <Text
          fontWeight="bold"
          fontSize="1.2rem"
          ml=".5rem"
          background="linear-gradient(128deg, #E23D3D 0%, #EC44BD 100%)"
          backgroundClip="text"
        >
          {convertLang({ ja: links.category_name_ja, en: links.category_name })}
        </Text>
      </Flex>
      <Divider my="1rem" />
      <Box w="100%" mb="1rem">
        {links.links.map((v, i) => {
          return (
            <Link href={v.site_url} isExternal key={i} width="500px">
              <Box
                mt="1rem"
                py=".5rem"
                cursor="pointer"
                fontSize={{ base: "1rem", sm: "1rem" }}
                bgColor={colorMode === "dark" ? "brand.600" : "brand.200"}
                borderRadius="25"
                pl="1rem"
                fontWeight={{ base: "bold", sm: "medium" }}
                _hover={{
                  bgColor: colorMode === "dark" ? "brand.500" : "brand.300",
                }}
                transition=".2s cubic-bezier(0.45, 0, 0.55, 1)"
              >
                <Flex alignItems="center">
                  {v.favicon_url ? (
                    <Box mr=".5rem">
                      <Image
                        src={v.favicon_url}
                        width="20px"
                        alt="favicon"
                        loading="lazy"
                      />
                    </Box>
                  ) : (
                    <Center mr=".5rem">
                      <TbPhoto size="20px" />
                    </Center>
                  )}
                  <Text maxW="calc(100% - 4rem)">
                    {convertLang({ ja: v.name_ja, en: v.name })}
                  </Text>
                </Flex>
              </Box>
            </Link>
          );
        })}
      </Box>
    </Box>
  );
};
