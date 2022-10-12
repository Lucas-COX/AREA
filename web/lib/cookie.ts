import { IncomingMessage } from 'http';
import cookie from 'cookie';

export function getToken(req: IncomingMessage): string | null {
  if (!req.headers.cookie) return null;
  const cookies = cookie.parse(req.headers.cookie);
  return cookies.area_token;
}
