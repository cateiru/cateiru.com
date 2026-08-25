import classNames from "classnames";
import styles from "./ScrollDown.module.css";

type ScrollDownProps = {
  className?: string;
};

export const ScrollDown = ({ className }: ScrollDownProps) => {
  return (
    <div
      role="img"
      aria-label="下にスクロール"
      className={classNames(styles.scroll_down, className)}
    />
  );
};
