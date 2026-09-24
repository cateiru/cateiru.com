import path from "node:path";
import type { Page, TestInfo } from "@playwright/test";

// reg-suit (VRT) 向けのスナップショットヘルパー。
// 撮影したPNGを `regconfig.json` の `core.actualDir` に書き出す。

const ACTUAL_DIR = path.join(process.cwd(), "vrt-screenshots", "actual");

function toFileName(testInfo: TestInfo): string {
  // PC / SP で同じテストを実行するため、プロジェクト名を含めてファイル名の衝突を防ぐ
  return `${[testInfo.project.name, testInfo.title]
    .join(" - ")
    .replace(/[\\/:*?"<>|]/g, "_")}.png`;
}

export async function takeSnapshot(page: Page, testInfo: TestInfo) {
  await page.evaluate(() => document.fonts.ready);

  await page.screenshot({
    path: path.join(ACTUAL_DIR, toFileName(testInfo)),
    // 2画面目 (リンク一覧) は position: fixed で描画されるため、
    // フルページではなく現在のビューポートを撮影する
    fullPage: false,
    // 有限アニメーション (フェードイン) は完了状態まで早送りし、
    // 無限アニメーション (スクロールヒント) は初期状態に戻してから撮影する
    animations: "disabled",
  });
}
