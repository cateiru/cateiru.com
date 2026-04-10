import { Logo } from "@/components/Logo";
import { LinkButton } from "@/components/LInkButton";
import styles from "./page.module.css";
import classNames from "classnames";
import { Footer } from "@/components/Footer";

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.page}>
        <div className={styles.flex_group}>
          <div className={styles.logo}>
            <Logo />
          </div>
          <h1 className={styles.title}>cateiru</h1>
        </div>
      </div>
      <div className={classNames(styles.page, styles.link_page)}>
        <div className={styles.flex_group}>
          <LinkButton href="https://blog.cateiru.com">Blog</LinkButton>
          <LinkButton href="https://github.com/cateiru">GitHub</LinkButton>
          <LinkButton href="https://x.com/cateiru">X</LinkButton>
        </div>
      </div>
      <Footer />
    </main>
  );
}
