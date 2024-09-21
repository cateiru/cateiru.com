/**********************************************************
 * markdown h{number} tag.
 *
 * @author Yuto Watanabe <yuto.w51942@gmail.com>
 * @version 1.0.0
 *
 * Copyright (C) 2021 hello-slide
 **********************************************************/
import { Heading, Divider, Box } from "@chakra-ui/react";
import React from "react";
import type { HeadingComponent } from "react-markdown/lib/ast-to-react";

export const H1: HeadingComponent = ({ children }) => {
  return (
    <Box marginBottom="1rem" marginTop="2rem">
      <Heading fontSize="1.9rem" paddingLeft=".5rem" mb=".5rem">
        {children}
      </Heading>
      <Divider />
    </Box>
  );
};

export const H2: HeadingComponent = ({ children }) => {
  return (
    <Box marginBottom="1rem" marginTop="2rem">
      <Heading fontSize="1.7rem" paddingLeft=".5rem" mb=".5rem">
        {children}
      </Heading>
      <Divider />
    </Box>
  );
};

export const H3: HeadingComponent = ({ children }) => {
  return (
    <Box marginBottom="1rem" marginTop="2rem">
      <Heading fontSize="1.5rem" paddingLeft=".5rem" mb=".5rem">
        {children}
      </Heading>
      <Divider />
    </Box>
  );
};

export const H4: HeadingComponent = ({ children }) => {
  return (
    <Box marginBottom="1rem" marginTop="1.5rem">
      <Heading fontSize="1.4rem" paddingLeft=".5rem">
        {children}
      </Heading>
    </Box>
  );
};

export const H5: HeadingComponent = ({ children }) => {
  return (
    <Box marginBottom="1rem" marginTop="1.5rem">
      <Heading fontSize="1.3rem" paddingLeft=".5rem">
        {children}
      </Heading>
    </Box>
  );
};

export const H6: HeadingComponent = ({ children }) => {
  return (
    <Box marginBottom="1rem" marginTop="1.5rem">
      <Heading fontSize="1.2rem" paddingLeft=".5rem">
        {children}
      </Heading>
    </Box>
  );
};
