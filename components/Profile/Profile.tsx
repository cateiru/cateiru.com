import { Center, Avatar, useColorMode } from "@chakra-ui/react";
import React from "react";
import FOG from "vanta/dist/vanta.fog.min";
import type { Public } from "../../utils/types";
import { Section } from "../Common/Section";
import ProfileText from "./ProfileText";

const Profile: React.FC<{ next: () => void; data: Public }> = ({
  next,
  data,
}) => {
  const { colorMode } = useColorMode();
  const myRef = React.useRef<HTMLDivElement>(null!);

  React.useEffect(() => {
    const effect = FOG({
      el: myRef.current,
      baseColor: colorMode === "dark" ? 0x0e121c : 0xffffff,
      highlightColor: 0x2bc4cf,
      midtoneColor: 0x572bcf,
      lowlightColor: 0xcf2ba1,
      mouseControls: true,
      touchControls: true,
      gyroControls: false,
      zoom: 0.5,
    });
    return () => {
      effect.destroy();
    };
  }, [colorMode]);

  return (
    <Section next={next} r={myRef}>
      <Center mb="1rem">
        <Avatar src={data.avatar_url} size="lg" />
      </Center>
      <ProfileText data={data} />
    </Section>
  );
};

export default Profile;
