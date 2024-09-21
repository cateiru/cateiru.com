import { UserEdit } from "../../components/Admin/UserEdit";
import { Head } from "../../components/Common/Head";
import { useRequire } from "../../components/Require/useRequire";

const AdminUser = () => {
  const { show } = useRequire(true, "/");

  return (
    <>
      <Head title={{ ja: "ユーザ詳細", en: "User Details" }} />
      {show && <UserEdit />}
    </>
  );
};

export default AdminUser;
