import {
  Center,
  Box,
  Heading,
  Text,
  UnorderedList,
  ListItem,
  Link,
} from '@chakra-ui/react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import config from '../../utils/config/config';
import {Product} from '../../utils/config/product';
import {Public} from '../../utils/types';
import useLanguage from '../useLanguage';

const Products: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
  data: Public;
}> = ({next, r, data}) => {
  const {lang, convertLang} = useLanguage();

  const element = (product: Product) => {
    const d = () => {
      if (lang === 'en') {
        return `${product.createDate.getFullYear()}/${product.createDate.getMonth()}`;
      }
      return `${product.createDate.getFullYear()}年${product.createDate.getMonth()}月`;
    };

    return (
      <ListItem my=".5rem" key={product.name.en}>
        <Text as="span">{d()}</Text>
        <Link fontWeight="bold" ml=".5rem" href={product.link} isExternal>
          {convertLang(product.name)}
        </Link>
        <Text as="span"> - {convertLang(product.description)}</Text>
      </ListItem>
    );
  };

  return (
    <Center height="100vh" ref={r}>
      <Box width={{base: '90%', md: '700px'}}>
        <Heading textAlign="center">
          {convertLang({en: 'Products', ja: '制作物'})}
        </Heading>
        <Box whiteSpace="nowrap" overflowX="auto">
          <UnorderedList mt="1.5rem" paddingLeft=".5rem">
            {config.products.map(v => element(v))}
          </UnorderedList>
        </Box>
        <Center mt="1.5rem" cursor="pointer" onClick={next}>
          <IoCaretDown size="25px" />
        </Center>
      </Box>
    </Center>
  );
};

export default Products;
