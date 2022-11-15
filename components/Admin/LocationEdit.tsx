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
  useDisclosure,
  FormControl,
  FormLabel,
  Input,
  FormErrorMessage,
  Select,
  useToast,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {useForm} from 'react-hook-form';
import {IoArrowBack} from 'react-icons/io5';
import useSWR from 'swr';
import {api} from '../../utils/api';
import {MultiLang} from '../../utils/config/lang';
import {fetcher, SWRError} from '../../utils/swr';
import {Location, LocationArray, LocationSchema} from '../../utils/types';
import useLanguage from '../useLanguage';

interface LocationForm {
  type: string;
  name: string;
  name_ja: string;
  address: string;
  address_ja: string;
}

export const LocationEdit = () => {
  const [lang, convertLang] = useLanguage();
  const {data, error, mutate} = useSWR<LocationArray, SWRError>(
    '/user/location',
    fetcher
  );
  const createModal = useDisclosure();
  const updateModal = useDisclosure();

  const [updateLocation, setUpdateLocation] = React.useState<Location>();

  const updateLocationHand = (l: Location, type: 'cre' | 'del' | 'upd') => {
    switch (type) {
      case 'cre':
      case 'upd':
        if (l) {
          if (data) {
            const d = [...data];
            const inLocationIndex = d.findIndex(v => v.id === l.id);
            if (inLocationIndex !== -1) {
              d[inLocationIndex] = l;
            } else {
              d.push(l);
            }
            mutate(d);
          } else {
            mutate([l]);
          }
        }
        break;
      case 'del':
        if (data) {
          const d = [...data];
          const inLocationIndex = d.findIndex(v => v.id === l.id);
          const deletedLocData = d.slice(inLocationIndex, inLocationIndex + 1);
          mutate(deletedLocData);
        }
        break;
    }
  };

  const edit = (id: number) => {
    const l = data?.find(v => v.id === id);
    if (l) {
      setUpdateLocation(l);
      updateModal.onOpen();
    }
  };

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
          {error.status} : {error.message}
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
                          <Button size="sm" onClick={() => edit(v.id)}>
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
        update={updateLocationHand}
      />
      {updateLocation && (
        <UpdateForm
          isOpen={updateModal.isOpen}
          onClose={updateModal.onClose}
          convertLang={convertLang}
          target={updateLocation}
          update={updateLocationHand}
        />
      )}
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
    handleSubmit,
    register,
    reset,
    formState: {errors, isSubmitting},
  } = useForm<LocationForm>();
  const toast = useToast();

  const onSubmit = async (d: LocationForm) => {
    const form = new FormData();
    form.append('type', d.type);
    form.append('name', d.name);
    form.append('name_ja', d.name_ja);
    form.append('address', d.address);
    form.append('address_ja', d.address_ja);

    const res = await fetch(api('/user/location'), {
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      body: form,
    });

    if (res.ok) {
      const location = LocationSchema.parse(await res.json());
      update(location, 'cre');
    } else {
      toast({
        status: 'error',
        title: (await res.json()).message,
      });
    }

    wrapperOnClose();

    return () => {};
  };

  const wrapperOnClose = () => {
    reset();
    onClose();
  };

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

const UpdateForm: React.FC<{
  target: Location;
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: Location, type: 'cre' | 'del' | 'upd') => void;
}> = ({convertLang, target, isOpen, onClose, update}) => {
  const {
    handleSubmit,
    register,
    reset,
    setValue,
    formState: {errors, isSubmitting},
  } = useForm<LocationForm>({});
  const toast = useToast();

  React.useEffect(() => {
    setValue('type', target.type);
    setValue('name', target.name);
    setValue('name_ja', target.name_ja);
    setValue('address', target.address);
    setValue('address_ja', target.address_ja);
  }, [target]);

  const onSubmit = async (d: LocationForm) => {
    const form = new FormData();
    form.append('location_id', `${target.id}`);
    let changed = false;
    if (d.type !== target.type) {
      form.append('type', d.type);
      changed = true;
    }
    if (d.name !== target.name) {
      form.append('name', d.name);
      changed = true;
    }
    if (d.name_ja !== target.name_ja) {
      form.append('name.name_ja', d.name_ja);
      changed = true;
    }
    if (d.address !== target.address) {
      form.append('address', d.address);
      changed = true;
    }
    if (d.address_ja !== target.address_ja) {
      form.append('address_ja', d.address_ja);
      changed = true;
    }

    if (!changed) {
      return;
    }

    const res = await fetch(api('/user/location'), {
      method: 'PUT',
      credentials: 'include',
      mode: 'cors',
      body: form,
    });

    if (res.ok) {
      const location = LocationSchema.parse(await res.json());
      update(location, 'upd');
    } else {
      toast({
        status: 'error',
        title: (await res.json()).message,
      });
    }

    wrapperOnClose();

    return () => {};
  };

  const onDelete = () => {
    const f = async () => {
      const res = await fetch(api(`/user/location?location_id=${target.id}`), {
        method: 'DELETE',
        credentials: 'include',
        mode: 'cors',
      });

      if (res.ok) {
        update(target, 'del');
        wrapperOnClose();
      } else {
        toast({
          status: 'error',
          title: (await res.json()).message,
        });
      }
    };

    f();
  };

  const wrapperOnClose = () => {
    reset();
    onClose();
  };

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
