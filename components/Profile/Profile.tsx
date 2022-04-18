import {Box, Center, Avatar} from '@chakra-ui/react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import config from '../../utils/config/config';
import ProfileText from './ProfileText';

const Profile: React.FC<{next: () => void}> = ({next}) => {
  return (
    <Center height="100vh">
      <Box>
        <Center mb="1rem">
          <Avatar src={config.avatar} size="lg" />
        </Center>
        <ProfileText />
        <Center mt="1rem" cursor="pointer" onClick={next}>
          <IoCaretDown size="25px" />
        </Center>
      </Box>
    </Center>
  );
};

export default Profile;
