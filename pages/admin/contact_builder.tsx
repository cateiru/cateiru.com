import {ContactBuilder} from '../../components/Admin/ContactBuilder';
import {Head} from '../../components/Common/Head';
import {useRequire} from '../../components/Require/useRequire';

const ContactBuilderPage = () => {
  const {show} = useRequire(true, '/');

  return (
    <>
      <Head title={{ja: 'お問い合わせメーカー', en: 'Contact Maker'}} />
      {show && <ContactBuilder />}
    </>
  );
};

export default ContactBuilderPage;
