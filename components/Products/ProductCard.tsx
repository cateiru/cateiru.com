import {
  Box,
  Center,
  Divider,
  Image,
  Text,
  useColorMode,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {parseShotDate} from '../../utils/parse';
import {PublicProduct} from '../../utils/types';
import useLanguage from '../useLanguage';

export const ProductCard: React.FC<{prod: PublicProduct}> = ({prod}) => {
  const {colorMode} = useColorMode();
  const {convertLang, lang} = useLanguage();

  return (
    <NextLink passHref href={`/product/${prod.id}`}>
      <Box
        w="100%"
        minH="100px"
        boxShadow={
          colorMode === 'light'
            ? '0px 1px 26px -3px #a0acc0'
            : '0px 1px 26px -3px #000'
        }
        borderRadius="35px"
        p="1rem"
        position="relative"
        _hover={{
          bgColor: colorMode === 'light' ? 'brand.200' : 'brand.600',
        }}
        transition=".2s cubic-bezier(0.45, 0, 0.55, 1)"
      >
        <Center display={{base: 'none', md: 'flex'}}>
          {prod.thumbnail ? (
            <Image
              src={prod.thumbnail}
              w="224px"
              h="126px"
              alt="thumbnail"
              borderRadius="20"
            />
          ) : (
            <Center
              w="224px"
              h="126px"
              bgColor={colorMode === 'dark' ? 'gray.700' : 'gray.100'}
              borderRadius="20"
            >
              <Text fontSize="1.2rem" color="gray.400">
                {convertLang({ja: '画像なし', en: 'No Image'})}
              </Text>
            </Center>
          )}
        </Center>
        <Box mt={{base: '0', md: '1rem'}}>
          <Center>
            <Text
              fontWeight="bold"
              fontSize="1.2rem"
              textAlign="center"
              whiteSpace="nowrap"
              textOverflow="ellipsis"
              overflow="hidden"
              background="linear-gradient(128deg, #E23D3D 0%, #EC44BD 100%)"
              backgroundClip="text"
            >
              {convertLang({ja: prod.name_ja, en: prod.name})}
            </Text>
          </Center>
          <Divider my=".5rem" />
          <Text
            textAlign="center"
            fontSize=".8rem"
            color="gray.500"
            fontWeight="bold"
          >
            {parseShotDate(prod.dev_time, lang)}
          </Text>
        </Box>
      </Box>
    </NextLink>
  );
};
