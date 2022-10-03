import AppLayout from "../components/AppLayout";
import { useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';
import { Alert, Button, TextField } from "@mui/material";
import { getSession } from "../lib/session";
import { GetServerSidePropsContext } from "next";
import { withSession } from "../config/withs";

export default function LoginPage({}: LoginPageProps) {
    const router = useRouter()
    const [state, setState] = useState({
        username: "",
        password: "",
        error: undefined
    })

    const handleNameChange = (e: any) => {
        setState({ ...state, username: e.target.value })
    }
    const handlePasswordChange = (e: any) => {
        setState({ ...state, password: e.target.value })
    }

    const handleLogin = () => {
        axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
            username: state.username,
            password: state.password
        }, { withCredentials: true }).then(function (response) {
            router.push('/')
        }).catch(function (error) {
            setState({ ...state, error });
        });
    }
    const handleRegister = (e: any) => {
        axios.post(`${process.env.NEXT_PUBLIC_API_URL}/register`, {
            username: state.username,
            password: state.password,
        }, { withCredentials: true }).then(function (response) {
            router.push('/')
        }).catch(function (error) {
            setState({ ...state, error });
        });
    }
    const canSubmit = () => ( state.username !== "" && state.password !== "" );

    return (
        <AppLayout>
            <div className="flex flex-col space-y-20 text-dark" >
                <form className="flex flex-col space-y-10 items-end" >
                    <TextField label="Username" variant="outlined" onChange={handleNameChange}/>
                    <TextField type="password" label="Password" variant="outlined" onChange={handlePasswordChange}/>
                </form>
                <Button disabled={!canSubmit()} variant="contained" onClick={handleLogin}>Log in</Button>
                <Button disabled={!canSubmit()} variant="outlined" onClick={handleRegister}>Register</Button>
                { state.error &&  <Alert severity="error">An error occured...</Alert> }
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
