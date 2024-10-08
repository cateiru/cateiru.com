import React from "react";
import type { Public } from "../utils/types";
import Bio from "./Bio/Bio";
import ContactLink from "./Contact/ContactLink";
import Products from "./Products/Products";
import Profile from "./Profile/Profile";
import Social from "./Social/Social";

const Index = React.memo<{
  profile: Public;
}>((props) => {
  const bioRef = React.useRef<HTMLDivElement>(null!);
  const productsRef = React.useRef<HTMLDivElement>(null!);
  const socialRef = React.useRef<HTMLDivElement>(null!);
  const contactRef = React.useRef<HTMLDivElement>(null!);

  const scrollToBio = React.useCallback(() => {
    bioRef.current.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });
  }, [bioRef]);

  const scrollToProducts = React.useCallback(() => {
    productsRef.current.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });
  }, [productsRef]);

  const scrollToSocial = React.useCallback(() => {
    socialRef.current.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });
  }, [socialRef]);

  const scrollToContact = React.useCallback(() => {
    contactRef.current.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });
  }, [contactRef]);

  return (
    <>
      <Profile
        next={() => {
          scrollToBio();
        }}
        data={props.profile}
      />
      <Bio
        next={() => {
          scrollToProducts();
        }}
        r={bioRef}
        data={props.profile}
      />
      <Products
        next={() => {
          scrollToSocial();
        }}
        r={productsRef}
        data={props.profile}
      />
      <Social
        next={() => {
          scrollToContact();
        }}
        r={socialRef}
        data={props.profile}
      />
      <ContactLink r={contactRef} />
    </>
  );
});

Index.displayName = "Index";

export default Index;
