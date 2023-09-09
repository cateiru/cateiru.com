import {
  Badge,
  Box,
  Button,
  Center,
  Flex,
  Heading,
  Spacer,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalCloseButton,
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  Select,
} from '@chakra-ui/react';
import React from 'react';
import {MultiLang} from '../../utils/config/lang';
import {Location, LocationSchema} from '../../utils/types';
import {Back} from '../Back';
import useLanguage from '../useLanguage';
import {useList} from './useList';
import {useNew} from './useNew';
import {useUpdate} from './useUpdate';

interface LocationForm {
  type: string;
  name: string;
  name_ja: string;
  address: string;
  address_ja: string;
}

export const LocationEdit = () => {
  const {lang, convertLang} = useLanguage();
  const {
    data,
    error,
    update,
    updateValue,
    onUpdate,
    closeUpdateModal,
    createModal,
    updateModal,
  } = useList<Location>('/user/location', (t, v) => t.id === v.id);

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ja: '場所編集', en: 'Location Edit'})}
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
                    <Th>{convertLang({ja: '名称', en: 'Name'})}</Th>
                    <Th>{convertLang({ja: '住所', en: 'Address'})}</Th>
                    <Th>{convertLang({ja: 'タイプ', en: 'Type'})}</Th>
                    <Th></Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map(v => {
                    return (
                      <Tr key={v.id}>
                        <Td>{lang === 'ja' ? v.name_ja : v.name}</Td>
                        <Td>{lang === 'ja' ? v.address_ja : v.address}</Td>
                        <Td>
                          <Badge
                            colorScheme={v.type === 'univ' ? 'blue' : 'green'}
                          >
                            {v.type}
                          </Badge>
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
      <NewForm
        isOpen={createModal.isOpen}
        onClose={createModal.onClose}
        convertLang={convertLang}
        update={update}
      />

      <UpdateForm
        isOpen={updateModal.isOpen}
        onClose={closeUpdateModal}
        convertLang={convertLang}
        target={updateValue}
        update={update}
      />
    </Box>
  );
};

const NewForm: React.FC<{
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: Location, type: 'cre' | 'del' | 'upd') => void;
}> = ({convertLang, isOpen, onClose, update}) => {
  const {
    onSubmit,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    wrapperOnClose,
  } = useNew<Location, LocationForm>({
    path: '/user/location',
    formFunc: d => {
      const form = new FormData();
      form.append('type', d.type);
      form.append('name', d.name);
      form.append('name_ja', d.name_ja);
      form.append('address', d.address);
      form.append('address_ja', d.address_ja);
      return form;
    },
    convertLang,
    parse: r => LocationSchema.parse(r),
    update,
    onClose,
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
          <form onSubmit={handleSubmit(onSubmit)}>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.type)}>
              <FormLabel htmlFor="type">
                {convertLang({ja: '種類', en: 'Type'})}
              </FormLabel>

              <Select
                id="type"
                placeholder={convertLang({ja: '種類を選択', en: 'Select Type'})}
                {...register('type', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              >
                <option value="univ">
                  {convertLang({ja: '大学', en: 'University'})}
                </option>
                <option value="corp">
                  {convertLang({ja: '会社', en: 'Corporation'})}
                </option>
              </Select>
              <FormErrorMessage>
                {errors.type && errors.type.message}
              </FormErrorMessage>
            </FormControl>
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
            <FormControl mt=".5rem" isInvalid={Boolean(errors.address)}>
              <FormLabel htmlFor="address">
                {convertLang({ja: '住所(英語)', en: 'Address'})}
              </FormLabel>
              <Input
                id="address"
                {...register('address', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.address && errors.address.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.address_ja)}>
              <FormLabel htmlFor="address_ja">
                {convertLang({ja: '住所', en: 'Address(japanese)'})}
              </FormLabel>
              <Input
                id="address_ja"
                {...register('address_ja', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.address_ja && errors.address_ja.message}
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
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

const UpdateForm: React.FC<{
  target: Location | undefined;
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: Location, type: 'cre' | 'del' | 'upd') => void;
}> = ({convertLang, target, isOpen, onClose, update}) => {
  const {
    onSubmit,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    wrapperOnClose,
    onDelete,
  } = useUpdate<Location, LocationForm>({
    formFunc: (d, id) => {
      const form = new FormData();
      form.append('location_id', `${id}`);
      let changed = false;
      if (d.type !== target?.type) {
        form.append('type', d.type);
        changed = true;
      }
      if (d.name !== target?.name) {
        form.append('name', d.name);
        changed = true;
      }
      if (d.name_ja !== target?.name_ja) {
        form.append('name_ja', d.name_ja);
        changed = true;
      }
      if (d.address !== target?.address) {
        form.append('address', d.address);
        changed = true;
      }
      if (d.address_ja !== target?.address_ja) {
        form.append('address_ja', d.address_ja);
        changed = true;
      }
      return [form, changed];
    },
    parse: r => LocationSchema.parse(r),
    update,
    onClose,
    target,
    targetId: t => t.id,
    setValues: (t, setValue) => {
      setValue('type', t.type);
      setValue('name', t.name);
      setValue('name_ja', t.name_ja);
      setValue('address', t.address);
      setValue('address_ja', t.address_ja);
    },
    path: '/user/location',
    deleteIdName: 'location_id',
    convertLang,
  });

  return (
    <Modal isOpen={isOpen} onClose={wrapperOnClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          {convertLang({ja: '場所を編集', en: 'Edit Location'})}
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
            <FormControl mt=".5rem" isInvalid={Boolean(errors.type)}>
              <FormLabel htmlFor="type">
                {convertLang({ja: '種類', en: 'Type'})}
              </FormLabel>

              <Select
                id="type"
                placeholder={convertLang({ja: '種類を選択', en: 'Select Type'})}
                {...register('type', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              >
                <option value="univ">
                  {convertLang({ja: '大学', en: 'University'})}
                </option>
                <option value="corp">
                  {convertLang({ja: '会社', en: 'Corporation'})}
                </option>
              </Select>
              <FormErrorMessage>
                {errors.type && errors.type.message}
              </FormErrorMessage>
            </FormControl>
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
            <FormControl mt=".5rem" isInvalid={Boolean(errors.address)}>
              <FormLabel htmlFor="address">
                {convertLang({ja: '住所(英語)', en: 'Address'})}
              </FormLabel>
              <Input
                id="address"
                {...register('address', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.address && errors.address.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.address_ja)}>
              <FormLabel htmlFor="address_ja">
                {convertLang({ja: '住所', en: 'Address(japanese)'})}
              </FormLabel>
              <Input
                id="address_ja"
                {...register('address_ja', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.address_ja && errors.address_ja.message}
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
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
