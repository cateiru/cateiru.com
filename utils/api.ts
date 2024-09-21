export const API_DOMAIN =
  process.env.NEXT_PUBLIC_API_DOMAIN ?? "https://api.cateiru.com";
export const BACKEND_API_DOMAIN =
  process.env.NEXT_PUBLIC_BACKEND_API_DOMAIN ?? "https://api.cateiru.com";

export const api = (path: string): string => `${API_DOMAIN}${path}`;

export const backendApi = (path: string): string =>
  `${BACKEND_API_DOMAIN}${path}`;
