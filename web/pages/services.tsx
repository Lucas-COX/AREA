import { getSession } from '../lib/session';
import type { GetServerSidePropsContext } from 'next';
import AppLayout from '../components/AppLayout';
import { Button } from '@mui/material';
import { withSession } from '../config/withs';
import axios from 'axios';
import { toast } from 'react-toastify';
import URLSafeBase64 from 'urlsafe-base64';
import useServices from '../hooks/useServices';

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
          toast.error('Failed to redirect to Google authentication page.');
        }
    };
    return (
        <AppLayout type="centered">
            <div>
                <Button>
                    Service Name
                </Button>
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
  