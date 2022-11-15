import {SimpleGrid} from '@chakra-ui/react';
import {AllUsersCard} from './AllusersCard';
import {UserCard} from './UserCard';

export const AdminTop = () => {
  return (
    <SimpleGrid
      columns={{base: 1, md: 2}}
      spacing={10}
      mt="5rem"
      mx={{base: '.5rem', md: '1rem'}}
    >
      <UserCard />
      <AllUsersCard />
    </SimpleGrid>
  );
};
