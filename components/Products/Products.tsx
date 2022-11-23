import {Center, Box, Heading, SimpleGrid} from '@chakra-ui/react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import {Public} from '../../utils/types';
import useLanguage from '../useLanguage';
import {ProductCard} from './ProductCard';

const Products: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
  data: Public;
}> = ({next, r, data}) => {
  const {convertLang} = useLanguage();

  return (
    <Center minHeight="100vh" ref={r}>
      <Box width={{base: '90%', md: '80%', lg: '1000px'}}>
        <Heading textAlign="center" mb="2rem">
          {convertLang({en: 'Products', ja: '制作物'})}
        </Heading>
        <SimpleGrid columns={{base: 1, sm: 2, lg: 3}} spacing={10} mt="1rem">
          {data.products.map(v => {
            return <ProductCard key={v.id} prod={v} />;
          })}
        </SimpleGrid>
        <Center mt="1.5rem" cursor="pointer" onClick={next}>
          <IoCaretDown size="25px" />
        </Center>
      </Box>
    </Center>
  );
};

export default Products;
