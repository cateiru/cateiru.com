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
  } else if (d === '0001-01-01T00:00:00Z') {
    return '-';
  }

  const date = new Date(d);

  if (lang === 'ja') {
    return `${date.getFullYear()}年 ${
      date.getMonth() + 1
    }月 ${date.getDate()}日`;
  }

  return date.toLocaleDateString();
}

// Convert emoji to Unicode
export const toUnicode = (str: string) => {
  if (str.length < 4) return str.codePointAt(0)?.toString(16);
  return (
    str.codePointAt(0)?.toString(16) + '-' + str.codePointAt(2)?.toString(16)
  );
};
