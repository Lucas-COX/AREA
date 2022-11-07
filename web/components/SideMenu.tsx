import clsx from 'clsx';
import { useState } from 'react';
import { useRouter } from 'next/router';
import MenuIcon from '@mui/icons-material/Menu';
import Offcanvas from './Offcanvas';

export default function SideMenu({}: SideMenuProps) {
    const [show, setShow] = useState(false);
    const router = useRouter();

    const handleShow = () => setShow(true);
    const handleClose = () => setShow(false);
    const handleLogout = () => router.push('/api/auth/logout');

    return (
        <div>
            <button type="button" onClick={handleShow} className="text-primary transition-transform hover:scale-105">
                <MenuIcon className="w-8 h-8" />
            </button>
            <Offcanvas show={show} onClose={handleClose} title={'Area'}>
                <div className='w-full h-full flex flex-col items-center px-8 py-4'>
                </div>
            </Offcanvas>
        </div>
    );
}

export interface SideMenuProps {}