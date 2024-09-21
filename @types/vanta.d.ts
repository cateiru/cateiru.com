declare module "vanta/dist/vanta.fog.min" {
  export default function set(c: Config): VantaEffect;
}

export interface Config {
  el: HTMLElement;
  mouseControls?: boolean;
  touchControls?: boolean;
  gyroControls?: boolean;
  minHeight?: number;
  minWidth?: number;
  highlightColor?: number;
  midtoneColor?: number;
  lowlightColor?: number;
  baseColor?: number;
  zoom?: number;
}

export interface VantaEffect {
  destroy(): void;
}
