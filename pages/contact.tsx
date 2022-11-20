import {Head} from '../components/Common/Head';
import Contact from '../components/Contact/Contact';

const ContactPage = () => {
  return (
    <>
      <Head title={{ja: 'お問い合わせ', en: 'Contact Us'}} />
      <Contact />
    </>
  );
};

export default ContactPage;
