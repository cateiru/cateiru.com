import {Center, Box, Heading, SimpleGrid} from '@chakra-ui/react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import {Public} from '../../utils/types';
import useLanguage from '../useLanguage';
import {CategoryCard} from './CategoryCard';

const Social: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
  data: Public;
}> = ({next, r, data}) => {
  const {convertLang} = useLanguage();

  return (
    <Center minHeight="100vh" ref={r}>
      <Box width={{base: '90%', md: '500px', lg: '900px'}}>
        <Heading textAlign="center" mb="2rem">
          {convertLang({en: 'Links', ja: 'リンク'})}
        </Heading>
        <SimpleGrid columns={{base: 1, sm: 2, lg: 3}} spacing={10} mt="1rem">
          {data.links.map(v => (
            <CategoryCard links={v} key={v.category_id} />
          ))}
        </SimpleGrid>
        <Center mt="1.5rem" cursor="pointer" onClick={next}>
          <IoCaretDown size="25px" />
        </Center>
      </Box>
    </Center>
  );
};

export default Social;
