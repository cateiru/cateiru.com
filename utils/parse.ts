/**
 * 日時をフォーマットする
 *
 * @param {string} d - date
 * @param {string} lang - languages
 * @returns {string} - formatted date
 */
export function parseDate(d: string, lang: string): string {
  if (d.length === 9) {
    return lang === 'ja' ? '不明' : 'Unknown';
  }

  const date = new Date(d);

  if (lang === 'ja') {
    return `${date.getFullYear()}年 ${
      date.getMonth() + 1
    }月 ${date.getDate()}日`;
  }

  return date.toLocaleDateString();
}
