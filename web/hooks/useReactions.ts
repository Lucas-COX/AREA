import { useState, useEffect } from "react"
import axios from "axios";
import { Reaction } from "../config/types";

export default function useReactions(accessToken: string) {
    const [reactions, setReactions] = useState<Reaction[]>([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null | undefined>(undefined)
    useEffect(() => {
        axios.get(`${process.env.NEXT_PUBLIC_API_URL}/reactions`, {
            headers: { Authorization: 'Bearer ' + accessToken }
        }).then((res) => {
            setReactions(res.data.reactions);
            setLoading(false);
        }).catch((err) => {
            setError(String(err));
            setLoading(false);
        })
    }, [accessToken])

    return {reactions, setReactions, loading, error}
}
