import {ContactDefaultEdit} from '../../components/Admin/ContactDefaultEdit';
import {Head} from '../../components/Common/Head';
import {useRequire} from '../../components/Require/useRequire';

const AdminBio = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head title={{ja: 'お問い合わせカスタム', en: 'Custom Contact'}} />
      {show && <ContactDefaultEdit />}
    </>
  );
};

export default AdminBio;
