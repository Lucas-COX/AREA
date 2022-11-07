import { GetServerSidePropsContext } from 'next';
import nookies from 'nookies';

export function getToken(context: GetServerSidePropsContext): string | null {
  const cookies = nookies.get(context);
  return cookies.area_token;
}
