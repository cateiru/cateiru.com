import {
  Center,
  Box,
  Heading,
  Skeleton,
  Text,
  useColorMode,
  useToast,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {api} from '../../utils/api';
import {ContactDefaultSchema} from '../../utils/types';
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

  const toast = useToast();

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
    if (typeof query['id'] === 'string') {
      getCustom(query['id']);
      return;
    }
    setData(d);
  }, [router.isReady, router.query]);

  const getCustom = async (id: string) => {
    try {
      const r = await fetch(api(`/contact/default?id=${id}`));

      if (r.status !== 200) {
        throw new Error((await r.json()).message);
      }

      const resData = ContactDefaultSchema.parse(await r.json());

      const d: ContactFormProps = {
        url: resData.url,
        name: resData.name,
        mail: resData.email,
        category: resData.category,
        custom_title: resData.custom_title,
      };
      setData(d);

      resData.description && setDescription(resData.description);
    } catch (e) {
      if (e instanceof Error) {
        toast({
          status: 'error',
          title: e.message,
        });
      }
    }
  };

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
