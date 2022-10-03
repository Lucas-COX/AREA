import CenteredLayout from "../components/CenteredLayout";
import { useState } from 'react';
import { Button } from "@mui/material";
import axios from "axios";
import { TextField } from "@mui/material";

function SigninPage() {
    const [state, setState] = useState({
        firstname: "",
        lastname: "",
        username: "",
        password: "",
        confirm: "",
        error: ""
    })


    const handleFirstNameChange = (e: any) => {
        setState({
            ...state,
            firstname: e.target.value
        })
    }
    const handleLastNameChange = (e: any) => {
        setState({
            ...state,
            lastname: e.target.value
        })
    }
    const handleNameChange = (e: any) => {
        setState({
            ...state,
            username: e.target.value
        })
    }
    const handlePasswordChange = (e: any) => {
        setState({
            ...state,
            password: e.target.value,
        })
    }
    const handleConfirmChange = (e: any) => {
        setState({
            ...state,
            confirm: e.target.value
        })
    }


    const handleSubmit = (e: any) => {
        axios.post('/api/register', {
            username: state.username,
            password: state.password,
            first_name: state.firstname,
            last_name: state.lastname,
        })
        .then(function (response) {
            localStorage.setItem('epytodo_token', response.data.token)
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    const canSubmit = () => ( state.username !== "" && state.password !== "" && state.confirm !== "" && state.password === state.confirm );
    
    return (
        <CenteredLayout title="Sign in">
            <div className="flex flex-col space-y-5 justify-center">
                <form className="flex flex-col space-y-5">
                    <TextField label="first name" variant="outlined" onChange={handleFirstNameChange} />
                    <TextField label="last name" variant="outlined" onChange={handleLastNameChange} />
                    <TextField label="username" variant="outlined" onChange={handleNameChange} />
                    <TextField type="password" label="password" variant="outlined" onChange={handlePasswordChange} color={"secondary"} />
                    <TextField type="password" label="confirm" variant="outlined" onChange={handleConfirmChange} />
                </form>
                <Button disabled={!canSubmit()} variant="contained" className="mt-4" onClick={handleSubmit} >Submit</Button>
            </div>
        </CenteredLayout>
    )
}

export default SigninPage
