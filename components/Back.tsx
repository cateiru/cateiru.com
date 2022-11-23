import {Button} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import {TbArrowBigLeft} from 'react-icons/tb';
import useLanguage from './useLanguage';

export const Back = () => {
  const {convertLang} = useLanguage();
  const router = useRouter();

  const handleClick = () => {
    router.back();
  };

  return (
    <Button
      my="1rem"
      variant="ghost"
      leftIcon={<TbArrowBigLeft size="20" />}
      onClick={handleClick}
    >
      {convertLang({ja: '戻る', en: 'Back'})}
    </Button>
  );
};
