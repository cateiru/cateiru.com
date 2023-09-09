import {
  Box,
  Button,
  Center,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Spacer,
  Table,
  TableContainer,
  Tbody,
  Td,
  Textarea,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import {FormProvider} from 'react-hook-form';
import {MultiLang} from '../../utils/config/lang';
import {ContactDefault, ContactDefaultSchema} from '../../utils/types';
import {Back} from '../Back';
import useLanguage from '../useLanguage';
import {useList} from './useList';
import {useNew} from './useNew';
import {useUpdate} from './useUpdate';

interface ContactDefaultForm {
  name?: string;
  email?: string;
  url?: string;
  category?: string;
  custom_title?: string;
  description?: string;
}

export const ContactDefaultEdit = () => {
  const {convertLang} = useLanguage();
  const {
    data,
    error,
    update,
    updateValue,
    onUpdate,
    closeUpdateModal,
    createModal,
    updateModal,
  } = useList<ContactDefault>('/user/contact/default', (t, v) => t.id === v.id);

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ja: 'お問い合わせカスタム', en: 'Custom Contact'})}
      </Heading>
      {error ? (
        <Center
          fontSize="1.5rem"
          fontWeight="bold"
          color="red.400"
          textAlign="center"
        >
          {error.message}
        </Center>
      ) : (
        <Box
          mx={{base: '.5rem', sm: '1.5rem', md: '0'}}
          display={{base: 'block', md: 'flex'}}
          alignItems="center"
          flexDirection="column"
        >
          <Box width={{base: 'auto', md: '700px'}}>
            <Flex>
              <Back href="/admin" />
              <Spacer />
              <Button my="1rem" onClick={createModal.onOpen}>
                {convertLang({ja: '新規作成', en: 'New'})}
              </Button>
            </Flex>
            <TableContainer>
              <Table variant="simple">
                <Thead>
                  <Tr>
                    <Th>{convertLang({ja: 'ID', en: 'ID'})}</Th>
                    <Th>{convertLang({ja: '名前', en: 'Name'})}</Th>
                    <Th>{convertLang({ja: 'メール', en: 'Email'})}</Th>
                    <Th>{convertLang({ja: 'URL', en: 'URL'})}</Th>
                    <Th>{convertLang({ja: 'カテゴリ', en: 'Category'})}</Th>
                    <Th>
                      {convertLang({
                        ja: 'カスタムタイトル',
                        en: 'Custom Title',
                      })}
                    </Th>
                    <Th>
                      {convertLang({
                        ja: '詳細',
                        en: 'Description',
                      })}
                    </Th>
                    <Th></Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map(v => {
                    return (
                      <Tr key={v.id}>
                        <Td>{v.id}</Td>
                        <Td>{v.name}</Td>
                        <Td>{v.email}</Td>
                        <Td>{v.url}</Td>
                        <Td>{v.category}</Td>
                        <Td>{v.custom_title}</Td>
                        <Td
                          whiteSpace="nowrap"
                          textOverflow="ellipsis"
                          overflow="hidden"
                          maxWidth="300px"
                        >
                          {v.description}
                        </Td>
                        <Td>
                          <Button size="sm" onClick={() => onUpdate(v)}>
                            {convertLang({ja: '編集', en: 'Edit'})}
                          </Button>
                        </Td>
                      </Tr>
                    );
                  })}
                </Tbody>
              </Table>
            </TableContainer>
          </Box>
        </Box>
      )}
      <NewContactDefault
        convertLang={convertLang}
        isOpen={createModal.isOpen}
        onClose={createModal.onClose}
        update={update}
      />
      <UpdateContactDefault
        convertLang={convertLang}
        isOpen={updateModal.isOpen}
        onClose={closeUpdateModal}
        update={update}
        target={updateValue}
      />
    </Box>
  );
};

export const NewContactDefault: React.FC<{
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: ContactDefault, type: 'cre' | 'del' | 'upd') => void;
}> = ({convertLang, isOpen, onClose, update}) => {
  const {
    onSubmit,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    wrapperOnClose,
    methods,
  } = useNew<ContactDefault, ContactDefaultForm>({
    path: '/user/contact/default',
    formFunc: d => {
      const form = new FormData();

      if (d.name) {
        form.append('name', d.name);
      }
      if (d.email) {
        form.append('email', d.email);
      }
      if (d.url) {
        form.append('url', d.url);
      }
      if (d.category) {
        form.append('category', d.category);
      }
      if (d.custom_title) {
        form.append('custom_title', d.custom_title);
      }
      if (d.description) {
        form.append('description', d.description);
      }

      return form;
    },
    convertLang,
    parse: r => {
      return ContactDefaultSchema.parse(r);
    },
    update,
    onClose,
  });

  return (
    <Modal isOpen={isOpen} onClose={wrapperOnClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          {convertLang({
            ja: 'お問い合わせを新規作成',
            en: 'Create new Contacts',
          })}
        </ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody mb="1rem">
          <FormProvider {...methods}>
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
              <FormControl mt=".5rem" isInvalid={Boolean(errors.email)}>
                <FormLabel htmlFor="email">
                  {convertLang({ja: 'メールアドレス', en: 'Email Address'})}
                </FormLabel>
                <Input
                  id="email"
                  {...register('email', {
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
                  {errors.email && errors.email.message}
                </FormErrorMessage>
              </FormControl>
              <FormControl mt=".5rem" isInvalid={Boolean(errors.url)}>
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
              <FormControl mt=".5rem" isInvalid={Boolean(errors.category)}>
                <FormLabel htmlFor="category">
                  {convertLang({ja: 'カテゴリ', en: 'Category'})}
                </FormLabel>
                <Input id="category" {...register('category')} />
                <FormErrorMessage>
                  {errors.category && errors.category.message}
                </FormErrorMessage>
              </FormControl>
              <FormControl mt=".5rem" isInvalid={Boolean(errors.custom_title)}>
                <FormLabel htmlFor="custom_title">
                  {convertLang({
                    ja: 'カスタムタイトル',
                    en: 'Custom Title',
                  })}
                </FormLabel>
                <Input id="custom_title" {...register('custom_title')} />
                <FormErrorMessage>
                  {errors.custom_title && errors.custom_title.message}
                </FormErrorMessage>
              </FormControl>
              <FormControl isInvalid={Boolean(errors.description)} mt=".5rem">
                <FormLabel htmlFor="description">
                  {convertLang({
                    ja: '詳細説明',
                    en: 'Description',
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
              <Button
                mt={4}
                w="100%"
                colorScheme="cateiru"
                isLoading={isSubmitting}
                type="submit"
              >
                {convertLang({ja: '登録', en: 'Submit'})}
              </Button>
            </form>
          </FormProvider>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

const UpdateContactDefault: React.FC<{
  target: ContactDefault | undefined;
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: ContactDefault, type: 'cre' | 'del' | 'upd') => void;
}> = ({convertLang, target, isOpen, onClose, update}) => {
  const {
    onSubmit,
    wrapperOnClose,
    onDelete,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    methods,
  } = useUpdate<ContactDefault, ContactDefaultForm>({
    convertLang,
    path: '/user/contact/default',
    deleteIdName: 'id',
    formFunc: (d, id) => {
      const form = new FormData();

      form.append('id', String(id));

      form.append('name', d.name ?? '');
      form.append('email', d.email ?? '');
      form.append('url', d.url ?? '');
      form.append('category', d.category ?? '');
      form.append('custom_title', d.custom_title ?? '');
      form.append('description', d.description ?? '');

      return [form, true];
    },
    parse: r => {
      return ContactDefaultSchema.parse(r);
    },
    update,
    onClose,
    target,
    targetId: t => t.id,
    setValues: (t, setValue) => {
      t.name && setValue('name', t.name);
      t.email && setValue('email', t.email);
      t.url && setValue('url', t.url);
      t.category && setValue('category', t.category);
      t.custom_title && setValue('custom_title', t.custom_title);
      t.description && setValue('description', t.description);
    },
  });

  return (
    <Modal isOpen={isOpen} onClose={wrapperOnClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          {convertLang({ja: 'カテゴリを編集', en: 'Edit new Location'})}
        </ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody mb="1rem">
          <Flex justifyContent="right">
            <Button
              colorScheme="red"
              variant="outline"
              size="sm"
              onClick={onDelete}
            >
              {convertLang({ja: '削除', en: 'Delete'})}
            </Button>
          </Flex>
          <FormProvider {...methods}>
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
              <FormControl mt=".5rem" isInvalid={Boolean(errors.email)}>
                <FormLabel htmlFor="email">
                  {convertLang({ja: 'メールアドレス', en: 'Email Address'})}
                </FormLabel>
                <Input
                  id="email"
                  {...register('email', {
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
                  {errors.email && errors.email.message}
                </FormErrorMessage>
              </FormControl>
              <FormControl mt=".5rem" isInvalid={Boolean(errors.url)}>
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
              <FormControl mt=".5rem" isInvalid={Boolean(errors.category)}>
                <FormLabel htmlFor="category">
                  {convertLang({ja: 'カテゴリ', en: 'Category'})}
                </FormLabel>
                <Input id="category" {...register('category')} />
                <FormErrorMessage>
                  {errors.category && errors.category.message}
                </FormErrorMessage>
              </FormControl>
              <FormControl mt=".5rem" isInvalid={Boolean(errors.custom_title)}>
                <FormLabel htmlFor="custom_title">
                  {convertLang({
                    ja: 'カスタムタイトル',
                    en: 'Custom Title',
                  })}
                </FormLabel>
                <Input id="custom_title" {...register('custom_title')} />
                <FormErrorMessage>
                  {errors.custom_title && errors.custom_title.message}
                </FormErrorMessage>
              </FormControl>
              <FormControl isInvalid={Boolean(errors.description)} mt=".5rem">
                <FormLabel htmlFor="description">
                  {convertLang({
                    ja: '詳細説明',
                    en: 'Description',
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

              <Button
                mt={4}
                w="100%"
                colorScheme="cateiru"
                isLoading={isSubmitting}
                type="submit"
              >
                {convertLang({ja: '更新', en: 'Update'})}
              </Button>
            </form>
          </FormProvider>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
