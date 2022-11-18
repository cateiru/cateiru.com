import {Box, useColorMode} from '@chakra-ui/react';
import React from 'react';

export const CardFrame: React.FC<{children: React.ReactNode}> = ({
  children,
}) => {
  const {colorMode} = useColorMode();

  return (
    <Box
      w="100%"
      minH="350px"
      boxShadow={
        colorMode === 'light'
          ? '0px 1px 26px -3px #a0acc0'
          : '0px 1px 26px -3px #000'
      }
      borderRadius="56px"
      p="1rem"
      position="relative"
    >
      {children}
    </Box>
  );
};
