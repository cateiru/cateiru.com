import { BioList } from "../../components/Admin/BioEdit";
import { Head } from "../../components/Common/Head";
import { useRequire } from "../../components/Require/useRequire";

const AdminBio = () => {
  const { show } = useRequire(true, "/");

  return (
    <>
      <Head title={{ ja: "略歴詳細", en: "Bio Detail" }} />
      {show && <BioList />}
    </>
  );
};

export default AdminBio;
