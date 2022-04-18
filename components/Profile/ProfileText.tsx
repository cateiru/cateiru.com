import {Box, Text, Heading, Center} from '@chakra-ui/react';
import {FaBirthdayCake} from 'react-icons/fa';
import {getAge} from '../../utils/birthday';
import config from '../../utils/config/config';
import useLanguage from '../useLanguage';

const ProfileText = () => {
  const [lang, convertLang] = useLanguage();
  return (
    <Box textAlign="center">
      <Heading>{convertLang(config.name)}</Heading>
      <Center mt="1rem">
        <FaBirthdayCake />
        <Text ml=".2rem">
          {lang === 'en-US'
            ? `${config.birthday.getMonth()}/${config.birthday.getDate()}/${config.birthday.getFullYear()} (Age.${getAge(
                config.birthday
              )})`
            : `${config.birthday.getFullYear()}年 ${config.birthday.getMonth()}月 ${config.birthday.getDate()}日 (${getAge(
                config.birthday
              )}歳)`}
        </Text>
      </Center>
      <Center>
        {lang === 'en-US' ? (
          <>
            <Text>
              Born in {convertLang(config.birthplace.prefecture)},{' '}
              {convertLang(config.birthplace.country)} and lives in{' '}
              {convertLang(config.residence.prefecture)},{' '}
              {convertLang(config.birthplace.country)}
            </Text>
          </>
        ) : (
          <>
            <Text as="span" mr=".5rem">
              {convertLang(config.birthplace.prefecture)}
            </Text>
            →<Text ml=".5rem">{convertLang(config.residence.prefecture)}</Text>
          </>
        )}
      </Center>
    </Box>
  );
};

export default ProfileText;
