import axios from 'axios';
import { IncomingMessage } from 'http';
import { Session } from '../config/types';
import { getToken } from './cookie';
import { GetServerSidePropsContext } from 'next';

export interface SessionContext {
    req: IncomingMessage
}

export async function getSession(context: GetServerSidePropsContext): Promise<Session> {
  return new Promise<Session>((resolve) => {
    const token = getToken(context);
    if (token == null) resolve({ user: undefined, authenticated: false });
    axios.get(`${process.env.API_URL}/me`, {
      headers: { Authorization: `Bearer ${token}` },
    }).then((res) => {
      resolve({ user: res.data.me, authenticated: true, token });
    }).catch((e) => {
      console.log(e);
      resolve({ user: undefined, authenticated: false });
    });
  });
}
