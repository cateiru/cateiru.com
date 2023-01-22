import {
  Center,
  Box,
  Heading,
  Skeleton,
  Text,
  useColorMode,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {Back} from '../Back';
import useLanguage from '../useLanguage';
import {ContactForm, ContactFormProps} from './ContactForm';
import {PreviewUserData} from './PreviewUserData';

const Contact = () => {
  const router = useRouter();
  const {colorMode} = useColorMode();
  const {convertLang} = useLanguage();
  const [data, setData] = React.useState<ContactFormProps | null>(null);
  const [description, setDescription] = React.useState<string>();

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
    if (typeof query['description'] === 'string') {
      setDescription(query['description']);
    }
    setData(d);
  }, [router.isReady, router.query]);

  return (
    <Center minHeight="100vh">
      <Box width={{base: '95%', md: '500px'}} mb="4rem" mt="3rem">
        <Back href="/" />
        <Heading mb="2rem" textAlign="center" mt="1.5rem">
          {convertLang({ja: 'お問い合わせ', en: 'Contact Us'})}
        </Heading>
        {description && (
          <Text
            as="pre"
            w="100%"
            minH="100px"
            py="1rem"
            px=".5rem"
            bgColor={colorMode === 'dark' ? 'gray.700' : 'gray.100'}
            borderRadius="5px"
            borderLeftWidth="4px"
            borderLeftColor={colorMode === 'dark' ? 'gray.500' : 'gray.400'}
            mb="1.5rem"
            whiteSpace="pre-wrap"
            fontSize="1rem"
            fontFamily="'Kosugi Maru', sans-serif"
          >
            {decodeURIComponent(description)}
          </Text>
        )}
        <PreviewUserData />
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
