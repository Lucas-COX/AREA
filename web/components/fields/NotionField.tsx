import { TextField } from '@mui/material';
import { ChangeEvent } from 'react';
import { asTextField } from '../../config/withs';

export default function NotionField({value = "", onChange = () => {}}: NotionFieldProps) {
    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => { onChange(e.target.value) }
    return (
        <TextField className="bg-white" size="small" value={value} onChange={handleChange} label={"Notion page url"} />
    )
}

export interface NotionFieldProps extends asTextField {}