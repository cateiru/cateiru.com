import {LinkEdit} from '../../components/Admin/LinkEdit';
import {Head} from '../../components/Common/Head';
import {useRequire} from '../../components/Require/useRequire';

const AdminBio = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head title={{ja: 'リンク詳細', en: 'Links Detail'}} />
      {show && <LinkEdit />}
    </>
  );
};

export default AdminBio;
