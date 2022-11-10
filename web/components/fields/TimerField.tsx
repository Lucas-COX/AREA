import { Input, TextField } from '@mui/material';
import { ChangeEvent } from 'react';
import { asTextField } from '../../config/withs';

export function TimerMinuteField({value = "", onChange = () => {}}: TimerFieldProps) {
    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => { onChange(e.target.value) }
    return (
        <TextField className="bg-white" size="small" type="number" value={value} onChange={handleChange} label="Minutes" />
    )
}

export function TimerTimeField({ value= "", onChange=() => {}}: TimerFieldProps) {
    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => { onChange(e.target.value) }
    return (
        <Input type="time" className="bg-white p-1 rounded-md" value={value} onChange={handleChange}></Input>
    )
}

export function TimerDateTimeField({ value= "", onChange=() => {}}: TimerFieldProps) {
    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => { onChange(e.target.value) }
    return (
        <Input type="datetime-local" className="bg-white p-1 rounded-md" value={value} onChange={handleChange}></Input>
    )
}

export interface TimerFieldProps extends asTextField {}