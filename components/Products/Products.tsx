import { SimpleGrid, Text } from "@chakra-ui/react";
import type React from "react";
import type { Public } from "../../utils/types";
import { Section } from "../Common/Section";
import useLanguage from "../useLanguage";
import { ProductCard } from "./ProductCard";

const Products: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
  data: Public;
}> = ({ next, r, data }) => {
  const { convertLang } = useLanguage();

  return (
    <Section next={next} r={r} heading={{ en: "Products", ja: "制作物" }}>
      {data.products.length === 0 ? (
        <Text textAlign="center" mt="1rem" color="yellow.500">
          {convertLang({
            ja: "表示できる制作物がありません。",
            en: "There is no product to display.",
          })}
        </Text>
      ) : (
        <SimpleGrid columns={{ base: 1, sm: 2, lg: 3 }} spacing={10} mt="1rem">
          {data.products.map((v) => {
            return <ProductCard key={v.id} prod={v} />;
          })}
        </SimpleGrid>
      )}
    </Section>
  );
};

export default Products;
