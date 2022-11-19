import {
  Badge,
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
  Switch,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {useForm} from 'react-hook-form';
import {IoArrowBack} from 'react-icons/io5';
import {TbAdjustmentsAlt} from 'react-icons/tb';
import {TbCheck} from 'react-icons/tb';
import useSWR from 'swr';
import type {SWRResponse} from 'swr';
import {api} from '../../utils/api';
import {MultiLang} from '../../utils/config/lang';
import {parseDate} from '../../utils/parse';
import {fetcher, SWRError} from '../../utils/swr';
import {
  Bio,
  BioLoc,
  BioLocArray,
  BioLocSchema,
  BioSchema,
  LocationArray,
} from '../../utils/types';
import useLanguage from '../useLanguage';
import {useList} from './useList';
import {useNew} from './useNew';
import {useUpdate} from './useUpdate';

interface BioForm {
  is_public: boolean;
  location_id: string;
  position: string;
  position_ja: string;
  join_date: string;
  leave_date?: string;
}

export const BioList = () => {
  const [lang, convertLang] = useLanguage();
  const {
    data,
    error,
    mutate,
    update,
    updateValue,
    onUpdate,
    closeUpdateModal,
    createModal,
    updateModal,
  } = useList<BioLoc>('/user/bio', (t, v) => t.biography.id === v.biography.id);
  const locations = useSWR<LocationArray, SWRError>('/user/location', fetcher);

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ja: '略歴編集', en: 'Biography Edit'})}
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
                    <Th>{convertLang({ja: '場所', en: 'Location'})}</Th>
                    <Th>{convertLang({ja: '住所', en: 'Address'})}</Th>
                    <Th>{convertLang({ja: 'タイプ', en: 'Type'})}</Th>
                    <Th>{convertLang({ja: 'ポジション', en: 'Position'})}</Th>
                    <Th>{convertLang({ja: '参加', en: 'Join'})}</Th>
                    <Th>{convertLang({ja: '脱退', en: 'Leave'})}</Th>
                    <Th>{convertLang({ja: '公開', en: 'Public'})}</Th>
                    <Th></Th>
                    <Th></Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map(v => {
                    return (
                      <Tr key={v.biography.id}>
                        <Td>
                          {convertLang({
                            ja: v.location?.name_ja ?? '',
                            en: v.location?.name ?? '',
                          })}
                        </Td>
                        <Td>
                          {convertLang({
                            ja: v.location?.address_ja ?? '',
                            en: v.location?.address ?? '',
                          })}
                        </Td>
                        <Td>
                          <Badge
                            colorScheme={
                              v.location
                                ? v.location?.type === 'univ'
                                  ? 'blue'
                                  : 'green'
                                : 'gray'
                            }
                          >
                            {v.location?.type ?? 'NULL'}
                          </Badge>
                        </Td>
                        <Td>
                          {convertLang({
                            ja: v.biography.position_ja,
                            en: v.biography.position,
                          })}
                        </Td>
                        <Td>{parseDate(v.biography.join, lang)}</Td>
                        <Td>{parseDate(v.biography.leave, lang)}</Td>
                        <Td color="green.500">
                          {v.biography.is_public ? <TbCheck size="20" /> : ''}
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
      <NewBio
        convertLang={convertLang}
        isOpen={createModal.isOpen}
        onClose={createModal.onClose}
        update={update}
        locations={locations}
      />
      <UpdateBio
        convertLang={convertLang}
        isOpen={updateModal.isOpen}
        onClose={closeUpdateModal}
        update={update}
        target={updateValue}
        locations={locations}
      />
    </Box>
  );
};

export const NewBio: React.FC<{
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: BioLoc, type: 'cre' | 'del' | 'upd') => void;
  locations: SWRResponse<LocationArray>;
}> = ({convertLang, isOpen, onClose, update, locations}) => {
  const {
    onSubmit,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    wrapperOnClose,
  } = useNew<BioLoc, BioForm>({
    path: '/user/bio',
    formFunc: v => {
      const form = new FormData();
      form.append('is_public', v.is_public ? 'true' : 'false');
      form.append('location_id', v.location_id);
      form.append('position', v.position);
      form.append('position_ja', v.position_ja);
      form.append('join_date', new Date(v.join_date).toISOString());
      if (v.leave_date) {
        form.append('leave_date', new Date(v.leave_date).toISOString());
      }
      return form;
    },
    convertLang,
    parse: r => BioLocSchema.parse(r),
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
            <FormControl mt=".5rem" isInvalid={Boolean(errors.location_id)}>
              <FormLabel htmlFor="location_id">
                {convertLang({ja: '場所', en: 'Location'})}
              </FormLabel>
              <Flex>
                <Select
                  placeholder={convertLang({
                    ja: '場所を選択',
                    en: 'Select Location',
                  })}
                  {...register('location_id', {
                    required: convertLang({
                      ja: 'この項目は必須です',
                      en: 'This is required',
                    }),
                  })}
                >
                  {locations.data?.map(v => {
                    return (
                      <option value={v.id} key={v.id}>
                        {convertLang({ja: v.name_ja, en: v.name})}
                      </option>
                    );
                  })}
                </Select>
                <NextLink passHref href="/admin/location">
                  <IconButton
                    ml=".5rem"
                    icon={<TbAdjustmentsAlt size="20" />}
                    aria-label="location config"
                    as="a"
                  />
                </NextLink>
              </Flex>
              <FormErrorMessage>
                {errors.location_id && errors.location_id.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.position)}>
              <FormLabel htmlFor="position">
                {convertLang({ja: '立場・ポジション(英語)', en: 'Position'})}
              </FormLabel>
              <Input
                id="position"
                {...register('position', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.position && errors.position.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.position_ja)}>
              <FormLabel htmlFor="position_ja">
                {convertLang({
                  ja: '立場・ポジション',
                  en: 'Position(japanese)',
                })}
              </FormLabel>
              <Input
                id="position_ja"
                {...register('position_ja', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.position_ja && errors.position_ja.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.is_public)}>
              <FormLabel htmlFor="is_public">
                {convertLang({
                  ja: '公開設定',
                  en: 'Public',
                })}
              </FormLabel>
              <Switch id="is_public" {...register('is_public')} />
              <FormErrorMessage>
                {errors.is_public && errors.is_public.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.join_date)}>
              <FormLabel htmlFor="join_date">
                {convertLang({
                  ja: '参加日時',
                  en: 'Join Date',
                })}
              </FormLabel>
              <Input
                id="join_date"
                type="date"
                {...register('join_date', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.join_date && errors.join_date.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.leave_date)}>
              <FormLabel htmlFor="leave_date">
                {convertLang({
                  ja: '脱退',
                  en: 'Join Date',
                })}
              </FormLabel>
              <Input id="leave_date" type="date" {...register('leave_date')} />
              <FormErrorMessage>
                {errors.leave_date && errors.leave_date.message}
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

const UpdateBio: React.FC<{
  target: BioLoc | undefined;
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: BioLoc, type: 'cre' | 'del' | 'upd') => void;
  locations: SWRResponse<LocationArray>;
}> = ({convertLang, target, isOpen, onClose, update, locations}) => {
  const {
    onSubmit,
    handleSubmit,
    register,
    formState: {errors, isSubmitting},
    wrapperOnClose,
    onDelete,
  } = useUpdate<BioLoc, BioForm>({
    formFunc: (d, id) => {
      const form = new FormData();
      form.append('bio_id', `${id}`);
      let changed = false;
      if (d.is_public !== Boolean(target?.biography.is_public)) {
        form.append('is_public', d.is_public ? 'true' : 'false');
        changed = true;
      }
      if (d.location_id !== String(target?.biography.location_id)) {
        form.append('location_id', d.location_id);
        changed = true;
      }
      if (d.position !== target?.biography.position) {
        form.append('position', d.position);
        changed = true;
      }
      if (d.position_ja !== target?.biography.position_ja) {
        form.append('position_ja', d.position_ja);
        changed = true;
      }
      if (
        new Date(d.join_date).toISOString() !==
        new Date(target?.biography.join ?? '').toISOString()
      ) {
        form.append('join_date', new Date(d.join_date).toISOString());
        changed = true;
      }
      if (
        d.leave_date &&
        new Date(d.leave_date).toISOString() !==
          new Date(target?.biography.leave ?? '').toISOString()
      ) {
        form.append('leave_date', new Date(d.leave_date).toISOString());
        changed = true;
      }
      return [form, changed];
    },
    parse: r => BioLocSchema.parse(r),
    update,
    onClose,
    target,
    targetId: t => t.biography.id,
    setValues: (t, setValue) => {
      setValue('is_public', t.biography.is_public ?? false);
      setValue('location_id', String(t.biography.location_id));
      setValue('position', t.biography.position);
      setValue('position_ja', t.biography.position_ja);
      setValue(
        'join_date',
        new Date(t.biography.join ?? '').toISOString().substring(0, 10)
      );

      if (t.biography.leave !== '0001-01-01T00:00:00Z') {
        setValue(
          'leave_date',
          new Date(t.biography.leave ?? '').toISOString().substring(0, 10)
        );
      }
    },
    path: '/user/bio',
    deleteIdName: 'bio_id',
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
            <FormControl mt=".5rem" isInvalid={Boolean(errors.location_id)}>
              <FormLabel htmlFor="location_id">
                {convertLang({ja: '場所', en: 'Location'})}
              </FormLabel>
              <Flex>
                <Select
                  placeholder={convertLang({
                    ja: '場所を選択',
                    en: 'Select Location',
                  })}
                  {...register('location_id', {
                    required: convertLang({
                      ja: 'この項目は必須です',
                      en: 'This is required',
                    }),
                  })}
                >
                  {locations.data?.map(v => {
                    return (
                      <option value={v.id} key={v.id}>
                        {convertLang({ja: v.name_ja, en: v.name})}
                      </option>
                    );
                  })}
                </Select>
                <NextLink passHref href="/admin/location">
                  <IconButton
                    ml=".5rem"
                    icon={<TbAdjustmentsAlt size="20" />}
                    aria-label="location config"
                    as="a"
                  />
                </NextLink>
              </Flex>
              <FormErrorMessage>
                {errors.location_id && errors.location_id.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.position)}>
              <FormLabel htmlFor="position">
                {convertLang({ja: '立場・ポジション(英語)', en: 'Position'})}
              </FormLabel>
              <Input
                id="position"
                {...register('position', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.position && errors.position.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.position_ja)}>
              <FormLabel htmlFor="position_ja">
                {convertLang({
                  ja: '立場・ポジション',
                  en: 'Position(japanese)',
                })}
              </FormLabel>
              <Input
                id="position_ja"
                {...register('position_ja', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.position_ja && errors.position_ja.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.is_public)}>
              <FormLabel htmlFor="is_public">
                {convertLang({
                  ja: '公開設定',
                  en: 'Public',
                })}
              </FormLabel>
              <Switch id="is_public" {...register('is_public')} />
              <FormErrorMessage>
                {errors.is_public && errors.is_public.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.join_date)}>
              <FormLabel htmlFor="join_date">
                {convertLang({
                  ja: '参加日時',
                  en: 'Join Date',
                })}
              </FormLabel>
              <Input
                id="join_date"
                type="date"
                {...register('join_date', {
                  required: convertLang({
                    ja: 'この項目は必須です',
                    en: 'This is required',
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.join_date && errors.join_date.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.leave_date)}>
              <FormLabel htmlFor="leave_date">
                {convertLang({
                  ja: '脱退',
                  en: 'Join Date',
                })}
              </FormLabel>
              <Input id="leave_date" type="date" {...register('leave_date')} />
              <FormErrorMessage>
                {errors.leave_date && errors.leave_date.message}
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
