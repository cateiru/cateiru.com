"use client";

import React from "react";
import styles from "./BgColorProvider.module.css";

export const BgColorProvider = () => {
  const [isScrollStarted, setIsScrollStarted] = React.useState(false);

  React.useEffect(() => {
    const handleScroll = () => {
      if (!isScrollStarted) {
        setIsScrollStarted(true);
      }
    };

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, [isScrollStarted]);

  const bgColor = React.useMemo(() => {
    if (isScrollStarted) {
      return "#fff";
    }
    return "#eef0f2";
  }, [isScrollStarted]);

  return (
    <div
      className={styles.background}
      style={{ "--bg-color": bgColor } as React.CSSProperties}
    />
  );
};
