import {
  FormControl,
  FormLabel,
  Input,
  Button,
  FormErrorMessage,
  Center,
  Box,
  Textarea,
  Heading,
  useToast,
} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useForm, FieldValues} from 'react-hook-form';
import {IoArrowBack} from 'react-icons/io5';
import {Form} from '../../utils/form';
import useLanguage from '../useLanguage';

interface FormValue {
  name: string;
  subject: string;
  url?: string;
  email: string;
  details: string;
}

const Contact = () => {
  const [useURL, setUseURL] = React.useState(false);
  const [load, setLoad] = React.useState(false);

  const router = useRouter();
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    reset,
    setValue,
  } = useForm<FormValue>();
  const {lang, convertLang} = useLanguage();
  const toast = useToast();

  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['subject'] === 'string') {
      setValue('subject', query['subject']);
    }
    if (typeof query['url'] === 'string') {
      setUseURL(true);
      setValue('url', query['url']);
    }
    if (typeof query['name'] === 'string') {
      setValue('name', query['name']);
    }
    if (typeof query['mail'] === 'string') {
      setValue('email', query['mail']);
    }
  }, [router.isReady, router.query]);

  const onSubmit = (values: FieldValues) => {
    const f = async () => {
      const sendForm: Form = {
        name: values.name,
        subject: values.subject,
        mail: values.email,
        url: values.url,
        body: values.details,
        date: new Date().toString(),
        lang: lang,
      };

      try {
        setLoad(true);
        const r = await fetch('/api/form', {
          method: 'POST',
          body: JSON.stringify(sendForm),
          headers: {'Content-Type': 'application/json'},
        });

        if (r.status !== 200) {
          throw new Error(r.statusText);
        }

        toast({
          status: 'success',
          title: convertLang({ja: '送信しました', en: 'Submitted.'}),
        });

        reset();
      } catch (e) {
        if (e instanceof Error) {
          toast({
            status: 'error',
            title: e.message,
          });
        }
      }

      setLoad(false);
    };

    f();
    return () => {};
  };

  return (
    <Center minHeight="100vh">
      <Box width={{base: '95%', md: '500px'}} mb="4rem" mt="2rem">
        {!useURL && (
          <Button
            variant="ghost"
            leftIcon={<IoArrowBack size="20px" />}
            onClick={() => {
              router.push('/');
            }}
            zIndex="1001"
          >
            {convertLang({ja: '戻る', en: 'back'})}
          </Button>
        )}
        <Heading mb="1.2rem" textAlign="center" mt="1.5rem">
          {convertLang({ja: 'お問い合わせ', en: 'Contact Us'})}
        </Heading>
        <form onSubmit={handleSubmit(onSubmit)}>
          <FormControl isInvalid={Boolean(errors.name)}>
            <FormLabel htmlFor="name">
              {convertLang({ja: 'お名前', en: 'Your Name'})}
            </FormLabel>
            <Input
              id="name"
              placeholder="name"
              type="text"
              {...register('name', {
                required: convertLang({
                  ja: 'この項目は必須です',
                  en: 'This is required',
                }),
              })}
            />
            <FormErrorMessage>
              {errors.name && errors.name.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.email)} mt=".5rem">
            <FormLabel htmlFor="email">
              {convertLang({ja: 'メールアドレス', en: 'Email Address'})}
            </FormLabel>
            <Input
              id="email"
              placeholder="email"
              type="email"
              {...register('email', {
                required: convertLang({
                  ja: 'この項目は必須です',
                  en: 'This is required',
                }),
                pattern: {
                  value:
                    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                  message: convertLang({
                    ja: '正しいメールアドレスを入力してください。',
                    en: 'Please enter your correct email address.',
                  }),
                },
              })}
            />
            <FormErrorMessage>
              {errors.email && errors.email.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.url)} mt=".5rem">
            <FormLabel htmlFor="url">
              {convertLang({ja: 'URL（オプション）', en: 'URL (Optional)'})}
            </FormLabel>
            <Input
              id="url"
              placeholder="url"
              type="url"
              disabled={useURL}
              {...register('url')}
            />
            <FormErrorMessage>
              {errors.url && errors.url.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.subject)} mt=".5rem">
            <FormLabel htmlFor="subject">
              {convertLang({ja: '件名', en: 'Subject'})}
            </FormLabel>
            <Input
              id="subject"
              placeholder="subject"
              type="text"
              {...register('subject', {
                required: convertLang({
                  ja: 'この項目は必須です',
                  en: 'This is required',
                }),
              })}
            />
            <FormErrorMessage>
              {errors.subject && errors.subject.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.details)} mt=".5rem">
            <FormLabel htmlFor="details">
              {convertLang({ja: 'お問い合わせ内容', en: 'Inquiry Details'})}
            </FormLabel>
            <Textarea
              id="details"
              placeholder="details"
              resize="vertical"
              height="150px"
              {...register('details', {
                required: convertLang({
                  ja: 'この項目は必須です',
                  en: 'This is required',
                }),
              })}
            />
            <FormErrorMessage>
              {errors.details && errors.details.message}
            </FormErrorMessage>
          </FormControl>
          <Button
            mt={4}
            isLoading={isSubmitting || load}
            type="submit"
            colorScheme="green"
            width={{base: '100%', md: 'auto'}}
          >
            {convertLang({ja: '送信する', en: 'Submit'})}
          </Button>
        </form>
      </Box>
    </Center>
  );
};

export default Contact;
