import {NoticeEdit} from '../../components/Admin/NoticeEdit';
import {Head} from '../../components/Common/Head';
import {useRequire} from '../../components/Require/useRequire';

const AdminLocation = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head title={{ja: '通知詳細', en: 'Notice Detail'}} />
      {show && <NoticeEdit />}
    </>
  );
};

export default AdminLocation;
