import { PropsWithChildren } from 'react';
import { Session } from './types'

export interface withSession {
    session: Session;
}

export interface withChildren extends PropsWithChildren {}

export interface withClassName {
    className?: string;
}