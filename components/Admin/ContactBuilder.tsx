import {
  Box,
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  InputGroup,
  InputRightElement,
  Textarea,
  useClipboard,
} from '@chakra-ui/react';
import {useForm} from 'react-hook-form';
import {Back} from '../Back';
import useLanguage from '../useLanguage';

interface ContactBuilderForm {
  description?: string;
  custom_title?: string;
  category?: string;

  name?: string;
  mail?: string;
  url?: string;
  subject?: string;
}

const CONTACT_PAGE_URL = 'https://cateiru.com/contact';

export const ContactBuilder = () => {
  const {convertLang} = useLanguage();
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<ContactBuilderForm>();
  const {onCopy, hasCopied, setValue, value} = useClipboard(CONTACT_PAGE_URL);

  const onSubmit = (data: ContactBuilderForm) => {
    const u = new URL(CONTACT_PAGE_URL);
    if (data.name) {
      u.searchParams.set('name', data.name);
    }
    if (data.mail) {
      u.searchParams.set('mail', data.mail);
    }
    if (data.url) {
      u.searchParams.set('url', data.url);
    }
    if (data.subject) {
      u.searchParams.set('subject', data.subject);
    }
    if (data.category) {
      u.searchParams.set('category', data.category);
    }
    if (data.custom_title) {
      u.searchParams.set('custom_title', data.custom_title);
    }
    if (data.description) {
      u.searchParams.set('description', data.description);
    }

    setValue(u.toString());
  };

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ja: 'お問い合わせビルダー', en: 'Contact Builder'})}
      </Heading>
      <Box
        mx={{base: '.5rem', sm: '1.5rem', md: '0'}}
        display={{base: 'block', md: 'flex'}}
        alignItems="center"
        flexDirection="column"
      >
        <Box width={{base: 'auto', md: '500px'}}>
          <Back href="/admin" />
          <form onSubmit={handleSubmit(onSubmit)}>
            <FormControl isInvalid={Boolean(errors.name)}>
              <FormLabel htmlFor="name">
                {convertLang({ja: 'お名前', en: 'Your Name'})}
              </FormLabel>
              <Input id="name" type="text" {...register('name')} />
              <FormErrorMessage>
                {errors.name && errors.name.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl isInvalid={Boolean(errors.mail)} mt=".5rem">
              <FormLabel htmlFor="mail">
                {convertLang({ja: 'メールアドレス', en: 'mail Address'})}
              </FormLabel>
              <Input
                id="mail"
                type="mail"
                {...register('mail', {
                  pattern: {
                    value:
                      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                    message: convertLang({
                      ja: '正しいメールアドレスを入力してください。',
                      en: 'Please enter your correct mail address.',
                    }),
                  },
                })}
              />
              <FormErrorMessage>
                {errors.mail && errors.mail.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl isInvalid={Boolean(errors.url)} mt=".5rem">
              <FormLabel htmlFor="url">
                {convertLang({ja: 'URL', en: 'URL'})}
              </FormLabel>
              <Input
                id="url"
                placeholder="https://"
                type="url"
                {...register('url')}
              />
              <FormErrorMessage>
                {errors.url && errors.url.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl isInvalid={Boolean(errors.category)} mt=".5rem">
              <FormLabel htmlFor="category">
                {convertLang({
                  ja: 'カテゴリ',
                  en: 'category',
                })}
              </FormLabel>
              <Input id="category" type="category" {...register('category')} />
              <FormErrorMessage>
                {errors.category && errors.category.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl isInvalid={Boolean(errors.custom_title)} mt=".5rem">
              <FormLabel htmlFor="custom_title">
                {convertLang({
                  ja: 'カスタムタイトル',
                  en: 'Custom Title',
                })}
              </FormLabel>
              <Input
                id="custom_title"
                type="custom_title"
                {...register('custom_title')}
              />
              <FormErrorMessage>
                {errors.custom_title && errors.custom_title.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl isInvalid={Boolean(errors.description)} mt=".5rem">
              <FormLabel htmlFor="description">
                {convertLang({
                  ja: 'カスタムタイトル',
                  en: 'Custom Title',
                })}
              </FormLabel>
              <Textarea
                id="description"
                {...register('description')}
                h="200px"
              />
              <FormErrorMessage>
                {errors.description && errors.description.message}
              </FormErrorMessage>
            </FormControl>
            <Button mt={4} isLoading={isSubmitting} type="submit" width="100%">
              {convertLang({ja: 'URLを作成する', en: 'Build URL'})}
            </Button>
          </form>
          <Box mt="1rem" mb="1rem">
            <InputGroup size="md">
              <Input
                pr="4.5rem"
                value={value}
                onFocus={e => e.target.select()}
                onChange={() => {}}
              />
              <InputRightElement width="4.5rem">
                <Button h="1.75rem" size="sm" onClick={onCopy}>
                  {hasCopied
                    ? convertLang({ja: 'OK', en: 'Copied'})
                    : convertLang({ja: 'コピー', en: 'Copy'})}
                </Button>
              </InputRightElement>
            </InputGroup>
          </Box>
        </Box>
      </Box>
    </Box>
  );
};
