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
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import {Emoji} from 'emoji-picker-react';
import NextLink from 'next/link';
import React from 'react';
import {FormProvider} from 'react-hook-form';
import {IoArrowBack} from 'react-icons/io5';
import {MultiLang} from '../../utils/config/lang';
import {toUnicode} from '../../utils/parse';
import {Category, CategorySchema} from '../../utils/types';
import {EmojiPick} from '../Form/EmojiPick';
import useLanguage from '../useLanguage';
import {useList} from './useList';
import {useNew} from './useNew';
import {useUpdate} from './useUpdate';

interface CategoryForm {
  name: string;
  name_ja: string;
  emoji: string;
}

export const CategoryEdit = () => {
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
  } = useList<Category>('/user/category', (t, v) => t.id === v.id);

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ja: 'カテゴリ編集', en: 'Categories Edit'})}
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
              <NextLink passHref href="/admin">
                <Button
                  my="1rem"
                  variant="ghost"
                  leftIcon={<IoArrowBack />}
                  as="a"
                >
                  {convertLang({ja: '戻る', en: 'Back'})}
                </Button>
              </NextLink>
              <Spacer />
              <Button my="1rem" onClick={createModal.onOpen}>
                {convertLang({ja: '新規作成', en: 'New'})}
              </Button>
            </Flex>
            <TableContainer>
              <Table variant="simple">
                <Thead>
                  <Tr>
                    <Th>{convertLang({ja: '絵文字', en: 'Emoji'})}</Th>
                    <Th>{convertLang({ja: '名前', en: 'Name'})}</Th>
                    <Th></Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map(v => {
                    return (
                      <Tr key={v.id}>
                        <Td>
                          <Emoji
                            size={20}
                            unified={v.emoji ? toUnicode(v.emoji) || '' : ''}
                          />
                        </Td>
                        <Td>{convertLang({ja: v.name_ja, en: v.name})}</Td>
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
      <NewCategory
        convertLang={convertLang}
        isOpen={createModal.isOpen}
        onClose={createModal.onClose}
        update={update}
      />
      <UpdateCategory
        convertLang={convertLang}
        isOpen={updateModal.isOpen}
        onClose={closeUpdateModal}
        update={update}
        target={updateValue}
      />
    </Box>
  );
};

export const NewCategory: React.FC<{
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: Category, type: 'cre' | 'del' | 'upd') => void;
}> = ({convertLang, isOpen, onClose, update}) => {
  const {
    onSubmit,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    wrapperOnClose,
    methods,
  } = useNew<Category, CategoryForm>({
    path: '/user/category',
    formFunc: d => {
      const form = new FormData();
      form.append('name', d.name);
      form.append('name_ja', d.name_ja);
      form.append('emoji', d.emoji);

      return form;
    },
    convertLang,
    parse: r => {
      return CategorySchema.parse(r);
    },
    update,
    onClose,
  });

  return (
    <Modal isOpen={isOpen} onClose={wrapperOnClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          {convertLang({ja: 'カテゴリを新規作成', en: 'Create new Category'})}
        </ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody mb="1rem">
          <FormProvider {...methods}>
            <form onSubmit={handleSubmit(onSubmit)}>
              <FormControl mt=".5rem" isInvalid={Boolean(errors.name)}>
                <FormLabel htmlFor="name">
                  {convertLang({ja: '名称(英語)', en: 'Name'})}
                </FormLabel>
                <Input
                  id="name"
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
              <FormControl mt=".5rem" isInvalid={Boolean(errors.name_ja)}>
                <FormLabel htmlFor="name_ja">
                  {convertLang({ja: '名称', en: 'Name(japanese)'})}
                </FormLabel>
                <Input
                  id="name_ja"
                  {...register('name_ja', {
                    required: convertLang({
                      ja: 'この項目は必須です',
                      en: 'This is required',
                    }),
                  })}
                />
                <FormErrorMessage>
                  {errors.name_ja && errors.name_ja.message}
                </FormErrorMessage>
              </FormControl>
              <EmojiPick convertLang={convertLang} />
              <Button
                mt={4}
                w="100%"
                colorScheme="blue"
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

const UpdateCategory: React.FC<{
  target: Category | undefined;
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: Category, type: 'cre' | 'del' | 'upd') => void;
}> = ({convertLang, target, isOpen, onClose, update}) => {
  const {
    onSubmit,
    wrapperOnClose,
    onDelete,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    methods,
  } = useUpdate<Category, CategoryForm>({
    convertLang,
    path: '/user/category',
    deleteIdName: 'category_id',
    formFunc: (d, id) => {
      const form = new FormData();
      form.append('category_id', `${id}`);
      let changed = false;
      if (d.name !== target?.name) {
        form.append('name', d.name);
        changed = true;
      }
      if (d.name_ja !== target?.name_ja) {
        form.append('name_ja', d.name_ja);
        changed = true;
      }
      if (d.emoji !== target?.emoji) {
        form.append('emoji', d.emoji);
        changed = true;
      }

      return [form, changed];
    },
    parse: r => {
      return CategorySchema.parse(r);
    },
    update,
    onClose,
    target,
    targetId: t => t.id,
    setValues: (t, setValue) => {
      setValue('name', t.name);
      setValue('name_ja', t.name_ja);
      setValue('emoji', t.emoji);
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
              <FormControl mt=".5rem" isInvalid={Boolean(errors.name)}>
                <FormLabel htmlFor="name">
                  {convertLang({ja: '名称(英語)', en: 'Name'})}
                </FormLabel>
                <Input
                  id="name"
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
              <FormControl mt=".5rem" isInvalid={Boolean(errors.name_ja)}>
                <FormLabel htmlFor="name_ja">
                  {convertLang({ja: '名称', en: 'Name(japanese)'})}
                </FormLabel>
                <Input
                  id="name_ja"
                  {...register('name_ja', {
                    required: convertLang({
                      ja: 'この項目は必須です',
                      en: 'This is required',
                    }),
                  })}
                />
                <FormErrorMessage>
                  {errors.name_ja && errors.name_ja.message}
                </FormErrorMessage>
              </FormControl>
              <EmojiPick convertLang={convertLang} />

              <Button
                mt={4}
                w="100%"
                colorScheme="blue"
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
