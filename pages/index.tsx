import type {NextPage} from 'next';
import Head from 'next/head';
import Index from '../components/Index';

const Home: NextPage = () => {
  return (
    <>
      <Head>
        <title>Cateiru</title>
      </Head>
      <Index />
    </>
  );
};

export default Home;
