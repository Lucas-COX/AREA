import { TextField } from '@mui/material';
import { ChangeEvent } from 'react';
import { asTextField } from '../../config/withs';

export default function GithubField({value = "", onChange = () => {}}: GithubFieldProps) {
    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => { onChange(e.target.value) }
    return (
        <TextField className="bg-white" size="small" value={value} onChange={handleChange} label="Repository name"></TextField>
    )
}

export interface GithubFieldProps extends asTextField {}