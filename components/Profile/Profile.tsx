import {Center, Avatar} from '@chakra-ui/react';
import React from 'react';
import {Public} from '../../utils/types';
import {Section} from '../Common/Section';
import ProfileText from './ProfileText';

const Profile: React.FC<{next: () => void; data: Public}> = ({next, data}) => {
  return (
    <Section next={next}>
      <Center mb="1rem">
        <Avatar src={data.avatar_url} size="lg" />
      </Center>
      <ProfileText data={data} />
    </Section>
  );
};

export default Profile;
