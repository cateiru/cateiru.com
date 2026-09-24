import { expect, test } from "@playwright/test";
import { takeSnapshot } from "./vrt-snapshot";

// ビジュアルリグレッションテスト (reg-suit)。
// PC / SP の各プロジェクト (playwright.vrt.config.ts) で実行される。

test("トップファーストビュー", async ({ page }, testInfo) => {
  await page.goto("/");
  await expect(page.getByRole("heading", { name: "cateiru" })).toBeVisible();
  await takeSnapshot(page, testInfo);
});

test("最下部までスクロール", async ({ page }, testInfo) => {
  await page.goto("/");
  await expect(page.getByRole("heading", { name: "cateiru" })).toBeVisible();

  await page.evaluate(() =>
    window.scrollTo({
      top: document.documentElement.scrollHeight,
      behavior: "instant",
    })
  );
  await page.waitForFunction(
    () =>
      window.scrollY + window.innerHeight >=
      document.documentElement.scrollHeight - 1
  );
  await expect(page.getByRole("link", { name: "GitHub" })).toBeInViewport();

  await takeSnapshot(page, testInfo);
});
