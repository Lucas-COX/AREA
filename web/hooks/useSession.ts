import jwt from "jsonwebtoken"
import { useState, useEffect } from "react"
import { User } from "../config/types"

export interface Session {
    user: User | null,
    loggedIn: boolean
}

export default function useSession() {
    const [session, setSession] = useState<Session | undefined>(undefined)
    const [loading, setLoading] = useState<boolean>(true)
    const [error, setError] = useState<any | undefined>(undefined)

    useEffect(() => {
        const token = localStorage.getItem("epytodo_token")
        if (token == null) {
            setSession({ user: null, loggedIn: false })
            setLoading(false)
        } else {
            try {
                const user : any = jwt.decode(token)
                setSession({
                    user: { 
                        username: user.username,
                        first_name: user.first_name,
                        last_name: user.last_name,
                     },
                    loggedIn: true
                })
                setLoading(false)
            } catch (e: any) {
                setError(e)
            }
        }
    }, [])
    return [session, setSession, loading, error]
}