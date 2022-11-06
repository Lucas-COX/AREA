import { useState, useEffect } from "react"
import axios from "axios";
import { Action } from "../config/types";

export default function useActions(accessToken: string) {
    const [actions, setActions] = useState<Action[]>([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null | undefined>(undefined)
    useEffect(() => {
        if (!accessToken)
            return
        axios.get(`${process.env.NEXT_PUBLIC_API_URL}/actions`, {
            headers: { Authorization: 'Bearer ' + accessToken }
        }).then((res) => {
            setActions(res.data.actions);
            setLoading(false);
        }).catch((err) => {
            setError(String(err));
            setLoading(false);
        })
    }, [accessToken])
    return {actions, setActions, loading, error}
}
