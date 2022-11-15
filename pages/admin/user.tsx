import Head from 'next/head';
import {UserEdit} from '../../components/Admin/UserEdit';
import {useRequire} from '../../components/Require/useRequire';

const AdminUser = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head>
        <title>User Edit</title>
      </Head>
      {show && <UserEdit />}
    </>
  );
};

export default AdminUser;
