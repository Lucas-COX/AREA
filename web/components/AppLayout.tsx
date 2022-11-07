import Head from 'next/head';
import { Button } from '@mui/material';
import { toast } from 'react-toastify';
import nookies from 'nookies';
import { useRouter } from 'next/router';
import clsx from 'clsx';
import CenteredLayout from './CenteredLayout';
import SideMenu from './SideMenu';
import { withChildren, withClassName } from '../config/withs';

export type LayoutType = 'centered'

const LayoutMappings = {
  centered: CenteredLayout,
};

export default function AppLayout({
  type = 'centered', className = '', children, loggedIn = false,
}: AppLayoutProps) {
  const Layout = LayoutMappings[type];
  const router = useRouter();
  const handleLogout = async () => {
    try {
      toast.success("You're now logged out.");
      nookies.destroy(null, 'area_token');
      router.push('/login');

    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div>
      <Head>
        <title>Area</title>
        <meta name="description" content="To Do List application" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className="w-screen h-screen">
        <div className="w-full h-20 flex justify-between items-center border-b top-0 left-0 border-secondary/30 bg-white shadow-md p-4 absolute">
          <SideMenu />
          {loggedIn && <Button variant="outlined" onClick={handleLogout}>Logout</Button>}
        </div>
        <Layout className={clsx(className, 'pt-20')}>
          {children}
        </Layout>
      </main>
      <footer />
    </div>
  );
}

export interface AppLayoutProps extends withChildren, withClassName {
    type?: LayoutType;
    loggedIn?: boolean;
}
