import {Center, Avatar, useColorMode} from '@chakra-ui/react';
import React from 'react';
import FOG from 'vanta/dist/vanta.fog.min';
import type {VantaEffect} from '../../@types/vanta';
import {Public} from '../../utils/types';
import {Section} from '../Common/Section';
import ProfileText from './ProfileText';

const Profile: React.FC<{next: () => void; data: Public}> = ({next, data}) => {
  const {colorMode} = useColorMode();
  const [vantaEffect, setVantaEffect] = React.useState<VantaEffect | null>(
    null
  );
  const myRef = React.useRef<HTMLDivElement>(null!);

  React.useEffect(() => {
    if (!vantaEffect) {
      setVantaEffect(
        FOG({
          el: myRef.current,
          baseColor: colorMode === 'dark' ? 0x1a202c : 0xffffff,
          highlightColor: 0xcf2ba1,
          midtoneColor: 0x2bc4cf,
          lowlightColor: 0x572bcf,
        })
      );
    }
    return () => {
      if (vantaEffect) vantaEffect.destroy();
    };
  }, []);

  React.useEffect(() => {
    if (!vantaEffect) return;

    vantaEffect.destroy();

    setVantaEffect(
      FOG({
        el: myRef.current,
        baseColor: colorMode === 'dark' ? 0x1a202c : 0xffffff,
        highlightColor: 0xcf2ba1,
        midtoneColor: 0x2bc4cf,
        lowlightColor: 0x572bcf,
      })
    );
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
