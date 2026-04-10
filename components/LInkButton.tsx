import type { FC } from "react";
import styles from "./LinkButton.module.css";

type LinkButtonProps = {
  href: string;
  children: React.ReactNode;
};

export const LinkButton: FC<LinkButtonProps> = (props) => {
  return (
    <a
      href={props.href}
      className={styles.link_button}
      target="_blank"
      rel="noopener noreferrer"
    >
      {props.children}
    </a>
  );
};
