import { Head } from "../components/Common/Head";
import { SSOPage } from "../components/Require/Sso";

const SSO = () => {
  return (
    <>
      <Head title={{ ja: "管理者向けログイン", en: "Admin Login" }} />
      <SSOPage />
    </>
  );
};

export default SSO;
