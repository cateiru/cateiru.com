import Link from "next/link";
import { Footer } from "@/components/Footer";
import styles from "./not-found.module.css";

export default function NotFound() {
  return (
    <main className={styles.main}>
      <div className={styles.page}>
        <p className={styles.code}>404</p>
        <p className={styles.message}>お探しのページは見つかりませんでした</p>
        <Link href="/" className={styles.link}>
          トップに戻る
        </Link>
      </div>
      <Footer />
    </main>
  );
}
