import { useState, useEffect } from "react"
import { User } from "../config/types"
import { getToken } from "../lib/jwt"

export interface Session {
    user: User | null,
}

export default function useSession() {
    const [session, setSession] = useState<Session | undefined>({ user: null })
    const [loading, setLoading] = useState<boolean>(true)
    const [error, setError] = useState<any | undefined>(undefined)

    useEffect(() => {
        const token = getToken();
        if (token == null) {
            setLoading(false)
        } else {
            try {
                // Todo: fetch the /me route giving the token and store the returned user
                // const user : any = jwt.decode(token)
                // setSession({
                //     user: {
                //         username: user.username,
                //         first_name: user.first_name,
                //         last_name: user.last_name,
                //      },
                //     loggedIn: true
                // })
                setLoading(false)
            } catch (e: any) {
                setError(e)
            }
        }
    }, [])
    return [session, loading, error]
}