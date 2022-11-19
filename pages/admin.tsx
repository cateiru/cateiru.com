import Head from 'next/head';
import {AdminTop} from '../components/Admin/Top';
import {useRequire} from '../components/Require/useRequire';

const AdminPage = () => {
  const {show} = useRequire(true, '/sso');

  return (
    <>
      <Head>
        <title>Admin</title>
      </Head>
      {show && <AdminTop />}
    </>
  );
};

export default AdminPage;
