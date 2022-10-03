import Head from 'next/head';
import React from 'react'

const CenteredLayout = ({
    title, icon = "/favicon.ico", children
}: CenteredLayoutProps) => {
    return (
        <div className='h-screen w-screen bg-light'>
            <Head>
                <title>{title}</title>
                <meta name="description" content="To Do List application" />
                <link rel="icon" href={icon} />
            </Head>

            <main className="flex h-full w-full items-center justify-center">
                {children}
            </main>

            <footer>
            </footer>
        </div>
    )
}

export interface CenteredLayoutProps {
    title: string;
    icon?: string;
    children: any;
}

export default CenteredLayout