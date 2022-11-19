import Head from 'next/head';
import {CategoryEdit} from '../../components/Admin/CategoryEdit';
import {useRequire} from '../../components/Require/useRequire';

const AdminBio = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head>
        <title>Categories</title>
      </Head>
      {show && <CategoryEdit />}
    </>
  );
};

export default AdminBio;
