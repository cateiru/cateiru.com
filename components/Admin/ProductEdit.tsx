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
  Image,
  Input,
  Link,
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
  Tooltip,
  Tr,
  useColorMode,
} from "@chakra-ui/react";
import type React from "react";
import { TbLink, TbBrandGithub } from "react-icons/tb";
import type { MultiLang } from "../../utils/config/lang";
import { parseDate } from "../../utils/parse";
import { type Product, ProductSchema } from "../../utils/types";
import { Back } from "../Back";
import useLanguage from "../useLanguage";
import { useList } from "./useList";
import { useNew } from "./useNew";
import { useUpdate } from "./useUpdate";

interface ProductForm {
  name: string;
  name_ja: string;
  detail: string;
  detail_ja: string;
  site_url: string;
  dev_time: string;
  github_url?: string;
  thumbnail?: string;
}

export const ProductEdit = () => {
  const { lang, convertLang } = useLanguage();
  const {
    data,
    error,
    update,
    updateValue,
    onUpdate,
    closeUpdateModal,
    createModal,
    updateModal,
  } = useList<Product>("/user/product", (t, v) => t.id === v.id);
  const { colorMode } = useColorMode();

  return (
    <Box mt="3rem">
      <Heading textAlign="center">
        {convertLang({ ja: "制作物編集", en: "Product Edit" })}
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
          mx={{ base: ".5rem", sm: "1.5rem", md: "0" }}
          display={{ base: "block", md: "flex" }}
          alignItems="center"
          flexDirection="column"
        >
          <Box width={{ base: "auto", md: "700px" }}>
            <Flex>
              <Back href="/admin" />
              <Spacer />
              <Button my="1rem" onClick={createModal.onOpen}>
                {convertLang({ ja: "新規作成", en: "New" })}
              </Button>
            </Flex>
            <TableContainer>
              <Table variant="simple">
                <Thead>
                  <Tr>
                    <Th></Th>
                    <Th>{convertLang({ ja: "名前", en: "Name" })}</Th>
                    <Th>{convertLang({ ja: "詳細", en: "Detail" })}</Th>
                    <Th>{convertLang({ ja: "開発日時", en: "Dev Time" })}</Th>
                    <Th></Th>
                    <Th></Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map((v) => {
                    return (
                      <Tr key={v.id}>
                        <Td p="0">
                          <Box width="50px">
                            {v.thumbnail && (
                              <Image
                                src={v.thumbnail}
                                alt="thumbnail"
                                width="100px"
                              />
                            )}
                          </Box>
                        </Td>
                        <Td>{convertLang({ ja: v.name_ja, en: v.name })}</Td>
                        <Td
                          whiteSpace="nowrap"
                          textOverflow="ellipsis"
                          overflow="hidden"
                          maxWidth="300px"
                        >
                          {convertLang({ ja: v.detail_ja, en: v.detail })}
                        </Td>
                        <Td>{parseDate(v.dev_time, lang)}</Td>
                        <Td>
                          <Box>
                            <Tooltip
                              label={v.site_url}
                              color={colorMode === "dark" ? "black" : "white"}
                            >
                              <IconButton
                                icon={<TbLink size="20" />}
                                aria-label="site url"
                                variant="ghost"
                                as={Link}
                                href={v.site_url}
                                isExternal
                              />
                            </Tooltip>
                            {v.github_url && (
                              <Tooltip
                                label={v.github_url}
                                color={colorMode === "dark" ? "black" : "white"}
                              >
                                <IconButton
                                  icon={<TbBrandGithub size="20" />}
                                  aria-label="site url"
                                  variant="ghost"
                                  as={Link}
                                  href={v.github_url}
                                  isExternal
                                />
                              </Tooltip>
                            )}
                          </Box>
                        </Td>
                        <Td>
                          <Button size="sm" onClick={() => onUpdate(v)}>
                            {convertLang({ ja: "編集", en: "Edit" })}
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
      <NewProduct
        convertLang={convertLang}
        isOpen={createModal.isOpen}
        onClose={createModal.onClose}
        update={update}
      />
      <UpdateProduct
        convertLang={convertLang}
        isOpen={updateModal.isOpen}
        onClose={closeUpdateModal}
        update={update}
        target={updateValue}
      />
    </Box>
  );
};

export const NewProduct: React.FC<{
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: Product, type: "cre" | "del" | "upd") => void;
}> = ({ convertLang, isOpen, onClose, update }) => {
  const {
    onSubmit,
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
    wrapperOnClose,
  } = useNew<Product, ProductForm>({
    path: "/user/product",
    formFunc: (d) => {
      const form = new FormData();
      form.append("name", d.name);
      form.append("name_ja", d.name_ja);
      form.append("detail", d.detail);
      form.append("detail_ja", d.detail_ja);
      form.append("site_url", d.site_url);
      form.append("dev_time", new Date(d.dev_time).toISOString());

      if (d.github_url) {
        form.append("github_url", d.github_url);
      }
      if (d.thumbnail) {
        form.append("thumbnail", d.thumbnail);
      }
      return form;
    },
    convertLang,
    parse: (r) => {
      return ProductSchema.parse(r);
    },
    update,
    onClose,
  });

  return (
    <Modal isOpen={isOpen} onClose={wrapperOnClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          {convertLang({ ja: "制作物を新規作成", en: "Create new Product" })}
        </ModalHeader>
        <ModalCloseButton size="lg" />
        <ModalBody mb="1rem">
          <form onSubmit={handleSubmit(onSubmit)}>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.name)}>
              <FormLabel htmlFor="name">
                {convertLang({ ja: "名称(英語)", en: "Name" })}
              </FormLabel>
              <Input
                id="name"
                {...register("name", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.name && errors.name.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.name_ja)}>
              <FormLabel htmlFor="name_ja">
                {convertLang({ ja: "名称", en: "Name(japanese)" })}
              </FormLabel>
              <Input
                id="name_ja"
                {...register("name_ja", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.name_ja && errors.name_ja.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.detail)}>
              <FormLabel htmlFor="detail">
                {convertLang({ ja: "詳細(英語)", en: "Detail" })}
              </FormLabel>
              <Textarea
                id="detail"
                {...register("detail", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.detail && errors.detail.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.detail_ja)}>
              <FormLabel htmlFor="detail_ja">
                {convertLang({ ja: "詳細", en: "Detail(japanese)" })}
              </FormLabel>
              <Textarea
                id="detail_ja"
                {...register("detail_ja", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.detail_ja && errors.detail_ja.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.site_url)}>
              <FormLabel htmlFor="site_url">
                {convertLang({ ja: "サイトURL", en: "Site URL" })}
              </FormLabel>
              <Input
                id="site_url"
                placeholder="https://"
                {...register("site_url", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                  pattern: {
                    value: /https?:\/\/([\da-z.-]+).([a-z.]{2,6})[/\w .-]*\/?/,
                    message: convertLang({
                      ja: "正しいURLを入力してください",
                      en: "Please enter the correct URL",
                    }),
                  },
                })}
              />
              <FormErrorMessage>
                {errors.site_url && errors.site_url.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.github_url)}>
              <FormLabel htmlFor="github_url">
                {convertLang({ ja: "GitHub URL", en: "GitHub URL" })}
              </FormLabel>
              <Input id="github_url" {...register("github_url")} />
              <FormErrorMessage>
                {errors.github_url && errors.github_url.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.dev_time)}>
              <FormLabel htmlFor="dev_time">
                {convertLang({ ja: "開発日時", en: "Dev Time" })}
              </FormLabel>
              <Input
                id="dev_time"
                type="date"
                {...register("dev_time", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.dev_time && errors.dev_time.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.thumbnail)}>
              <FormLabel htmlFor="thumbnail">
                {convertLang({ ja: "サムネイルURL", en: "Thumbnail URL" })}
              </FormLabel>
              <Input
                id="thumbnail"
                placeholder="https://"
                {...register("thumbnail", {
                  pattern: {
                    value: /https?:\/\/([\da-z.-]+).([a-z.]{2,6})[/\w .-]*\/?/,
                    message: convertLang({
                      ja: "正しいURLを入力してください",
                      en: "Please enter the correct URL",
                    }),
                  },
                })}
              />
              <FormErrorMessage>
                {errors.thumbnail && errors.thumbnail.message}
              </FormErrorMessage>
            </FormControl>
            <Button
              mt={4}
              w="100%"
              colorScheme="cateiru"
              isLoading={isSubmitting}
              type="submit"
            >
              {convertLang({ ja: "登録", en: "Submit" })}
            </Button>
          </form>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

const UpdateProduct: React.FC<{
  target: Product | undefined;
  convertLang: (e: MultiLang) => string;
  isOpen: boolean;
  onClose: () => void;
  update: (newV: Product, type: "cre" | "del" | "upd") => void;
}> = ({ convertLang, target, isOpen, onClose, update }) => {
  const {
    onSubmit,
    wrapperOnClose,
    onDelete,
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useUpdate<Product, ProductForm>({
    convertLang,
    path: "/user/product",
    deleteIdName: "product_id",
    formFunc: (d, id) => {
      const form = new FormData();
      form.append("product_id", `${id}`);
      if (d.name !== target?.name) {
        form.append("name", d.name);
      }
      if (d.name_ja !== target?.name_ja) {
        form.append("name_ja", d.name_ja);
      }
      if (d.detail !== target?.detail) {
        form.append("detail", d.detail);
      }
      if (d.detail_ja !== target?.detail_ja) {
        form.append("detail_ja", d.detail_ja);
      }
      if (d.site_url !== target?.site_url) {
        form.append("site_url", d.site_url);
      }
      if (
        new Date(d.dev_time).toISOString() !==
        new Date(target?.dev_time ?? "").toISOString()
      ) {
        form.append("dev_time", new Date(d.dev_time).toISOString());
      }
      form.append("github_url", d.github_url ?? "");
      form.append("thumbnail", d.thumbnail ?? "");

      return [form, true];
    },
    parse: (r) => {
      return ProductSchema.parse(r);
    },
    update,
    onClose,
    target,
    targetId: (t) => t.id,
    setValues: (t, setValue) => {
      setValue("name", t.name);
      setValue("name_ja", t.name_ja);
      setValue("detail", t.detail);
      setValue("detail_ja", t.detail_ja);
      setValue("site_url", t.site_url);
      setValue("dev_time", new Date(t.dev_time).toISOString().substring(0, 10));

      if (t.github_url) {
        setValue("github_url", t.github_url);
      }
      if (t.thumbnail) {
        setValue("thumbnail", t.thumbnail);
      }
    },
  });

  return (
    <Modal isOpen={isOpen} onClose={wrapperOnClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          {convertLang({ ja: "場所を新規作成", en: "Create new Location" })}
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
              {convertLang({ ja: "削除", en: "Delete" })}
            </Button>
          </Flex>
          <form onSubmit={handleSubmit(onSubmit)}>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.name)}>
              <FormLabel htmlFor="name">
                {convertLang({ ja: "名称(英語)", en: "Name" })}
              </FormLabel>
              <Input
                id="name"
                {...register("name", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.name && errors.name.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.name_ja)}>
              <FormLabel htmlFor="name_ja">
                {convertLang({ ja: "名称", en: "Name(japanese)" })}
              </FormLabel>
              <Input
                id="name_ja"
                {...register("name_ja", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.name_ja && errors.name_ja.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.detail)}>
              <FormLabel htmlFor="detail">
                {convertLang({ ja: "詳細(英語)", en: "Detail" })}
              </FormLabel>
              <Textarea
                id="detail"
                {...register("detail", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.detail && errors.detail.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.detail_ja)}>
              <FormLabel htmlFor="detail_ja">
                {convertLang({ ja: "詳細", en: "Detail(japanese)" })}
              </FormLabel>
              <Textarea
                id="detail_ja"
                {...register("detail_ja", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.detail_ja && errors.detail_ja.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.site_url)}>
              <FormLabel htmlFor="site_url">
                {convertLang({ ja: "サイトURL", en: "Site URL" })}
              </FormLabel>
              <Input
                id="site_url"
                {...register("site_url", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.site_url && errors.site_url.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.github_url)}>
              <FormLabel htmlFor="github_url">
                {convertLang({
                  ja: "GitHub URL(オプション)",
                  en: "GitHub URL(Optional)",
                })}
              </FormLabel>
              <Input id="github_url" {...register("github_url")} />
              <FormErrorMessage>
                {errors.github_url && errors.github_url.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.dev_time)}>
              <FormLabel htmlFor="dev_time">
                {convertLang({ ja: "開発日時", en: "Dev Time" })}
              </FormLabel>
              <Input
                id="dev_time"
                type="date"
                {...register("dev_time", {
                  required: convertLang({
                    ja: "この項目は必須です",
                    en: "This is required",
                  }),
                })}
              />
              <FormErrorMessage>
                {errors.dev_time && errors.dev_time.message}
              </FormErrorMessage>
            </FormControl>
            <FormControl mt=".5rem" isInvalid={Boolean(errors.thumbnail)}>
              <FormLabel htmlFor="thumbnail">
                {convertLang({
                  ja: "サムネイルURL(オプション)",
                  en: "Thumbnail URL(Optional)",
                })}
              </FormLabel>
              <Input id="thumbnail" {...register("thumbnail")} />
              <FormErrorMessage>
                {errors.thumbnail && errors.thumbnail.message}
              </FormErrorMessage>
            </FormControl>
            <Button
              mt={4}
              w="100%"
              colorScheme="cateiru"
              isLoading={isSubmitting}
              type="submit"
            >
              {convertLang({ ja: "更新", en: "Update" })}
            </Button>
          </form>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
