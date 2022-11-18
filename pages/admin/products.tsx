import Head from 'next/head';
import {ProductEdit} from '../../components/Admin/ProductEdit';
import {useRequire} from '../../components/Require/useRequire';

const AdminProducts = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head>
        <title>Location Edit</title>
      </Head>
      {show && <ProductEdit />}
    </>
  );
};

export default AdminProducts;
