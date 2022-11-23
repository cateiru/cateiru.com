import {ProductEdit} from '../../components/Admin/ProductEdit';
import {Head} from '../../components/Common/Head';
import {useRequire} from '../../components/Require/useRequire';

const AdminProducts = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head title={{ja: '制作物詳細', en: 'Products Detail'}} />
      {show && <ProductEdit />}
    </>
  );
};

export default AdminProducts;
