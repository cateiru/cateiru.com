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
import {parseShotDate, sliceStr} from '../../utils/parse';
import {PublicProduct} from '../../utils/types';
import useLanguage from '../useLanguage';

export const ProductCard: React.FC<{prod: PublicProduct}> = ({prod}) => {
  const {colorMode} = useColorMode();
  const {convertLang, lang} = useLanguage();

  return (
    <NextLink passHref href={`/product/${prod.id}`}>
      <Box
        w="100%"
        minH="350px"
        boxShadow={
          colorMode === 'light'
            ? '0px 1px 26px -3px #a0acc0'
            : '0px 1px 26px -3px #000'
        }
        borderRadius="56px"
        p="1rem"
        position="relative"
        _hover={{
          bgColor: colorMode === 'light' ? 'gray.100' : 'gray.700',
        }}
        transition=".2s cubic-bezier(0.45, 0, 0.55, 1)"
      >
        {prod.thumbnail ? (
          <Image
            src={prod.thumbnail}
            w="100%"
            alt="thumbnail"
            borderRadius="40"
          />
        ) : (
          <Center
            w="100%"
            h="160px"
            bgColor={colorMode === 'dark' ? 'gray.700' : 'gray.100'}
            borderRadius="40"
          >
            <Text fontSize="1.2rem" color="gray.400">
              {convertLang({ja: '画像なし', en: 'No Image'})}
            </Text>
          </Center>
        )}
        <Box mt="1rem">
          <Center>
            <Text fontWeight="bold" fontSize="1.2rem">
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

        <Box mb="1rem" mt=".5rem" overflow="auto" minH="100px" maxH="150px">
          <Text
            as="pre"
            fontFamily="'Noto Sans JP', sans-serif"
            fontSize="1rem"
            maxW="300px"
            minH="120px"
            maxH="200px"
          >
            {convertLang({
              ja: sliceStr(prod.detail_ja),
              en: sliceStr(prod.detail),
            })}
          </Text>
        </Box>
      </Box>
    </NextLink>
  );
};
