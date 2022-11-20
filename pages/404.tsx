import {Head} from '../components/Common/Head';
import {NotFound} from '../components/NotFound';

const Page = () => {
  return (
    <>
      <Head title={{ja: 'お探しのページは見つかりません', en: 'Not Found'}} />
      <NotFound />
    </>
  );
};

export default Page;
