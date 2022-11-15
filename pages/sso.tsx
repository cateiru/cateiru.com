import Head from 'next/head';
import {SSOPage} from '../components/Require/Sso';

const SSO = () => {
  return (
    <>
      <Head>
        <title>Login</title>
      </Head>
      <SSOPage />
    </>
  );
};

export default SSO;
