import axios from 'axios'
import { IncomingMessage } from 'http'
import { Session } from '../config/types'
import { getToken } from './cookie'

export interface SessionContext {
    req: IncomingMessage
}

export async function getSession(context: SessionContext): Promise<Session> {
    return new Promise<Session>((resolve) => {
        const token = getToken(context.req)
        if (token == null) resolve({ user: undefined, authenticated: false })
        axios.get(`${process.env.NEXT_PUBLIC_API_URL}/me`, {
            headers: { 'Authorization': 'Bearer ' + token },
        }).then((res) => {
            resolve({ user: res.data.me, authenticated: true })
        }).catch(() => {
            resolve({ user: undefined, authenticated: false});
        })
    })
}