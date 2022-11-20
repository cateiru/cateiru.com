import {
  Center,
  Box,
  Heading,
  Text,
  UnorderedList,
  ListItem,
} from '@chakra-ui/react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import {Area} from '../../utils/config/area';
import {BiographyDetail, WorksDetail} from '../../utils/config/bio';
import config from '../../utils/config/config';
import useLanguage from '../useLanguage';

const Bio: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
}> = ({next, r}) => {
  const {lang, convertLang} = useLanguage();

  const area = (area: Area) => {
    if (lang === 'en') {
      return `${convertLang(area.prefecture)}, ${convertLang(area.country)}`;
    }

    return convertLang(area.prefecture);
  };

  const biographyElement = (bio: BiographyDetail) => {
    return (
      <ListItem my=".5rem">
        <Text as="span" display="inline-block" minW="10ch">
          {bio.admission.getFullYear()}&nbsp;-&nbsp;
          {bio.graduation ? bio.graduation.getFullYear() : ''}
        </Text>
        <Text as="span" ml="1rem" fontWeight="bold">
          {convertLang(bio.name)}
        </Text>
        <Text as="span" ml=".5rem">
          ({area(bio.area)})
        </Text>
      </ListItem>
    );
  };

  const workElement = (work: WorksDetail) => {
    return (
      <ListItem my=".5rem" key={work.name.en}>
        <Text as="span" display="inline-block" minW="10ch">
          {work.admission.getFullYear()}&nbsp;-&nbsp;
          {work.graduation ? work.graduation.getFullYear() : ''}
        </Text>
        <Text as="span" ml="1rem" fontWeight="bold">
          {convertLang(work.name)}
        </Text>
        <Text as="span" ml=".5rem">
          ({area(work.area)})
        </Text>
        <Text as="span" ml=".5rem">
          - {convertLang(work.occupation)}
        </Text>
      </ListItem>
    );
  };

  return (
    <Center height="100vh" ref={r}>
      <Box width={{base: '90%', md: '700px'}}>
        <Heading textAlign="center">
          {convertLang({en: 'Brief personal record', ja: '略歴'})}
        </Heading>
        <Box whiteSpace="nowrap" overflowX="auto">
          <UnorderedList mt="1.5rem" paddingLeft=".5rem">
            {config.bio.primarySchool &&
              biographyElement(config.bio.primarySchool)}
            {config.bio.juniorHighSchool &&
              biographyElement(config.bio.juniorHighSchool)}
            {config.bio.highSchool && biographyElement(config.bio.highSchool)}
            {config.bio.university && biographyElement(config.bio.university)}
            {config.bio.works.map(v => workElement(v))}
          </UnorderedList>
        </Box>
        <Center mt="1.5rem" cursor="pointer" onClick={next}>
          <IoCaretDown size="25px" />
        </Center>
      </Box>
    </Center>
  );
};

export default Bio;
