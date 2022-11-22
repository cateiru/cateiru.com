import {
  Box,
  Center,
  Heading,
  IconButton,
  Image,
  Link,
  Text,
  useColorMode,
} from '@chakra-ui/react';
import React from 'react';
import {TbBrandGithub, TbLink} from 'react-icons/tb';
import {parseShotDate} from '../../utils/parse';
import {PublicProduct} from '../../utils/types';
import {Back} from '../Back';
import Markdown from '../Markdown/Markdown';
import useLanguage from '../useLanguage';

export const ProductPage: React.FC<{
  product: PublicProduct;
}> = ({product}) => {
  const {convertLang, lang} = useLanguage();
  const {colorMode} = useColorMode();

  return (
    <Center minH="100vh">
      <Box w={{base: '95%', md: '500px'}} minH="700px">
        {product.thumbnail ? (
          <Image
            src={product.thumbnail}
            w="100%"
            alt="thumbnail"
            borderRadius="25"
            boxShadow={
              colorMode === 'light'
                ? '0px 1px 26px -3px #a0acc0'
                : '0px 1px 26px -3px #000'
            }
          />
        ) : (
          <Center
            w="100%"
            h="300px"
            bgColor={colorMode === 'dark' ? 'gray.700' : 'gray.100'}
            borderRadius="25"
            boxShadow={
              colorMode === 'light'
                ? '0px 1px 26px -3px #a0acc0'
                : '0px 1px 26px -3px #000'
            }
          >
            <Text fontSize="1.2rem" color="gray.400">
              {convertLang({ja: '画像なし', en: 'No Image'})}
            </Text>
          </Center>
        )}
        <Back />
        <Heading textAlign="center" mt="1rem">
          {convertLang({ja: product.name_ja, en: product.name})}
        </Heading>
        <Text
          textAlign="center"
          fontSize=".8rem"
          color="gray.500"
          fontWeight="bold"
        >
          {parseShotDate(product.dev_time, lang)}
        </Text>
        <Box mb="1rem" mt=".5rem" overflow="auto" mx={{base: '.5rem', md: '0'}}>
          {/* <Text
            as="pre"
            fontFamily="'Noto Sans JP', sans-serif"
            fontSize="1rem"
            maxW="500px"
          > */}
          <Markdown>
            {convertLang({ja: product.detail_ja, en: product.detail})}
          </Markdown>
          {/* </Text> */}
        </Box>
        <Center mt="1rem">
          <IconButton
            aria-label="github"
            icon={<TbLink size="25" />}
            as={Link}
            href={product.site_url}
            size="lg"
            isExternal
            borderRadius="25"
          />
          {product.github_url && (
            <IconButton
              aria-label="github"
              icon={<TbBrandGithub size="25" />}
              as={Link}
              href={product.github_url}
              size="lg"
              isExternal
              borderRadius="25"
              ml="1rem"
            />
          )}
        </Center>
      </Box>
    </Center>
  );
};
