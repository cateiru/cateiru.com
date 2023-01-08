import {SimpleGrid} from '@chakra-ui/react';
import React from 'react';
import {Public} from '../../utils/types';
import {Section} from '../Common/Section';
import {CategoryCard} from './CategoryCard';

const Social: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
  data: Public;
}> = ({next, r, data}) => {
  return (
    <Section next={next} r={r} heading={{en: 'Links', ja: 'リンク'}}>
      <SimpleGrid columns={{base: 1, sm: 2, lg: 3}} spacing={10} mt="1rem">
        {data.links.map(v => (
          <CategoryCard links={v} key={v.category_id} />
        ))}
      </SimpleGrid>
    </Section>
  );
};

export default Social;
