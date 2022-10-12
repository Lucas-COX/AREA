import React from 'react';
import clsx from 'clsx';
import { withChildren, withClassName } from '../config/withs';

export default function CenteredLayout({ children, className = '' }: CenteredLayoutProps) {
  return (
    <div className={clsx('h-full w-full flex items-center justify-center', className)}>
      {children}
    </div>
  );
}

export interface CenteredLayoutProps extends withChildren, withClassName {}
