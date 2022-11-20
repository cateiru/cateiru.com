import chalk from 'chalk';
import type {Task} from 'gantt-task-react';
import stc from 'string-to-color';
import {Public, PublicBio} from './types';

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
  console.log('ID: ' + chalk.bgHex('#086f83').white(d.user_id));
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
    if (d[i].leave === '0001-01-01T00:00:00Z') {
      end = new Date();
    } else {
      end = new Date(d[i].leave);
    }
    const barColor = stc(`${d[i].name}${d[i].position}`);

    tasks.push({
      id: `${i}`,
      type: 'task',
      name:
        lang === 'ja'
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
