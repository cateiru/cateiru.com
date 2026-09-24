import { defineConfig, devices } from "@playwright/test";

// 開発サーバー (3000番) と衝突しないよう、VRT 専用のポートで起動する
const PORT = 3100;
const BASE_URL = `http://localhost:${PORT}`;

// ビジュアルリグレッションテスト (reg-suit) 用の Playwright 設定
export default defineConfig({
  testDir: "./e2e",
  testMatch: /vrt\.spec\.ts$/,
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  reporter: "html",
  use: {
    baseURL: BASE_URL,
    trace: "on-first-retry",
  },
  // PC / SP それぞれで撮影する。CI では chromium のみインストールするため、
  // SP も Chromium ベースのデバイス (Pixel 7) を使う
  projects: [
    {
      name: "pc",
      use: { ...devices["Desktop Chrome"] },
    },
    {
      name: "sp",
      use: { ...devices["Pixel 7"] },
    },
  ],
  // 開発用オーバーレイなどが写り込まないよう、本番ビルドに対して撮影する
  webServer: {
    command: `pnpm build && pnpm start --port ${PORT}`,
    url: BASE_URL,
    reuseExistingServer: !process.env.CI,
    timeout: 180_000,
  },
});
