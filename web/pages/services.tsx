import { getSession } from '../lib/session';
import { useState } from 'react';
import type { GetServerSidePropsContext } from 'next';
import AppLayout from '../components/AppLayout';
import { Button } from '@mui/material';
import { withSession } from '../config/withs';
import axios from 'axios';
import { toast } from 'react-toastify';
import URLSafeBase64 from 'urlsafe-base64';
import useServices from '../hooks/useServices';
import { Color } from '../config/types';

export default function ServicesPage({ session }: ServicesPageProps) {
    const {services, setServices, loading, error} = useServices(session.token as string)
    const [registered, setRegistered] = useState(session?.user?.services || [])

    const handleServiceLogin = async (service: string) => {
        try {
          const url = URLSafeBase64.encode(Buffer.from(`${process.env.NEXT_PUBLIC_API_URL}/login/done`));

          const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/providers/${service}/auth?callback=${url}`, {
            headers: { Authorization: `Bearer ${session.token}` },
          });
          if (response.data.url !== null && !['discord', 'timer'].includes(service)) {
            window.open(response.data.url, '_blank')?.focus();
          }
          if (['discord', 'timer'].includes(service)) {
            setRegistered([...registered, service])
          }
          toast.success(`Successfully activated ${service} service.`)
        } catch (e) {
          toast.error(`Failed to redirect to ${service} authentication page.`);
        }
    };

    const handleServiceLogout = async (service: string) => {
        try {
            await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/providers/${service}/logout`, {
                headers: { Authorization: `Bearer ${session.token}` },
            });
            toast.success(`Successfully disconnected from ${service}.`)
            var index = registered.indexOf(service) || -1;
            if (index > -1) {
                setRegistered(service => service.filter((_, i) => i !== index))
            }
        } catch (e) {
            toast.error(`Failed to disconnect from ${service}`);
        }
    };

    const colors = ["error", "inherit", "primary", "secondary", "success", "info", "warning"]

    return (
        <AppLayout type="centered">
            <div className="space-x-5">
                {services.map((service, index) => {
                    const handleClickLogin = () => {
                        handleServiceLogin(service.name)
                        setRegistered
                    };
                    const handleClickLogout = () => handleServiceLogout(service.name);
                    return (
                        <Button
                            key={`service-${service.name}`}
                            color={colors[index % colors.length] as Color}
                            variant={registered.includes(service.name) == true ? "outlined" : "text"}
                            onClick={registered.includes(service.name) == true ? handleClickLogout : handleClickLogin}>
                            {service.name}
                        </Button>
                    )
                })}
            </div>
        </AppLayout>
    );
}

export interface ServicesPageProps extends withSession{}

export async function getServerSideProps(context: GetServerSidePropsContext) {
    const session = await getSession(context);
    if (session.authenticated == false) {
        return {
        redirect: {
            destination: '/login',
            permanent: false,
        },
        };
    }
    return {
        props: { session },
    };
}
