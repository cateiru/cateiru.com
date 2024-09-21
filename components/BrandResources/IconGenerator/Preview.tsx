import { Box, Center, Image, useColorMode } from "@chakra-ui/react";
import React from "react";
import { TbDownload } from "react-icons/tb";

export const PreviewImage: React.FC<{ src: string; download: () => void }> = ({
  src,
  download,
}) => {
  const [hover, setHover] = React.useState(false);
  const { colorMode } = useColorMode();

  return (
    <Center mt="1rem">
      <Box
        width="fit-content"
        height="fit-content"
        boxShadow={
          colorMode === "light"
            ? "0px 1px 26px -3px #a0acc0"
            : "0px 1px 26px -3px #000"
        }
        onClick={download}
        cursor="pointer"
        position="relative"
        onMouseEnter={() => setHover(true)}
        onMouseLeave={() => setHover(false)}
      >
        <Box
          position="absolute"
          top="50%"
          left="50%"
          transform="translate(-50%, -50%)"
          borderRadius="50%"
          bgColor="white"
          p=".5rem"
          visibility={hover ? "visible" : "hidden"}
          opacity={hover ? 1 : 0}
          transition=".2s cubic-bezier(0.45, 0, 0.55, 1)"
        >
          <TbDownload size="100px" color="#000" />
        </Box>
        <Image
          src={src}
          alt="generated image"
          backgroundPosition="0px 0px, 7px 7px"
          backgroundSize="14px 14px"
          backgroundRepeat="repeat"
          backgroundImage="linear-gradient(45deg, #fff 25%, transparent 25%, transparent 75%, #fff 75%, #fff 100%),linear-gradient(45deg, #fff 25%, #e3e3e3 25%, #e3e3e3 75%, #fff 75%, #fff 100%)"
        />
      </Box>
    </Center>
  );
};
