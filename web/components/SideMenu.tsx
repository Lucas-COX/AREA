import clsx from 'clsx';
import { useState } from 'react';
import { useRouter } from 'next/router';
import MenuIcon from '@mui/icons-material/Menu';
import Offcanvas from './Offcanvas';
import Link from 'next/link';

export default function SideMenu({}: SideMenuProps) {
    const [show, setShow] = useState(false);
    const router = useRouter();

    const handleShow = () => setShow(true);
    const handleClose = () => setShow(false);
    const handleLogout = () => router.push('/api/auth/logout');
    const elements = [
        {name: "Home", href: "/"},
        {name: "Services", href: "/services"},
        {name: "Download APK", href: "/client.apk"},
    ]

    return (
        <div>
            <button type="button" onClick={handleShow} className="text-primary transition-transform hover:scale-105">
                <MenuIcon className="w-8 h-8" />
            </button>
            <Offcanvas show={show} onClose={handleClose} title={'Area'}>
                <div className='w-full h-full flex flex-col items-center pb-8'>
                    {elements.map((element, index) => (
                        <Link key={`link-${element.name}`} href={element.href}>
                            <div className={"w-full text-center hover:bg-primary/10 cursor-pointer border-b border-gray-200 py-4" + (index === 0 ? " border-t" : "")}>{element.name}</div>
                        </Link>
                    ))}
                </div>
            </Offcanvas>
        </div>
    );
}

export interface SideMenuProps {}