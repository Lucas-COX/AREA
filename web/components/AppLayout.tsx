import Head from "next/head"
import { withChildren, withClassName } from "../config/withs";
import CenteredLayout from "./CenteredLayout";
import { Button } from "@mui/material";
import { toast } from "react-toastify";
import axios from 'axios';
import { useRouter } from "next/router";

export type LayoutType = "centered"

const LayoutMappings = {
    "centered": CenteredLayout,
}

export default function AppLayout({ type = "centered", className = "", children, loggedIn = false }: AppLayoutProps) {
    const Layout = LayoutMappings[type]
    const router = useRouter()
    const handleLogout = async () => {
        try {
            await toast.promise(axios.get(`${process.env.NEXT_PUBLIC_API_URL}/logout`), {
                pending: "Logging you out...",
                error: "An error occurend while loggin you out.",
                success: "Successfully logged out.",
            })
            router.push('/login')
        } catch (e) {
            console.error(e);
        }
    }

    return (
        <div>
            <Head>
                <title>Area</title>
                <meta name="description" content="To Do List application" />
                <link rel="icon" href={"/favicon.ico"} />
            </Head>
            <main className="w-screen h-screen">
                <div className="w-full h-20 flex justify-end border-b border-secondary/30 bg-white shadow-md p-4 absolute">
                    {loggedIn && <Button variant={"outlined"} onClick={handleLogout}>Logout</Button>}
                </div>
                <Layout className={"pt-20 " + className}>
                    {children}
                </Layout>
            </main>
            <footer></footer>
        </div>
    )
}

export interface AppLayoutProps extends withChildren, withClassName {
    type?: LayoutType;
    loggedIn?: boolean;
}