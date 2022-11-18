import {SimpleGrid} from '@chakra-ui/react';
import {AllUsersCard} from './AllusersCard';
import {BioCard} from './BioCard';
import {ControllerCard} from './ControllerCard';
import {ProductCard} from './ProductCard';
import {UserCard} from './UserCard';

export const AdminTop = () => {
  return (
    <SimpleGrid
      columns={{base: 1, md: 2, xl: 3}}
      spacing={10}
      mt="5rem"
      mx={{base: '.5rem', md: '1rem'}}
    >
      <UserCard />
      <AllUsersCard />
      <BioCard />
      <ProductCard />
      <ControllerCard />
    </SimpleGrid>
  );
};
