import {Box, Button, SimpleGrid, useClipboard} from '@chakra-ui/react';
import React from 'react';
import {TbCheck} from 'react-icons/tb';

interface Props {
  color1: string;
  color2: string;
  color3?: string;
}

export const ColorCircle: React.FC<Props> = props => {
  return (
    <>
      <Box
        h="20px"
        background={
          props.color3
            ? `linear-gradient(90deg, ${props.color1}, ${props.color2}, ${props.color3})`
            : `linear-gradient(90deg, ${props.color1}, ${props.color2})`
        }
        borderRadius="25px"
      ></Box>
      <SimpleGrid
        columns={props.color3 ? 3 : 2}
        spacing={10}
        mt=".5rem"
        fontSize="1.2rem"
        mb="1rem"
      >
        <Box textAlign="left">
          <ColorCode>{props.color1}</ColorCode>
        </Box>
        <Box textAlign={props.color3 ? 'center' : 'right'}>
          <ColorCode>{props.color2}</ColorCode>
        </Box>
        {props.color3 && (
          <Box textAlign="right">
            <ColorCode>{props.color3}</ColorCode>
          </Box>
        )}
      </SimpleGrid>
    </>
  );
};

const ColorCode: React.FC<{children: string}> = props => {
  const {onCopy, hasCopied} = useClipboard(props.children);

  return (
    <Button variant="ghost" color={props.children} onClick={onCopy}>
      {hasCopied ? (
        <TbCheck color={props.children} size="25px" />
      ) : (
        props.children
      )}
    </Button>
  );
};
