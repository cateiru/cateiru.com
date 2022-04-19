import React from 'react';
import Bio from './Bio/Bio';
import ContactLink from './Contact/ContactLink';
import Products from './Products/Products';
import Profile from './Profile/Profile';
import Social from './Social/Social';

const Index = () => {
  const bioRef = React.useRef<HTMLDivElement>(null!);
  const productsRef = React.useRef<HTMLDivElement>(null!);
  const socialRef = React.useRef<HTMLDivElement>(null!);
  const contactRef = React.useRef<HTMLDivElement>(null!);

  const scrollToBio = React.useCallback(() => {
    bioRef.current.scrollIntoView({
      behavior: 'smooth',
      block: 'end',
    });
  }, [bioRef]);

  const scrollToProducts = React.useCallback(() => {
    productsRef.current.scrollIntoView({
      behavior: 'smooth',
      block: 'end',
    });
  }, [productsRef]);

  const scrollToSocial = React.useCallback(() => {
    socialRef.current.scrollIntoView({
      behavior: 'smooth',
      block: 'end',
    });
  }, [socialRef]);

  const scrollToContact = React.useCallback(() => {
    contactRef.current.scrollIntoView({
      behavior: 'smooth',
      block: 'end',
    });
  }, [contactRef]);

  return (
    <>
      <Profile
        next={() => {
          scrollToBio();
        }}
      />
      <Bio
        next={() => {
          scrollToProducts();
        }}
        r={bioRef}
      />
      <Products
        next={() => {
          scrollToSocial();
        }}
        r={productsRef}
      />
      <Social
        next={() => {
          scrollToContact();
        }}
        r={socialRef}
      />
      <ContactLink r={contactRef} />
    </>
  );
};

export default Index;
