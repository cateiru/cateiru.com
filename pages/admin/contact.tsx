import {ContactEdit} from '../../components/Admin/ContactEdit';
import {Head} from '../../components/Common/Head';
import {useRequire} from '../../components/Require/useRequire';

const AdminBio = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head title={{ja: 'お問い合わせ詳細', en: 'Contact Detail'}} />
      {show && <ContactEdit />}
    </>
  );
};

export default AdminBio;
