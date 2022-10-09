import AppLayout from "../components/AppLayout";
import { useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';
import { Alert, Button, TextField } from "@mui/material";
import { getSession } from "../lib/session";
import { GetServerSidePropsContext } from "next";
import { toast } from "react-toastify";

export default function LoginPage({}: LoginPageProps) {
    const router = useRouter()
    const [state, setState] = useState({
        username: "",
        password: "",
    })

    const handleNameChange = (e: any) => {
        setState({ ...state, username: e.target.value })
    }
    const handlePasswordChange = (e: any) => {
        setState({ ...state, password: e.target.value })
    }

    const handleLogin = async () => {
        await toast.promise(axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
            username: state.username,
            password: state.password,
        }, { withCredentials: true }), {
            pending: "Loading...",
            error: "An error has occured",
            success: `Hello ${state.username} !`,
        })
        router.push('/')
    }
    const handleRegister = async (e: any) => {
        await toast.promise(axios.post(`${process.env.NEXT_PUBLIC_API_URL}/register`, {
            username: state.username,
            password: state.password,
        }, { withCredentials: true }), {
            pending: "Loading...",
            error: "An error has occurred",
            success: `Hello ${state.username} !`,
        })
        router.push('/')
    }
    const canSubmit = () => ( state.username !== "" && state.password !== "" );

    return (
        <AppLayout className="bg-blue-50/50">
            <div className="flex flex-col space-y-10 text-dark border border-gray-200 p-10 rounded-md shadow-md bg-white">
                <form className="flex flex-col space-y-10 items-center">
                    <h1 className="text-xl">Login to <span className="text-blue-600 font-bold">Area</span></h1>
                    <TextField label="Username" variant="outlined" onChange={handleNameChange}/>
                    <TextField type="password" label="Password" variant="outlined" onChange={handlePasswordChange}/>
                </form>
                <div className="flex justify-between space-x-2 h-12">
                    <Button disabled={!canSubmit()} variant="outlined" onClick={handleLogin} color={"primary"} className="w-1/2">Log in</Button>
                    <Button disabled={!canSubmit()} variant="outlined" onClick={handleRegister} color={"secondary"} className="w-1/2">Register</Button>
                </div>
            </div>
        </AppLayout>
    )
}

export interface LoginPageProps {}

export async function getServerSideProps(context: GetServerSidePropsContext) {
  const session = await getSession(context)
  if (session.authenticated == true) {
    return {
      redirect: {
        destination: '/',
        permanent: false,
      },
    }
  }
  return {
    props: {}
  }
}
