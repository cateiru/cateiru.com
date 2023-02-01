import {
  Box,
  Button,
  Center,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  IconButton,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Select,
  Spacer,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tooltip,
  Tr,
  useColorMode,
  Link as ChakraLink,
  Image,
  Checkbox,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {TbAdjustmentsAlt} from 'react-icons/tb';
import {TbLink, TbPhoto} from 'react-icons/tb';
import useSWR from 'swr';
import {MultiLang} from '../../utils/config/lang';
import {fetcher, SWRError} from '../../utils/swr';
import {Category, LinkCategory, LinkCategorySchema} from '../../utils/types';
import {Back} from '../Back';
import useLanguage from '../useLanguage';
import {useList} from './useList';
import {useNew} from './useNew';
import {useUpdate} from './useUpdate';

interface LinkForm {
  name: string;
  name_ja: string;
  site_url: string;
  category_id: string;
  update_favicon: boolean;
}

export const LinkEdit = () => {
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
  } = useList<LinkCategory>('/user/link', (t, v) => t.link.id === v.link.id);
  const category = useSWR<Category[], SWRError>('/user/category', fetcher);
  const {colorMode} = useColorMode();

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ja: 'リンク編集', en: 'Link Edit'})}
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
                    <Th></Th>
                    <Th>{convertLang({ja: '名前', en: 'Name'})}</Th>
                    <Th>
                      {convertLang({ja: 'カテゴリ名', en: 'Category Name'})}
                    </Th>
                    <Th>{convertLang({ja: 'URL', en: 'URL'})}</Th>
                    <Th></Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map(v => {
                    return (
                      <Tr key={v.link.id}>
                        <Td>
                          {v.link.favicon_url ? (
                            <Center w="25px">
                              <Image
                                src={v.link.favicon_url}
                                width="25px"
                                alt="favicon"
                              />
                            </Center>
                          ) : (
                            <Center>
                              <TbPhoto size="25px" />
                            </Center>
                          )}
                        </Td>
                        <Td>
                          {convertLang({ja: v.link.name_ja, en: v.link.name})}
                        </Td>
                        <Td>
                          {convertLang({
                            ja: v.category.name_ja,
                            en: v.category.name,
                          })}
                        </Td>
                        <Td>
                          <Tooltip
                            label={v.link.site_url}
                            color={colorMode === 'dark' ? 'black' : 'white'}
                          >
                            <IconButton
                              icon={<TbLink size="20" />}
                              aria-label="site url"
                              variant="ghost"
                              as={ChakraLink}
                              href={v.link.site_url}
                              isExternal
                            />
                          </Tooltip>
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
      <NewLink
        convertLang={convertLang}
        isOpen={createModal.isOpen}
        onClose={createModal.onClose}
        update={update}
        categories={category.data}
      />
      <UpdateLink
        convertLang={convertLang}
        isOpen={updateModal.isOpen}
        onClose={closeUpdateModal}
        update={update}
        target={updateValue}
        categories={category.data}
      />
    </Box>
  );
};

export const NewLink: React.FC<{
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: LinkCategory, type: 'cre' | 'del' | 'upd') => void;
  categories: Category[] | undefined;
}> = ({convertLang, isOpen, onClose, update, categories}) => {
  const {
    onSubmit,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    wrapperOnClose,
  } = useNew<LinkCategory, LinkForm>({
    path: '/user/link',
    formFunc: d => {
      const form = new FormData();
      form.append('name', d.name);
      form.append('name_ja', d.name_ja);
      form.append('category_id', d.category_id);
      form.append('site_url', d.site_url);
      return form;
    },
    convertLang,
    parse: r => {
      return LinkCategorySchema.parse(r);
    },
    update,
    onClose,
  });

  return (
    <Modal isOpen={isOpen} onClose={wrapperOnClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          {convertLang({ja: '制作物を新規作成', en: 'Create new Product'})}
        </ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody mb="1rem">
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
            <FormControl mt=".5rem" isInvalid={Boolean(errors.category_id)}>
              <FormLabel htmlFor="category_id">
                {convertLang({ja: 'カテゴリ', en: 'Category'})}
              </FormLabel>

              <Flex>
                <Select
                  placeholder={convertLang({
                    ja: 'カテゴリを選択',
                    en: 'Select category',
                  })}
                  id="category_id"
                  {...register('category_id', {
                    required: convertLang({
                      ja: 'この項目は必須です',
                      en: 'This is required',
                    }),
                  })}
                >
                  {categories?.map(v => (
                    <option key={v.id} value={v.id}>
                      {convertLang({ja: v.name_ja, en: v.name})}
                    </option>
                  ))}
                </Select>
                <NextLink passHref href="/admin/category">
                  <IconButton
                    ml=".5rem"
                    icon={<TbAdjustmentsAlt size="20" />}
                    aria-label="location config"
                    as="a"
                  />
                </NextLink>
              </Flex>
              <FormErrorMessage>
                {errors.category_id && errors.category_id.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.site_url)}>
              <FormLabel htmlFor="site_url">
                {convertLang({ja: 'サイトURL', en: 'Site URL'})}
              </FormLabel>
              <Input
                id="site_url"
                placeholder="https://"
                {...register('site_url', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
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
                {errors.site_url && errors.site_url.message}
              </FormErrorMessage>
            </FormControl>

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
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

const UpdateLink: React.FC<{
  target: LinkCategory | undefined;
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: LinkCategory, type: 'cre' | 'del' | 'upd') => void;
  categories: Category[] | undefined;
}> = ({convertLang, target, isOpen, onClose, update, categories}) => {
  const {
    onSubmit,
    wrapperOnClose,
    onDelete,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
  } = useUpdate<LinkCategory, LinkForm>({
    convertLang,
    path: '/user/link',
    deleteIdName: 'link_id',
    formFunc: (d, id) => {
      const form = new FormData();
      form.append('link_id', `${id}`);
      let changed = false;
      if (d.name !== target?.link.name) {
        form.append('name', d.name);
        changed = true;
      }
      if (d.name_ja !== target?.link.name_ja) {
        form.append('name_ja', d.name_ja);
        changed = true;
      }
      if (d.category_id !== String(target?.link.category_id)) {
        form.append('category_id', d.category_id);
        changed = true;
      }
      if (d.site_url !== target?.link.site_url) {
        form.append('site_url', d.site_url);
        changed = true;
      }
      // site_urlが変更されておらず、update_faviconがtrueの場合はfaviconを更新
      if (d.update_favicon && d.site_url === target?.link.site_url) {
        form.append('update_favicon', 'true');
        changed = true;
      }

      return [form, changed];
    },
    parse: r => {
      return LinkCategorySchema.parse(r);
    },
    update,
    onClose,
    target,
    targetId: t => t.link.id,
    setValues: (t, setValue) => {
      setValue('name', t.link.name);
      setValue('name_ja', t.link.name_ja);
      setValue('category_id', String(t.link.category_id));
      setValue('site_url', t.link.site_url);
      setValue('update_favicon', false);
    },
  });

  return (
    <Modal isOpen={isOpen} onClose={wrapperOnClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          {convertLang({ja: '場所を新規作成', en: 'Create new Location'})}
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
            <FormControl mt=".5rem" isInvalid={Boolean(errors.category_id)}>
              <FormLabel htmlFor="category_id">
                {convertLang({ja: 'カテゴリ', en: 'Category'})}
              </FormLabel>

              <Flex>
                <Select
                  placeholder={convertLang({
                    ja: 'カテゴリを選択',
                    en: 'Select category',
                  })}
                  id="category_id"
                  {...register('category_id', {
                    required: convertLang({
                      ja: 'この項目は必須です',
                      en: 'This is required',
                    }),
                  })}
                >
                  {categories?.map(v => (
                    <option key={v.id} value={v.id}>
                      {convertLang({ja: v.name_ja, en: v.name})}
                    </option>
                  ))}
                </Select>
                <NextLink passHref href="/admin/category">
                  <IconButton
                    ml=".5rem"
                    icon={<TbAdjustmentsAlt size="20" />}
                    aria-label="location config"
                    as="a"
                  />
                </NextLink>
              </Flex>
              <FormErrorMessage>
                {errors.category_id && errors.category_id.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.site_url)}>
              <FormLabel htmlFor="site_url">
                {convertLang({ja: 'サイトURL', en: 'Site URL'})}
              </FormLabel>
              <Input
                id="site_url"
                placeholder="https://"
                {...register('site_url', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
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
                {errors.site_url && errors.site_url.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.update_favicon)}>
              <Checkbox
                id="update_favicon"
                placeholder="https://"
                {...register('update_favicon')}
              >
                {convertLang({ja: 'Favicon更新', en: 'Update Favicon'})}
              </Checkbox>
              <FormErrorMessage>
                {errors.update_favicon && errors.update_favicon.message}
              </FormErrorMessage>
            </FormControl>

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
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
