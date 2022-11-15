import {Button, Center} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import {api} from '../../utils/api';

export const SSOPage = () => {
  const router = useRouter();

  const onClick = () => {
    router.push(api('/login/url'));
  };

  return (
    <Center h="100vh">
      <Button onClick={onClick}>Login with Cateiru SSO</Button>
    </Center>
  );
};
