import clsx from 'clsx';
import CloseIcon from '@mui/icons-material/Close';
import { withChildren, withShow, withTitle, withOnClose } from '../config/withs';

export default function Offcanvas({
    children, show = false, onClose, title = '',
}: OffcanvasProps) {
    const handleClose = () => { if (show && onClose) onClose(); };
    const handleClick = (e: any) => e.stopPropagation();

    return (
        <div>
            <div
                className={clsx(
                    'top-0 left-0 flex h-screen w-screen z-50 bg-black/40',
                    'transition-all duration-700 ease-in-out',
                    show ? 'fixed' : 'hidden',
                )}
                onClick={handleClose}
            />
            <div
                className={clsx(
                    'fixed top-0 left-0 h-screen bg-gray-100 z-50 w-screen sm:w-96',
                    'transition-all duration-300 ease-in-out',
                    show ? 'translate-x-0' : '-translate-x-full',
                )}
                onClick={handleClick}
            >
                <div className="w-full flex items-center border-b border-gray-300 h-24 px-8">
                    <div className='text-center text-xl text-primary w-full'>{title}</div>
                    <div className='flex w-full items-center justify-end absolute px-16'>
                        <button onClick={handleClose}>
                            <CloseIcon />
                        </button>
                    </div>
                </div>
                {children}
            </div>
        </div>
    );
}

export interface OffcanvasProps extends withChildren, withShow, withOnClose, withTitle {}
