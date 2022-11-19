import Head from 'next/head';
import {NoticeEdit} from '../../components/Admin/NoticeEdit';
import {useRequire} from '../../components/Require/useRequire';

const AdminLocation = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head>
        <title>Notice Edit</title>
      </Head>
      {show && <NoticeEdit />}
    </>
  );
};

export default AdminLocation;
