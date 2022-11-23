import {Center, Box, Heading, Skeleton} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {Back} from '../Back';
import useLanguage from '../useLanguage';
import {ContactForm, ContactFormProps} from './ContactForm';

const Contact = () => {
  const router = useRouter();
  const {convertLang} = useLanguage();
  const [data, setData] = React.useState<ContactFormProps | null>(null);

  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;
    const d: ContactFormProps = {};

    if (typeof query['subject'] === 'string') {
      d.subject = query['subject'];
    }
    if (typeof query['url'] === 'string') {
      d.url = query['url'];
    }
    if (typeof query['name'] === 'string') {
      d.name = query['name'];
    }
    if (typeof query['mail'] === 'string') {
      d.mail = query['mail'];
    }
    if (typeof query['category'] === 'string') {
      d.category = query['category'];
    }
    if (typeof query['custom_title'] === 'string') {
      d.custom_title = query['custom_title'];
    }
    setData(d);
  }, [router.isReady, router.query]);

  return (
    <Center minHeight="100vh">
      <Box width={{base: '95%', md: '500px'}} mb="4rem" mt="3rem">
        <Back />
        <Heading mb="2rem" textAlign="center" mt="1.5rem">
          {convertLang({ja: 'お問い合わせ', en: 'Contact Us'})}
        </Heading>
        {data ? (
          <ContactForm {...data} />
        ) : (
          <Skeleton>
            <Box width="100%" h="700px" />
          </Skeleton>
        )}
      </Box>
    </Center>
  );
};

export default Contact;
