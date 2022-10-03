import Head from "next/head"
import { withChildren, withClassName } from "../config/withs";
import CenteredLayout from "./CenteredLayout";

export type LayoutType = "centered"

const LayoutMappings = {
    "centered": CenteredLayout,
}

export default function AppLayout({ type = "centered", className = "", children }: AppLayoutProps) {
    const Layout = LayoutMappings[type]
    return (
        <div>
            <Head>
                <title>Area</title>
                <meta name="description" content="To Do List application" />
                <link rel="icon" href={"/favicon.ico"} />
            </Head>
            <main className="w-screen h-screen">
                <Layout className={className}>
                    {children}
                </Layout>
            </main>
            <footer></footer>
        </div>
    )
}

export interface AppLayoutProps extends withChildren, withClassName {
    type?: LayoutType;
}