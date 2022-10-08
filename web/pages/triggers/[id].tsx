import { withSession } from '../../config/withs';
import type { GetServerSidePropsContext } from 'next';
import { useRouter } from 'next/router';
import { getSession } from '../../lib/session';
import AppLayout from '../../components/AppLayout';
import React from 'react'
import { useState } from 'react'
import { Paper, TextField, Button } from '@mui/material';
import { toast } from 'react-toastify';
import axios from 'axios'

export default function TriggerPage({ session }: TriggerProps) {
    const router = useRouter()
    const { id } = router.query
    const trigger = session?.user?.triggers?.find((t) => {return t.id === Number(id)})

    if (trigger === undefined)
      return router.push("/");

    const [state, setState] = useState({
      title: trigger.title,
      description: trigger?.description
    })

    const handleTitleChange = (e: any) => {
      setState({
        ...state,
        title: e.target.value
      })
    }
    const handleDescriptionChange = (e: any) => {
      setState({
        ...state,
        description: e.target.value
      })
    }

    const handleApply = async (e: any) => {
      try {
        const response = await toast.promise(axios.put(`${process.env.NEXT_PUBLIC_API_URL}/triggers/${trigger.id}`, {
          title: state.title,
          description: state.description,
        }, {
          headers: { 'Authorization': 'Bearer ' + (session.token as string), 'Content-Type': 'application/json' }
        }), {
          pending: "Loading...",
          error: "An error occured while updating trigger.",
          success: "Trigger successfully updated !"
        })
        router.back
      } catch (e) {
        console.error(e)
      }
    }
    console.log(trigger)

    return (
        <AppLayout type="centered" className="flex flex-col items-center justify-center bg-blue-50/50">
            <Paper className="w-2/3 h-2/3 flex flex-col justify-between p-10">
                <div className="flex flex-col space-y-6">
                    <TextField label="Title" variant="outlined" defaultValue={ trigger?.title } onChange={ handleTitleChange }/>
                    <TextField label="Description" multiline rows={4} defaultValue={ trigger?.description } onChange={ handleDescriptionChange }/>
                </div>
                <div className="justify-self-end space-x-4">
                      <Button variant="outlined" onClick={handleApply}>Apply changes</Button>
                      <Button variant="outlined" color={"error"} onClick={router.back}>Cancel</Button>
                </div>
            </Paper>
        </AppLayout>
    )
}


export interface HomeProps extends withSession {}

export async function getServerSideProps(context: GetServerSidePropsContext) {
  const session = await getSession(context)
  if (session.authenticated == false) {
    return {
      redirect: {
        destination: '/login',
        permanent: false,
      },
    }
  }
  const triggers = session.user?.triggers;
  if (triggers) {
    for (var i = 0; i < triggers.length; i++) {
        if (triggers[i].id === Number(context.query.id)) {
            return { props: { session }}
        }
    }
  }
  return {
    redirect: {
        destination: '/',
        permanent: false
    }
  }
}

export interface TriggerProps extends withSession {}
