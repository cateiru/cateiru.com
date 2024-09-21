import {
  Box,
  Center,
  Checkbox,
  FormControl,
  FormErrorMessage,
  FormLabel,
  useColorMode,
} from "@chakra-ui/react";
import React from "react";
import { ColorPicker, useColor } from "react-color-palette";
import { useFormContext, useWatch } from "react-hook-form";

import "react-color-palette/lib/css/styles.css";
import useLanguage from "../../useLanguage";

export interface ColorPickerFormType {
  isTransparent: boolean;
  r?: number;
  g?: number;
  b?: number;
}

export const ColorPickerForm = () => {
  const {
    setValue,
    register,
    control,
    formState: { errors },
  } = useFormContext<ColorPickerFormType>();
  const [color, setColor] = useColor("hex", "#fff");
  const isTransparent = useWatch({
    control,
    name: "isTransparent",
  });
  const { convertLang } = useLanguage();
  const { colorMode } = useColorMode();

  return (
    <>
      <FormControl mt="1rem" isInvalid={Boolean(errors.isTransparent)}>
        <Checkbox id="isTransparent" {...register("isTransparent")}>
          {convertLang({
            ja: "背景を塗りつぶす",
            en: "fill in the background",
          })}
        </Checkbox>
        <FormErrorMessage>
          {errors.isTransparent?.message}
        </FormErrorMessage>
      </FormControl>
      {isTransparent && (
        <Center>
          <Box
            mt=".5rem"
            boxShadow={
              colorMode === "light"
                ? "0px 1px 26px -3px #a0acc0"
                : "0px 1px 26px -3px #000"
            }
            w="300px"
            borderRadius="10px"
          >
            <ColorPicker
              width={300}
              height={228}
              color={color}
              onChange={(color) => {
                setColor(color);
              }}
              onChangeComplete={(color) => {
                setValue("r", color.rgb.r);
                setValue("g", color.rgb.g);
                setValue("b", color.rgb.b);
              }}
              hideHSV
              dark={colorMode === "dark"}
            />
          </Box>
        </Center>
      )}
    </>
  );
};
