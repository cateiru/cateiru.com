import {
  Box,
  Center,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  useColorMode,
  Skeleton,
} from '@chakra-ui/react';
import EmojiPicker from 'emoji-picker-react';
import {Theme as EmojiTheme, EmojiStyle} from 'emoji-picker-react';
import React from 'react';
import {useFormContext} from 'react-hook-form';
import {MultiLang} from '../../utils/config/lang';

interface EmojiForm {
  emoji: string;
}

export const EmojiPick = React.memo<{
  convertLang: (e: MultiLang) => string;
}>(({convertLang}) => {
  const {
    register,
    setValue,
    getValues,
    formState: {errors},
  } = useFormContext<EmojiForm>();
  const {colorMode} = useColorMode();
  const [previewEmoji, setPreviewEmoji] = React.useState('');
  const [showEmojiSearch, setShowEmojiSearch] = React.useState(false);

  React.useEffect(() => {
    const e = getValues('emoji');
    if (e.length !== 0) {
      setPreviewEmoji(e);
    }

    setTimeout(() => {
      setShowEmojiSearch(true);
    }, 1000);
  }, []);

  const setEmoji = (emoji: string) => {
    setValue('emoji', emoji);
    setPreviewEmoji(emoji);
  };

  return (
    <FormControl mt=".5rem" isInvalid={Boolean(errors.emoji)}>
      <FormLabel htmlFor="emoji">
        {convertLang({ja: '絵文字', en: 'Emoji'})}
      </FormLabel>
      <Input
        as={Box}
        alignItems="center"
        display="flex"
        mb=".5rem"
        fontSize="1.5rem"
      >
        {previewEmoji}
      </Input>
      <Input
        display="none"
        id="emoji"
        {...register('emoji', {
          required: convertLang({
            ja: 'この項目は必須です',
            en: 'This is required',
          }),
        })}
      />
      <FormErrorMessage>
        {errors.emoji && errors.emoji.message}
      </FormErrorMessage>
      <Skeleton isLoaded={showEmojiSearch} h="450px">
        <Center>
          {showEmojiSearch && (
            <EmojiPicker
              onEmojiClick={e => setEmoji(e.emoji)}
              theme={colorMode === 'dark' ? EmojiTheme.DARK : EmojiTheme.LIGHT}
              emojiStyle={EmojiStyle.NATIVE}
              lazyLoadEmojis
              width="100%"
            />
          )}
        </Center>
      </Skeleton>
    </FormControl>
  );
});

EmojiPick.displayName = 'EmojiPick';
