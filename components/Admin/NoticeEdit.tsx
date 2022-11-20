import {
  Box,
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  useToast,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {useForm} from 'react-hook-form';
import {IoArrowBack} from 'react-icons/io5';
import useSWR from 'swr';
import {api} from '../../utils/api';
import {fetcher, SWRError} from '../../utils/swr';
import {Notice} from '../../utils/types';
import useLanguage from '../useLanguage';
interface Form {
  discord_webhook?: string;
  slack_webhook?: string;
  mail?: string;
}

export const NoticeEdit = () => {
  const {convertLang} = useLanguage();
  const toast = useToast();
  const {
    handleSubmit,
    register,
    setValue,
    formState: {errors, isSubmitting},
  } = useForm<Form>();

  const {data} = useSWR<Notice, SWRError>('/user/notice', fetcher);

  React.useEffect(() => {
    if (data) {
      if (data.discord_webhook) {
        setValue('discord_webhook', data.discord_webhook);
      }
      if (data.slack_webhook) {
        setValue('slack_webhook', data.slack_webhook);
      }
      if (data.mail) {
        setValue('mail', data.mail);
      }
    }
  }, [data]);

  const onSubmit = async (d: Form) => {
    if (!data) {
      return () => {};
    }

    const form = new FormData();
    let changed = false;
    if (d.discord_webhook && data?.discord_webhook !== d.discord_webhook) {
      form.append('discord_webhook', d.discord_webhook);
      changed = true;
    }
    if (d.slack_webhook && data?.slack_webhook !== d.slack_webhook) {
      form.append('slack_webhook', d.slack_webhook);
      changed = true;
    }
    if (d.mail && data?.mail !== d.mail) {
      form.append('mail', d.mail);
      changed = true;
    }

    if (!changed) {
      return () => {};
    }

    const res = await fetch(api('/user/notice'), {
      method: 'PUT',
      credentials: 'include',
      mode: 'cors',
      body: form,
    });

    if (res.ok) {
      toast({
        status: 'success',
        title: convertLang({ja: '更新しました', en: 'Success updated'}),
      });
    } else {
      toast({
        status: 'error',
        title: (await res.json()).message,
      });
    }

    return () => {};
  };

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ja: '通知先編集', en: 'Notice destination Edit'})}
      </Heading>
      <Box
        mx={{base: '.5rem', sm: '1.5rem', md: '0'}}
        display={{base: 'block', md: 'flex'}}
        alignItems="center"
        flexDirection="column"
      >
        <Box width={{base: 'auto', md: '500px'}}>
          <NextLink passHref href="/admin">
            <Button my="1rem" variant="ghost" leftIcon={<IoArrowBack />} as="a">
              {convertLang({ja: '戻る', en: 'Back'})}
            </Button>
          </NextLink>

          <form onSubmit={handleSubmit(onSubmit)}>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.discord_webhook)}>
              <FormLabel htmlFor="discord_webhook">
                {convertLang({ja: 'Discoed WebHook', en: 'Discoed WebHook'})}
              </FormLabel>
              <Input
                id="discord_webhook"
                placeholder="https://"
                {...register('discord_webhook', {
                  pattern: {
                    value: /https?:\/\/([\da-z.-]+).([a-z.]{2,6})[/\w .-]*\/?/,
                    message: convertLang({
                      ja: '正しいURLを入力してください',
                      en: 'Please enter the correct URL',
                    }),
                  },
                })}
              />
              <FormErrorMessage>
                {errors.discord_webhook && errors.discord_webhook.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.slack_webhook)}>
              <FormLabel htmlFor="slack_webhook">
                {convertLang({ja: 'Slack WebHook', en: 'Slack WebHook'})}
              </FormLabel>
              <Input
                id="slack_webhook"
                placeholder="https://"
                {...register('slack_webhook', {
                  pattern: {
                    value: /https?:\/\/([\da-z.-]+).([a-z.]{2,6})[/\w .-]*\/?/,
                    message: convertLang({
                      ja: '正しいURLを入力してください',
                      en: 'Please enter the correct URL',
                    }),
                  },
                })}
              />
              <FormErrorMessage>
                {errors.slack_webhook && errors.slack_webhook.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.mail)}>
              <FormLabel htmlFor="mail">
                {convertLang({ja: 'メールアドレス', en: 'Email Address'})}
              </FormLabel>
              <Input
                id="mail"
                {...register('mail', {
                  pattern: {
                    value:
                      /^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$/,
                    message: convertLang({
                      ja: '正しいメールアドレスを入力してください',
                      en: 'Please enter the correct Email address',
                    }),
                  },
                })}
              />
              <FormErrorMessage>
                {errors.mail && errors.mail.message}
              </FormErrorMessage>
            </FormControl>

            <Button
              mt={4}
              w={{base: '100%', md: 'auto'}}
              colorScheme="blue"
              isLoading={isSubmitting}
              type="submit"
            >
              {convertLang({ja: '更新', en: 'Submit'})}
            </Button>
          </form>
        </Box>
      </Box>
    </Box>
  );
};
