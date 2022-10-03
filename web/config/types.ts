import { NextApiRequest, NextApiResponse } from 'next';

export type User = user;

export type UserBody = Omit<User, "password">;

export type ErrorBody = {
    message: string,
}

export type ProtectedApiHandler<T = any> = (req: NextApiRequest, res: NextApiResponse<T>, body: UserBody) => unknown | Promise<unknown>;