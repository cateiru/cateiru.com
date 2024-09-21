import { Box, Text, Heading, Center } from "@chakra-ui/react";
import { FaBirthdayCake } from "react-icons/fa";
import { getAge, parseDate } from "../../utils/parse";
import type { Public } from "../../utils/types";
import useLanguage from "../useLanguage";

const ProfileText: React.FC<{ data: Public }> = ({ data }) => {
  const { lang, convertLang } = useLanguage();
  return (
    <Box>
      <Center>
        <Heading textAlign="center">
          {convertLang({
            ja: `${data.family_name_ja} ${data.given_name_ja}`,
            en: `${data.given_name} ${data.family_name}`,
          })}
        </Heading>
      </Center>
      <Center mt="1rem">
        <FaBirthdayCake />
        <Text ml=".2rem">
          {parseDate(data.birth_date, lang)} ({getAge(data.birth_date)})
        </Text>
      </Center>
      <Center>
        <Text>{convertLang({ ja: data.location_ja, en: data.location })}</Text>
      </Center>
    </Box>
  );
};

export default ProfileText;
