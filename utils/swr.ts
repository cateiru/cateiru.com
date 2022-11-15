import {api} from './api';

export interface SWRError extends Error {
  message: string;
  status?: number;
}
/**
 * fetch used by swr
 *
 * @param {string} url - fetch url
 */
export async function fetcher(url: string) {
  const res = await fetch(api(url), {
    credentials: 'include',
    mode: 'cors',
  });

  if (!res.ok) {
    const error: SWRError = new Error(
      'An error occurred while fetching the data.'
    );
    error.message = (await res.json()).message;
    error.status = res.status;
    throw error;
  }
  return await res.json();
}
