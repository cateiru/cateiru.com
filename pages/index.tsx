import type {NextPage} from 'next';
import {Head} from '../components/Common/Head';
import Index from '../components/Index';

const Home: NextPage = () => {
  return (
    <>
      <Head title={{ja: 'Cateiruのページ', en: "Cateiru's Page"}} />
      <Index />
    </>
  );
};

export default Home;
