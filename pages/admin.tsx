import { AdminTop } from "../components/Admin/Top";
import { Head } from "../components/Common/Head";
import { useRequire } from "../components/Require/useRequire";

const AdminPage = () => {
  const { show } = useRequire(true, "/sso");

  return (
    <>
      <Head title={{ ja: "管理者向けページ", en: "Admin Page" }} />
      {show && <AdminTop />}
    </>
  );
};

export default AdminPage;
