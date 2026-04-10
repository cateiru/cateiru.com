import type { Metadata } from "next";
import "./globals.css";
import "@fontsource/line-seed-jp";

export const metadata: Metadata = {
  metadataBase: new URL("https://cateiru.com"),
  title: "cateiru",
  description: "cateiru の公式サイト！",
  icons: {
    icon: "/favicon.svg",
  },
  openGraph: {
    title: "cateiru",
    description: "cateiru の公式サイト！",
    url: "https://cateiru.com",
    siteName: "cateiru",
    images: [{ url: "/ogp.png", width: 1200, height: 630 }],
    locale: "ja_JP",
    type: "website",
  },
  twitter: {
    card: "summary_large_image",
    title: "cateiru",
    description: "cateiru の公式サイト！",
    images: ["/ogp.png"],
  },
  robots: {
    index: true,
    follow: true,
  },
  alternates: {
    canonical: "https://cateiru.com",
  },
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
