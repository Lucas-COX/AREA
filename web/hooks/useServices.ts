import { useState, useEffect } from "react"
import axios from "axios";
import { Service } from "../config/types";

export default function useServices(accessToken: string) {
    const [services, setServices] = useState<Service[]>([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null | undefined>(undefined)
    useEffect(() => {
        axios.get(`${process.env.NEXT_PUBLIC_API_URL}/services`, {
            headers: { Authorization: 'Bearer ' + accessToken }
        }).then((res) => {
            setServices(res.data.services);
            setLoading(false);
        }).catch((err) => {
            setError(String(err));
            setLoading(false);
        })
    }, [accessToken])

    return {services, setServices, loading, error}
}
