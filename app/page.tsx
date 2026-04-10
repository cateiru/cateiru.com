import { Logo } from "@/components/Logo";
import styles from "./page.module.css";

export default function Home() {
  return (
    <main className={styles.main}>
      <div>
      <Logo />
      </div>
       <h1 className={styles.title}>cateiru.com</h1>
    </main>
  );
}
