import { BgColorProvider } from "@/components/BgColorProvider";
import { Footer } from "@/components/Footer";
import { LinkButton } from "@/components/LInkButton";
import { Logo } from "@/components/Logo";
import { ScrollDown } from "@/components/ScrollDown";
import classNames from "classnames";
import React from "react";
import styles from "./page.module.css";

export default function Home() {
  return (
    <main className={styles.main}>
      <BgColorProvider />
      <div className={styles.page}>
        <div className={styles.flex_group}>
          <div className={styles.logo}>
            <Logo />
          </div>
          <h1 className={styles.title}>cateiru</h1>
        </div>
        <ScrollDown className={styles.scroll_down} />
      </div>
      <div className={classNames(styles.page, styles.link_page)}>
        <div className={styles.flex_group}>
          <LinkButton href="https://blog.cateiru.com">Blog</LinkButton>
          <LinkButton href="https://github.com/cateiru">GitHub</LinkButton>
          <LinkButton href="https://x.com/cateiru">X</LinkButton>
        </div>
        <Footer />
      </div>
    </main>
  );
}
