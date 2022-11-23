import {Box, Center, Avatar} from '@chakra-ui/react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import {Public} from '../../utils/types';
import ProfileText from './ProfileText';

const Profile: React.FC<{next: () => void; data: Public}> = ({next, data}) => {
  return (
    <Center height="100vh">
      <Box>
        <Center mb="1rem">
          <Avatar src={data.avatar_url} size="lg" />
        </Center>
        <ProfileText data={data} />
        <Center mt="1rem" cursor="pointer" onClick={next}>
          <IoCaretDown size="25px" />
        </Center>
      </Box>
    </Center>
  );
};

export default Profile;
