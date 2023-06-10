import {
  FormControl,
  FormLabel,
  Input,
  Button,
  FormErrorMessage,
  Textarea,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {api} from '../../utils/api';
import useLanguage from '../useLanguage';

export interface ContactFormProps {
  name?: string;
  subject?: string;
  url?: string;
  mail?: string;
  category?: string;
  custom_title?: string;
}

interface FormValue {
  name: string;
  subject: string;
  mail: string;
  details: string;

  url?: string;
  category?: string;
  custom_title?: string;
  custom_value?: string;
}

export const ContactForm: React.FC<ContactFormProps> = props => {
  const {convertLang, lang} = useLanguage();
  const toast = useToast();
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    reset,
  } = useForm<FormValue>({defaultValues: props});

  const onSubmit = async (v: FormValue) => {
    const form = new FormData();
    form.append('name', v.name);
    form.append('subject', v.subject);
    form.append('mail', v.mail);
    form.append('detail', v.details);

    form.append('lang', lang);

    if (v.url) {
      form.append('url', v.url);
    }
    if (v.category) {
      form.append('category', v.category);
    }
    if (v.custom_title && v.custom_value) {
      form.append('custom_title', v.custom_title);
      form.append('custom_value', v.custom_value);
    }

    try {
      const r = await fetch(api('/contact'), {
        method: 'POST',
        credentials: 'include',
        mode: 'cors',
        body: form,
      });

      if (r.status !== 200) {
        throw new Error((await r.json()).message);
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

    return () => {};
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={Boolean(errors.name)}>
        <FormLabel htmlFor="name">
          {convertLang({ja: 'お名前', en: 'Your Name'})}
        </FormLabel>
        <Input
          id="name"
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
      <FormControl isInvalid={Boolean(errors.mail)} mt="1rem">
        <FormLabel htmlFor="mail">
          {convertLang({ja: 'メールアドレス', en: 'mail Address'})}
        </FormLabel>
        <Input
          id="mail"
          type="mail"
          {...register('mail', {
            required: convertLang({
              ja: 'この項目は必須です',
              en: 'This is required',
            }),
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
      <FormControl isInvalid={Boolean(errors.url)} mt="1rem">
        <FormLabel htmlFor="url">
          {convertLang({ja: 'URL（オプション）', en: 'URL (Optional)'})}
        </FormLabel>
        <Input
          id="url"
          placeholder="https://"
          type="url"
          {...register('url')}
        />
        <FormErrorMessage>{errors.url && errors.url.message}</FormErrorMessage>
      </FormControl>
      {props.category && (
        <FormControl isInvalid={Boolean(errors.category)} mt="1rem">
          <FormLabel htmlFor="category">
            {convertLang({
              ja: 'カテゴリ',
              en: 'category',
            })}
          </FormLabel>
          <Input
            id="category"
            type="category"
            isDisabled
            {...register('category')}
          />
          <FormErrorMessage>
            {errors.category && errors.category.message}
          </FormErrorMessage>
        </FormControl>
      )}
      {props.custom_title && (
        <FormControl isInvalid={Boolean(errors.custom_value)} mt="1rem">
          <FormLabel htmlFor="custom_value">{props.custom_title}</FormLabel>
          <Input
            id="custom_value"
            type="custom_value"
            {...register('custom_value')}
          />
          <FormErrorMessage>
            {errors.custom_value && errors.custom_value.message}
          </FormErrorMessage>
        </FormControl>
      )}
      <FormControl isInvalid={Boolean(errors.subject)} mt="1rem">
        <FormLabel htmlFor="subject">
          {convertLang({ja: '件名', en: 'Subject'})}
        </FormLabel>
        <Input
          id="subject"
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
      <FormControl isInvalid={Boolean(errors.details)} mt="1rem">
        <FormLabel htmlFor="details">
          {convertLang({ja: 'お問い合わせ内容', en: 'Inquiry Details'})}
        </FormLabel>
        <Textarea
          id="details"
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
        isLoading={isSubmitting}
        type="submit"
        colorScheme="cateiru"
        width="100%"
      >
        {convertLang({ja: '送信する', en: 'Submit'})}
      </Button>
    </form>
  );
};
