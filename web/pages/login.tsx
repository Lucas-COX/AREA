import CenteredLayout from "../components/CenteredLayout";
import { useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';

function LoginPage() {
    const router = useRouter()
    const [state, setState] = useState({
        name: "",
        password: "",
        error: ""
    })

    const handleNameChange = (e: any) => {
        setState({
            ...state,
            name: e.target.value
        })
    }
    const handlePasswordChange = (e: any) => {
        setState({
            ...state,
            password: e.target.value,
        })
    }
    const handleSubmit = (e: any) => {
        axios.post('/api/login', {
            username: state.name,
            password: state.password
        })
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            console.log(error);
        });
    }
    const canSubmit = () => ( state.name !== "" && state.password !== "" );
    return (
        <CenteredLayout title="Login page">
            <div className="flex flex-col space-y-20 text-dark" >
                <form className="flex flex-col space-y-10 items-end" onSubmit={handleSubmit}>
                    <div>
                        {state.name === "" ?
                        "No username specified" : state.name}
                    </div>
                    <div>
                        {state.password === "" ?
                        "No password specified" : state.password}
                    </div>
                    <label>
                        User ID : 
                        <input className="border border-primary flex space-x-2 rounded-md" type="text" name="name" onChange={handleNameChange} />
                    </label>
                    <label>
                        Password : 
                        <input className="border border-primary flex space-x-2 rounded-md" type="password" name="name" onChange={handlePasswordChange} />
                    </label>
                </form>
                <button disabled={!canSubmit()} className="border border-primary/50 bg-primary rounded-md p-3 hover:bg-secondary" onClick={handleSubmit}>Submit</button>
                <button type="button" onClick={() => router.push('/signin')}>Pas de compte ? Inscris toi</button>
            </div>
        </CenteredLayout>
    )
}

export default LoginPage