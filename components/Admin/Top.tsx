import {SimpleGrid} from '@chakra-ui/react';
import {AllUsersCard} from './AllusersCard';
import {BioCard} from './BioCard';
import {ContactCard} from './ContactCard';
import {ContactDefaultCard} from './ContactDefaultCard';
import {ControllerCard} from './ControllerCard';
import {LinkCard} from './LinkCard';
import {NoticeCard} from './NoticeCard';
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
      <ContactCard />
      <BioCard />
      <ProductCard />
      <LinkCard />
      <NoticeCard />
      <ContactDefaultCard />
      <ControllerCard />
    </SimpleGrid>
  );
};
