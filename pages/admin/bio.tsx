import Head from 'next/head';
import {BioList} from '../../components/Admin/BioEdit';
import {useRequire} from '../../components/Require/useRequire';

const AdminBio = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head>
        <title>Bios</title>
      </Head>
      {show && <BioList />}
    </>
  );
};

export default AdminBio;
