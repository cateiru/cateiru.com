import chalk from "chalk";
import type { Task } from "gantt-task-react";
import stc from "string-to-color";
import type { Contact, Public, PublicBio } from "./types";

/**
 * 日時をフォーマットする
 *
 * @param {string} d - date
 * @param {string} lang - languages
 * @returns {string} - formatted date
 */
export function parseDate(d: string, lang: string): string {
  if (d.length === 9) {
    return lang === "ja" ? "不明" : "Unknown";
  }
  if (d === "0001-01-01T00:00:00Z") {
    return "-";
  }

  const date = new Date(d);

  if (lang === "ja") {
    return `${date.getFullYear()}年 ${
      date.getMonth() + 1
    }月 ${date.getDate()}日`;
  }

  return date.toLocaleDateString();
}

/**
 * 詳細データにパースする
 *
 * @param {string} d - date
 * @returns {string} - formatted date
 */
export function parseDetailDate(d: string): string {
  if (d === "0001-01-01T00:00:00Z") {
    return "-";
  }

  const date = new Date(d);
  return date.toLocaleString();
}

/**
 *
 *
 * @param {string} d - date
 * @param {string} lang - languages
 * @returns {string} - formatted date
 */
export function parseAgo(d: string, lang: string): string {
  if (d === "0001-01-01T00:00:00Z") {
    return "-";
  }

  const date = new Date(d);
  const now = new Date();
  const diffSec = Math.floor((now.getTime() - date.getTime()) / 1000);

  if (diffSec < 3600) {
    return lang === "ja"
      ? `${Math.floor(diffSec / 60)}分前`
      : `${Math.floor(diffSec / 60)} min ago`;
  }
  if (diffSec < 86400) {
    return lang === "ja"
      ? `${Math.floor(diffSec / 3600)}時間前`
      : `${Math.floor(diffSec / 3600)} hors ago`;
  }
  if (diffSec < 86400 * 7) {
    return lang === "ja"
      ? `${Math.floor(diffSec / 86400)}日前`
      : `${Math.floor(diffSec / 86400)} day ago`;
  }
  if (diffSec < 86400 * 30) {
    return lang === "ja"
      ? `${Math.floor(diffSec / (86400 * 7))}週間前`
      : `${Math.floor(diffSec / (86400 * 7))} week ago`;
  }
  return lang === "ja"
    ? `${Math.floor(diffSec / (86400 * 30))}ヶ月以上前`
    : `over ${Math.floor(diffSec / (86400 * 30))} months ago`;
}

/**
 * 年月 表示用のDate
 *
 * @param {string} d - date
 * @param {string} lang - languages
 * @returns {string} - formatted date
 */
export function parseShotDate(d: string, lang: string): string {
  if (d === "0001-01-01T00:00:00Z") {
    return "";
  }

  const date = new Date(d);

  if (lang === "ja") {
    return `${date.getFullYear()}年 ${date.getMonth() + 1}月`;
  }

  return `${date.getMonth() + 1}/${date.getFullYear()}`;
}

/**
 * Convert emoji to Unicode
 *
 * 1f468 dc68 200d 1f469 dc69 200d 1f466 dc66
 * ↓
 * 1f468-200d-1f469-200d-1f466
 *
 * 1f1ef-ddef-1f1f5-ddf5
 * ↓
 * 1f1ef-1f1f5
 *
 * @param {string} str - emoji
 * @returns {string} - unicode
 */
export const toUnicode = (str: string): string => {
  const utf8: string[] = [];
  const len = str.length;

  for (let i = 0; len > i; i++) {
    const u = str.codePointAt(i)?.toString(16);
    if (u) {
      if (!u.match(/d[a-z][0-9a-z]+/)) {
        utf8.push(u);
      }
    }
  }

  return utf8.join("-");
};

export const getAge = (d: string): number => {
  const birthDate = new Date(d);
  const today = new Date();
  let age = today.getFullYear() - birthDate.getFullYear();
  const m = today.getMonth() - birthDate.getMonth();
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--;
  }

  return age;
};

/**
 * show Public data in console
 *
 * @param {Public} d - public data
 */
export function consolePublic(d: Public) {
  console.log(`ID: ${chalk.bgHex("#086f83").white(d.user_id)}`);
}

/**
 * parse tasks
 *
 * @param {PublicBio} d - bio data
 * @param {string} lang - language
 * @returns {Task[]} - tasks
 */
export function getTasks(d: PublicBio, lang: string): Task[] {
  const tasks: Task[] = [];

  for (let i = 0; d.length > i; i++) {
    const start = new Date(d[i].join);
    let end;
    if (d[i].leave === "0001-01-01T00:00:00Z") {
      end = new Date();
    } else {
      end = new Date(d[i].leave);
    }
    const barColor = stc(`${d[i].name}${d[i].position}`);

    tasks.push({
      id: `${i}`,
      type: "task",
      name:
        lang === "ja"
          ? `${d[i].name_ja} - ${d[i].position_ja}`
          : `${d[i].name} - ${d[i].position}`,
      start: start,
      end: end,
      progress: 0,
      styles: {
        backgroundColor: barColor,
        backgroundSelectedColor: barColor,
      },
      isDisabled: true,
    });
  }

  return tasks;
}

/**
 *
 * @param {string} str - target
 * @returns {string} - sliced str
 */
export function sliceStr(str: string): string {
  const returnValueMatch = str.match(/\n/g);
  if (returnValueMatch && returnValueMatch.length > 3) {
    const sliceLen = returnValueMatch.reduce((_, v, a) => v.length + a, 0);
    return str.slice(0, sliceLen);
  }

  if (100 < str.length) {
    return `${str.slice(0, 100)}...`;
  }
  return str;
}

/**
 *
 * @param {Contact} form - form data
 * @returns {string} - format copy value
 */
export function copyElement(form: Contact): string {
  const formTexts: string[] = [
    "## Inquiry",
    `- name: ${form.name}`,
    `- email: ${form.mail}`,
    `- date: ${new Date(form.created).toLocaleString()}`,
    "",
    `### ${form.title}`,
    "",
    ...form.detail.split("\n"),
  ];

  return formTexts.map((v) => `> ${v}`).join("\n");
}
