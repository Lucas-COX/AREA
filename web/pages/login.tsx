import CenteredLayout from "../components/CenteredLayout";
import { useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';
import useSession from "../hooks/useSession";
import { setToken } from "../lib/jwt";
import { Alert, Button, TextField } from "@mui/material";

function LoginPage() {
    const router = useRouter()
    const [session, loading, error] = useSession()
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
        }).then(function (response) {
            setToken(response.data.token);
            router.push('/')
        }).catch(function (error) {
            setState({ ...state, error });
        });
    }
    const handleRegister = (e: any) => {
        console.log(process.env)
        axios.post(`${process.env.NEXT_PUBLIC_API_URL}/register`, {
            username: state.username,
            password: state.password,
        }).then(function (response) {
            setToken(response.data.token)
            router.push('/')
        }).catch(function (error) {
            console.log(error);
            setState({ ...state, error });
        });
    }
    const canSubmit = () => ( state.username !== "" && state.password !== "" );

    if (loading)
        return <div>Spinner</div>

    if (session.user != null && !error)
        router.push("/")

    return (
        <CenteredLayout title="Authentication">
           {  loading ?
            <div>Spinner</div> :
            <div className="flex flex-col space-y-20 text-dark" >
                <form className="flex flex-col space-y-10 items-end" >
                    <TextField label="Username" variant="outlined" onChange={handleNameChange}/>
                    <TextField type="password" label="Password" variant="outlined" onChange={handlePasswordChange}/>
                </form>
                <Button disabled={!canSubmit()} variant="contained" onClick={handleLogin}>Log in</Button>
                <Button disabled={!canSubmit()} variant="outlined" onClick={handleRegister}>Register</Button>
                { state.error &&  <Alert severity="error">An error occured...</Alert> }
            </div> }
        </CenteredLayout>
    )
}

export default LoginPage