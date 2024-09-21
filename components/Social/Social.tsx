import { SimpleGrid, Text } from "@chakra-ui/react";
import type React from "react";
import type { Public } from "../../utils/types";
import { Section } from "../Common/Section";
import useLanguage from "../useLanguage";
import { CategoryCard } from "./CategoryCard";

const Social: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
  data: Public;
}> = ({ next, r, data }) => {
  const { convertLang } = useLanguage();

  return (
    <Section next={next} r={r} heading={{ en: "Links", ja: "リンク" }}>
      {data.links.length === 0 ? (
        <Text textAlign="center" mt="1rem" color="yellow.500">
          {convertLang({
            ja: "表示できるリンクがありません。",
            en: "There is no link to display.",
          })}
        </Text>
      ) : (
        <SimpleGrid columns={{ base: 1, sm: 2, lg: 3 }} spacing={10} mt="1rem">
          {data.links.map((v) => (
            <CategoryCard links={v} key={v.category_id} />
          ))}
        </SimpleGrid>
      )}
    </Section>
  );
};

export default Social;
