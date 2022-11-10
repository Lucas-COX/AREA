import { PropsWithChildren } from 'react';
import { Session } from './types'

export interface withSession {
    session: Session;
}

export interface withChildren extends PropsWithChildren {}

export interface withClassName {
    className?: string;
}

export interface withTitle {
    title?: string;
}

export interface withShow {
    show?: boolean;
}

export interface withOnClose {
    onClose?: () => void;
}

export interface asTextField {
    value?: string;
    onChange?: (value: string) => void;
}