import type { Metadata } from "next";
import "./globals.css";
import "@fontsource/line-seed-jp";

export const metadata: Metadata = {
  title: "cateiru",
  description: "cateiru の公式サイト！",
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  );
}
