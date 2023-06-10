import {
  Button,
  Checkbox,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Image,
  Input,
  NumberDecrementStepper,
  NumberIncrementStepper,
  NumberInput,
  NumberInputField,
  NumberInputStepper,
  Select,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {Controller, FormProvider, useForm} from 'react-hook-form';

import useLanguage from '../../useLanguage';
import {ColorPickerForm, ColorPickerFormType} from './ColorPicker';
import {PreviewImage} from './Preview';
import {Transform, getSize} from './Transform';

// CORS制限があるので同一オリジンにホストする
const SOURCE_IMAGE = '/transparent.png';

interface GeneratorForm extends ColorPickerFormType {
  width?: number;
  height?: number;

  imageType: 'image/png' | 'image/jpeg' | 'image/webp';
  quality?: number;

  isCircle: boolean;

  image?: FileList;
}

export const Generator = () => {
  const methods = useForm<GeneratorForm>({
    defaultValues: {imageType: 'image/png'},
  });
  const {
    handleSubmit,
    register,
    watch,
    setValue,
    control,
    formState: {errors, isSubmitting},
  } = methods;
  const {convertLang} = useLanguage();
  const toast = useToast();

  const [generatedURL, setGeneratedURL] = React.useState<string>();
  const [extension, setExtension] = React.useState('');

  const handler = async (data: GeneratorForm) => {
    const t = new Transform(data.width, data.height);

    if (data.image && data.image[0]) {
      await t.loadImageFromSrc(URL.createObjectURL(data.image[0]));
    } else {
      await t.loadImageFromSrc(SOURCE_IMAGE);
    }

    if (data.isTransparent) {
      const r = data.r ?? 255;
      const g = data.g ?? 255;
      const b = data.b ?? 255;

      t.setBg(r, g, b);
    }
    if (data.isCircle) {
      t.setCircle();
    }

    const image = await t.export(
      data.imageType,
      Number.isNaN(data.quality) ? undefined : data.quality
    );

    if (image) {
      setGeneratedURL(URL.createObjectURL(image));
      setExtension(image.type.split('/')[1]);
    } else {
      toast({
        status: 'error',
        title: convertLang({
          ja: '画像の生成に失敗しました',
          en: 'Image generation failed.',
        }),
      });
    }
  };

  // 画像がアップロードされたか見たりするやつ
  const selectedImage = React.useCallback(() => {
    const files = watch('image');
    if (typeof files === 'undefined') return;
    const f = files[0];
    if (!f) return;
    return f;
  }, [watch]);

  // 画像をダウンロードする
  const download = () => {
    if (typeof generatedURL === 'undefined') return;

    const link = document.createElement('a');
    link.href = generatedURL;
    link.download = `icon.${extension}`;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  };

  React.useEffect(() => {
    setValue('width', 1000);
    setValue('height', 1000);
    console.log('reset');
  }, []);

  // 画像の幅高さのフォームを埋める
  React.useEffect(() => {
    const f = async () => {
      const files = watch('image');
      if (typeof files === 'undefined') return;
      const f = files[0];
      if (!f) return;

      const [width, height] = await getSize(f);
      setValue('width', width);
      setValue('height', height);
    };
    f();
  }, [watch('image')]);

  return (
    <>
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(handler)}>
          <FormControl mt="1rem" isInvalid={Boolean(errors.image)}>
            <FormLabel htmlFor="image">
              {convertLang({
                ja: 'アイコン（オプション）',
                en: 'Icon (Optional)',
              })}
            </FormLabel>
            <label htmlFor="image">
              {(() => {
                const i = selectedImage();

                if (!i) {
                  return (
                    <Button as="p" cursor="pointer">
                      {convertLang({ja: 'ファイルを選択', en: 'Select File'})}
                    </Button>
                  );
                }
                return (
                  <Image
                    src={URL.createObjectURL(i)}
                    alt="select image"
                    w="50px"
                    cursor="pointer"
                  />
                );
              })()}
            </label>
            <Input
              display="none"
              id="image"
              type="file"
              accept="image/*"
              {...register('image')}
            />
            <FormErrorMessage>
              {errors.image && errors.image.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl mt="1rem" isInvalid={Boolean(errors.width)}>
            <FormLabel htmlFor="width">
              {convertLang({ja: '画像幅 (px)', en: 'Image Width (px)'})}
            </FormLabel>
            <Controller
              name="width"
              control={control}
              rules={{
                min: {
                  value: 0,
                  message: convertLang({
                    ja: '0px以上で指定してください',
                    en: 'Please specify at least 0px.',
                  }),
                },
              }}
              render={({field: {ref, onChange, ...restField}}) => (
                <NumberInput
                  {...restField}
                  onChange={(_, text) => onChange(text)}
                  min={0}
                  step={10}
                >
                  <NumberInputField ref={ref} name={restField.name} />
                  <NumberInputStepper>
                    <NumberIncrementStepper />
                    <NumberDecrementStepper />
                  </NumberInputStepper>
                </NumberInput>
              )}
            />
            <FormErrorMessage>
              {errors.width && errors.width.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl mt="1rem" isInvalid={Boolean(errors.height)}>
            <FormLabel htmlFor="height">
              {convertLang({ja: '画像高さ (px)', en: 'Image Height (px)'})}
            </FormLabel>
            <Controller
              name="height"
              control={control}
              rules={{
                min: {
                  value: 0,
                  message: convertLang({
                    ja: '0px以上で指定してください',
                    en: 'Please specify at least 0px.',
                  }),
                },
              }}
              render={({field: {ref, onChange, ...restField}}) => (
                <NumberInput
                  {...restField}
                  onChange={(_, text) => onChange(text)}
                  min={0}
                  step={10}
                >
                  <NumberInputField ref={ref} name={restField.name} />
                  <NumberInputStepper>
                    <NumberIncrementStepper />
                    <NumberDecrementStepper />
                  </NumberInputStepper>
                </NumberInput>
              )}
            />
            <FormErrorMessage>
              {errors.height && errors.height.message}
            </FormErrorMessage>
          </FormControl>
          <ColorPickerForm />
          <FormControl mt="1rem" isInvalid={Boolean(errors.isCircle)}>
            <Checkbox id="isCircle" {...register('isCircle')}>
              {convertLang({
                ja: '円でマスクする',
                en: 'Masking with a circle',
              })}
            </Checkbox>
            <FormErrorMessage>
              {errors.isCircle && errors.isCircle.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl mt="1rem" isInvalid={Boolean(errors.imageType)}>
            <FormLabel htmlFor="imageType">
              {convertLang({ja: '画像形式', en: 'Image Type'})}
            </FormLabel>
            <Select
              placeholder={convertLang({
                ja: '画像形式を選択',
                en: 'Select image type',
              })}
              id="imageType"
              {...register('imageType', {
                required: convertLang({
                  ja: 'この項目は必須です',
                  en: 'This is required',
                }),
              })}
            >
              <option value="image/png">png</option>
              <option value="image/jpeg">jpeg</option>
              <option value="image/webp">webp</option>
            </Select>
            <FormErrorMessage>
              {errors.imageType && errors.imageType.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl mt="1rem" isInvalid={Boolean(errors.quality)}>
            <FormLabel htmlFor="quality">
              {convertLang({
                ja: 'クオリティ（jpegとwebpのみ）',
                en: 'Quality (jpeg and webp only)',
              })}
            </FormLabel>
            <NumberInput min={0} max={1} step={0.1}>
              <NumberInputField
                id="quality"
                {...register('quality', {
                  valueAsNumber: true,
                  max: {
                    value: 1.1,
                    message: convertLang({
                      ja: '1以内で指定してください',
                      en: '',
                    }),
                  },
                  min: {
                    value: 0,
                    message: convertLang({
                      ja: '0以上で指定してください',
                      en: '',
                    }),
                  },
                })}
              />
              <NumberInputStepper>
                <NumberIncrementStepper />
                <NumberDecrementStepper />
              </NumberInputStepper>
            </NumberInput>
            <FormErrorMessage>
              {errors.quality && errors.quality.message}
            </FormErrorMessage>
          </FormControl>
          <Button
            mt={4}
            w="100%"
            colorScheme="cateiru"
            isLoading={isSubmitting}
            type="submit"
            size="sm"
          >
            {convertLang({ja: '生成', en: 'Generate'})}
          </Button>
        </form>
      </FormProvider>
      {generatedURL && <PreviewImage src={generatedURL} download={download} />}
    </>
  );
};
