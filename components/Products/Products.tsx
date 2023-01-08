import {SimpleGrid} from '@chakra-ui/react';
import React from 'react';
import {Public} from '../../utils/types';
import {Section} from '../Common/Section';
import {ProductCard} from './ProductCard';

const Products: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
  data: Public;
}> = ({next, r, data}) => {
  return (
    <Section next={next} r={r} heading={{en: 'Products', ja: '制作物'}}>
      <SimpleGrid columns={{base: 1, sm: 2, lg: 3}} spacing={10} mt="1rem">
        {data.products.map(v => {
          return <ProductCard key={v.id} prod={v} />;
        })}
      </SimpleGrid>
    </Section>
  );
};

export default Products;
