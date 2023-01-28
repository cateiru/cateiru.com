import {Box, Button, Center, Heading, Image, Link} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import useLanguage from '../useLanguage';

export const Icon = () => {
  const {convertLang} = useLanguage();

  return (
    <Box>
      <Heading textAlign="center" mb="1rem" id="icons">
        {convertLang({ja: 'アイコン', en: 'Icons'})}
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
      <Center mt="1.5rem">
        <NextLink href="/brand_resources/icon_generator" passHref>
          <Button as="p">
            {convertLang({ja: 'アイコンジェネレーター', en: 'Icon Generator'})}
          </Button>
        </NextLink>
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
