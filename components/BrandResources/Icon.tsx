import {Box, Center, Heading, Image, Link} from '@chakra-ui/react';
import React from 'react';
import useLanguage from '../useLanguage';

export const Icon = () => {
  const {convertLang} = useLanguage();

  return (
    <Box mt="15vh">
      <Heading textAlign="center" mb="1rem">
        {convertLang({ja: 'アイコン', en: 'Icon'})}
      </Heading>
      <Center mx=".5rem">
        <IconContent
          src="https://storage.googleapis.com/cateiru/brand_resource/icon.png"
          alt="main icon"
        />
        <IconContent
          src="https://storage.googleapis.com/cateiru/brand_resource/bone.png"
          alt="bone icon"
          ml
        />
        <IconContent
          src="https://storage.googleapis.com/cateiru/brand_resource/transparent.png"
          alt="transparent icon"
          ml
        />
      </Center>
    </Box>
  );
};

const IconContent: React.FC<{
  src: string;
  alt: string;
  ml?: boolean;
}> = props => {
  return (
    <Link
      ml={props.ml ? {base: '.5rem', sm: '1.5rem', md: '2.5rem'} : undefined}
      href={props.src}
      isExternal
    >
      <Image
        src={props.src}
        w="200px"
        borderRadius="50%"
        bgColor="white"
        alt={props.alt}
        backgroundPosition="0px 0px, 7px 7px"
        backgroundSize="14px 14px"
        backgroundRepeat="repeat"
        backgroundImage="linear-gradient(45deg, #fff 25%, transparent 25%, transparent 75%, #fff 75%, #fff 100%),linear-gradient(45deg, #fff 25%, #e3e3e3 25%, #e3e3e3 75%, #fff 75%, #fff 100%)"
      />
    </Link>
  );
};
