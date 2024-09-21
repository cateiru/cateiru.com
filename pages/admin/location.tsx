import { LocationEdit } from "../../components/Admin/LocationEdit";
import { Head } from "../../components/Common/Head";
import { useRequire } from "../../components/Require/useRequire";

const AdminLocation = () => {
  const { show } = useRequire(true, "/");

  return (
    <>
      <Head title={{ ja: "場所詳細", en: "Location Detail" }} />
      {show && <LocationEdit />}
    </>
  );
};

export default AdminLocation;
