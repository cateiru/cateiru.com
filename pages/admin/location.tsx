import Head from 'next/head';
import {LocationEdit} from '../../components/Admin/LocationEdit';
import {useRequire} from '../../components/Require/useRequire';

const AdminLocation = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head>
        <title>Location Edit</title>
      </Head>
      {show && <LocationEdit />}
    </>
  );
};

export default AdminLocation;
