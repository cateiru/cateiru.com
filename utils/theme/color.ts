/**
 * 補色を計算する
 *
 * @param {string} color - HEXのカラーコード
 * @returns {string} - 補色のHEXカラーコード
 */
export function complementary(color: string): string {
  color = color.replace("#", "");

  const r = Number.parseInt(color.substring(0, 2), 16);
  const g = Number.parseInt(color.substring(2, 4), 16);
  const b = Number.parseInt(color.substring(4, 6), 16);

  const buf = Math.max(r, g, b) + Math.min(r, g, b);

  const cr = buf - r;
  const cg = buf - g;
  const cb = buf - b;

  // 桁を揃えて結合
  return `#${cr.toString(16)}${cg.toString(16)}${cb.toString(16)}`;
}

/**
 * 文字色を計算する
 *
 * @param {string} color - HEXのカラーコード
 * @returns {string} - 文字色
 */
export function textColor(color: string): string {
  const t = Number.parseInt(color.replace("#", ""), 16);

  return +(
    (299 * ((t >> 16) & 255) + 587 * ((t >> 8) & 255) + 114 * (255 & t)) /
    1e3 /
    255
  ).toFixed(2) < 0.6
    ? "white"
    : "black";
}
