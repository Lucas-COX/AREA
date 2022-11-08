import { getSession } from '../lib/session';
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

    const handleServiceLogin = async (service: string) => {
        try {
          const url = URLSafeBase64.encode(Buffer.from(`${process.env.NEXT_PUBLIC_API_URL}/login/done`));
    
          const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/providers/${service}/auth?callback=${url}`, {
            headers: { Authorization: `Bearer ${session.token}` },
          });
          if (response.data.url !== null) {
            window.open(response.data.url, '_blank')?.focus();
          }
        } catch (e) {
          toast.error(`Failed to redirect to ${service} authentication page.`);
        }
    };

    const handleServiceLogout = async (service: string) => {
        try {
          await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/providers/${service}/logout`, {
            headers: { Authorization: `Bearer ${session.token}` },
          });
          toast.success(`Successfully disconnected to ${service}!`)
        } catch (e) {
          toast.error(`Failed to disconnect to ${service}`);
        }
    };

    const colors = ["error", "inherit", "primary", "secondary", "success", "info", "warning"]

    console.log(session?.user?.services)
    return (
        <AppLayout type="centered">
            <div className="space-x-5">
                {services.map((service, index) => {
                    const handleClickLogin = () => handleServiceLogin(service.name);
                    const handleClickLogout = () => handleServiceLogout(service.name);
                    if (session?.user?.services.includes(service.name) == true) {
                        return (
                            <Button color={colors[index % colors.length] as Color} variant="outlined" onClick={handleClickLogout}>
                                {service.name}
                            </Button>
                        )
                    } else {
                        return (
                            <Button color={colors[index % colors.length] as Color} onClick={handleClickLogin}>
                                {service.name}
                            </Button>
                        )
                    }
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
  