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
import {useForm} from 'react-hook-form';
import {IoArrowBack} from 'react-icons/io5';
import {useRecoilState} from 'recoil';
import {api} from '../../utils/api';
import {UserState} from '../../utils/state/atoms';
import useLanguage from '../useLanguage';

interface Form {
  given_name: string;
  family_name: string;
  given_name_ja: string;
  family_name_ja: string;
  user_id: string;
  mail: string;
  birth_date: string;
  location: string;
  location_ja: string;
  sso_token: string;
  avatar_url?: string;
}

export const UserEdit = () => {
  const [lang, convertLang] = useLanguage();
  const [user, setUser] = useRecoilState(UserState);
  const toast = useToast();
  const {
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useForm<Form>({
    defaultValues: {
      given_name: user?.given_name,
      given_name_ja: user?.given_name_ja,
      family_name: user?.family_name,
      family_name_ja: user?.family_name_ja,
      birth_date: new Date(user?.birth_date ?? '')
        .toISOString()
        .substring(0, 10),
      location: user?.location,
      location_ja: user?.location_ja,
    },
  });

  const onSubmit = async (v: Form) => {
    if (!user) {
      return () => {};
    }

    const form = new FormData();
    let changed = false;
    const newUser = {...user};

    if (v.family_name !== user?.family_name) {
      form.append('family_name', v.family_name);
      newUser.family_name = v.family_name;
      changed = true;
    }
    if (v.family_name_ja !== user?.family_name_ja) {
      form.append('family_name_ja', v.family_name_ja);
      newUser.family_name_ja = v.family_name_ja;
      changed = true;
    }
    if (v.given_name !== user?.given_name) {
      form.append('given_name', v.given_name);
      newUser.given_name = v.given_name;
      changed = true;
    }
    if (v.given_name_ja !== user?.given_name_ja) {
      form.append('given_name_ja', v.given_name_ja);
      newUser.given_name_ja = v.given_name_ja;
      changed = true;
    }
    if (
      v.birth_date !==
      new Date(user?.birth_date ?? '').toISOString().substring(0, 10)
    ) {
      const d = new Date(v.birth_date);
      form.append('birth_date', d.toISOString());
      newUser.birth_date = d.toISOString();
      changed = true;
    }
    if (v.location !== user?.location) {
      form.append('location', v.location);
      newUser.location = v.location;
      changed = true;
    }
    if (v.location_ja !== user?.location_ja) {
      form.append('location_ja', v.location_ja);
      newUser.location_ja = v.location_ja;
      changed = true;
    }

    if (!changed) {
      return () => {};
    }

    const res = await fetch(api('/user'), {
      method: 'PUT',
      credentials: 'include',
      mode: 'cors',
      body: form,
    });

    if (res.ok) {
      setUser(newUser);
      toast({
        status: 'success',
        title: convertLang({ja: '更新しました', en: 'Success Updated'}),
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
        {convertLang({ja: 'ユーザ編集', en: 'User Edit'})}
      </Heading>
      <Box
        mx={{base: '.5rem', sm: '1.5rem', md: '0'}}
        display={{base: 'block', md: 'flex'}}
        alignItems="center"
        flexDirection="column"
      >
        <Box width={{base: 'auto', md: '500px'}}>
          <NextLink passHref href="/admin">
            <Button my="1rem" variant="ghost" leftIcon={<IoArrowBack />}>
              {convertLang({ja: '戻る', en: 'Back'})}
            </Button>
          </NextLink>

          <form onSubmit={handleSubmit(onSubmit)}>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.given_name)}>
              <FormLabel htmlFor="given_name">
                {convertLang({ja: '名前(英語)', en: 'Given Name'})}
              </FormLabel>
              <Input
                id="given_name"
                {...register('given_name', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.given_name && errors.given_name.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.given_name_ja)}>
              <FormLabel htmlFor="given_name_ja">
                {convertLang({ja: '名前', en: 'Given Name (japanese)'})}
              </FormLabel>
              <Input
                id="given_name_ja"
                {...register('given_name_ja', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.given_name_ja && errors.given_name_ja.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.family_name)}>
              <FormLabel htmlFor="family_name">
                {convertLang({ja: '名字(英語)', en: 'Family Name'})}
              </FormLabel>
              <Input
                id="family_name"
                {...register('family_name', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.family_name && errors.family_name.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.family_name_ja)}>
              <FormLabel htmlFor="family_name_ja">
                {convertLang({ja: '名字', en: 'Family Name (japanese)'})}
              </FormLabel>
              <Input
                id="family_name_ja"
                {...register('family_name_ja', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.family_name_ja && errors.family_name_ja.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.birth_date)}>
              <FormLabel htmlFor="birth_date">
                {convertLang({ja: '誕生日', en: 'Birth Date'})}
              </FormLabel>
              <Input
                id="birth_date"
                type="date"
                {...register('birth_date', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.birth_date && errors.birth_date.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.location)}>
              <FormLabel htmlFor="location">
                {convertLang({ja: '出身地(英語)', en: 'Location'})}
              </FormLabel>
              <Input
                id="location"
                {...register('location', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.location && errors.location.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.location_ja)}>
              <FormLabel htmlFor="location_ja">
                {convertLang({ja: '出身地', en: 'location_ja (japanese)'})}
              </FormLabel>
              <Input
                id="location_ja"
                {...register('location_ja', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.location_ja && errors.location_ja.message}
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
