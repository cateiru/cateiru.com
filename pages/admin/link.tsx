import Head from 'next/head';
import {LinkEdit} from '../../components/Admin/LinkEdit';
import {useRequire} from '../../components/Require/useRequire';

const AdminBio = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head>
        <title>Links</title>
      </Head>
      {show && <LinkEdit />}
    </>
  );
};

export default AdminBio;
