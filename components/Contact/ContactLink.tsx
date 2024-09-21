import { Center, Button } from "@chakra-ui/react";
import { useRouter } from "next/router";
import React from "react";
import useLanguage from "../useLanguage";

const ContactLink: React.FC<{
  r: React.MutableRefObject<HTMLDivElement>;
}> = ({ r }) => {
  const { convertLang } = useLanguage();
  const router = useRouter();

  return (
    <Center height="100vh" ref={r}>
      <Button
        variant="outline"
        onClick={() => {
          router.push("/contact");
        }}
      >
        {convertLang({ ja: "お問い合わせ", en: "Contact Us" })}
      </Button>
      <Button
        variant="outline"
        onClick={() => {
          router.push("/brand_resources");
        }}
        ml=".5rem"
      >
        {convertLang({ ja: "ブランドリソース", en: "Brand Resources" })}
      </Button>
    </Center>
  );
};

export default ContactLink;
