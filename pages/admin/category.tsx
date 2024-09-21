import { CategoryEdit } from "../../components/Admin/CategoryEdit";
import { Head } from "../../components/Common/Head";
import { useRequire } from "../../components/Require/useRequire";

const AdminBio = () => {
  const { show } = useRequire(true, "/");

  return (
    <>
      <Head title={{ ja: "カテゴリ詳細", en: "Category Detail" }} />
      {show && <CategoryEdit />}
    </>
  );
};

export default AdminBio;
