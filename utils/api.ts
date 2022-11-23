export const API_DOMAIN =
  process.env.NEXT_PUBLIC_API_DOMAIN ?? 'https://api.cateiru.com';

export const api = (path: string): string => `${API_DOMAIN}${path}`;
