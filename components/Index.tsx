import {Center, Heading} from '@chakra-ui/react';
import React from 'react';
import {Public} from '../utils/types';
import Bio from './Bio/Bio';
import ContactLink from './Contact/ContactLink';
import Products from './Products/Products';
import Profile from './Profile/Profile';
import Social from './Social/Social';

const Index = React.memo<{
  profile?: Public;
  error?: string;
}>(props => {
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

  if (typeof props.profile === 'undefined') {
    return (
      <Center color="red.500" h="100vh">
        <Heading>{props.error ?? 'ERROR'}</Heading>
      </Center>
    );
  }

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
});

Index.displayName = 'Index';

export default Index;
