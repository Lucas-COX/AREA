import { TextField } from '@mui/material';
import { ChangeEvent } from 'react';
import { asTextField } from '../../config/withs';

export default function DiscordField({value = "", onChange = () => {}}: DiscordFieldProps) {
    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => { onChange(e.target.value) }
    return (
        <TextField className="bg-white" size="small" label="Webhook url" value={value} onChange={handleChange}></TextField>
    )
}

export interface DiscordFieldProps extends asTextField {}