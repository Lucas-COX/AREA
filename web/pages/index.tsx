import React, { useState } from 'react';
import type { GetServerSidePropsContext } from 'next'
import { useRouter } from 'next/router';
import AppLayout from '../components/AppLayout'
import { getSession } from '../lib/session'
import { withSession } from '../config/withs'
import { IconButton, Card, Switch, CardActions, CardActionArea, CardContent, Typography, Button } from '@mui/material'
import DeleteIcon from '@mui/icons-material/Delete'
import AddIcon from '@mui/icons-material/Add';
import axios from 'axios'
import { toast } from 'react-toastify'

export default function Home({ session }: HomeProps) {
  
  const router = useRouter();
  const [state, setState] = useState({triggers: session.user?.triggers ? session.user.triggers : []});

  const handleCreate = async (e: any) => {
    const newTrigger = {
      title: "New trigger",
      user_id: session.user?.id,
    }
    try {
      const response = await toast.promise(axios.post(`${process.env.NEXT_PUBLIC_API_URL}/triggers`, {
        title: newTrigger.title,
        user_id: newTrigger.user_id,
        action: { type: "gmail", event: "receive" },
        reaction: { type: "discord", event: "send" },
      }, {
          headers: { 'Authorization': 'Bearer ' + (session.token as string), 'Content-Type': 'application/json' }
      }), {
        pending: "Loading...",
        error: "An error occured while creating trigger.",
        success: "Trigger successfully created !"
      })
      console.log(response);
      router.push(`/triggers/${response.data.trigger.id}`);
    } catch (e) {
      console.error(e)
    }
  }

  return (
    <AppLayout type="centered" className="flex flex-col space-y-4 bg-blue-50/50">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {state.triggers.map(function (trigger) {
            const handleDelete = async (e: any) => {
              try {
                await toast.promise(axios.delete(`${process.env.NEXT_PUBLIC_API_URL}/triggers/${trigger?.id}`, {
                  headers: { 'Authorization': 'Bearer ' + session.token }
                }), {
                  pending: 'Loading...',
                  error: 'An error occured while deleting trigger.',
                  success: 'Trigger successfully created !'
                })
                setState({ triggers: state.triggers.filter((t) => t.id !== trigger.id )});
              } catch (e) {
                console.error(e);
              }
            }
            return (
              <Card key={`trigger_${trigger.id}`} className="flex flex-col items-center">
                <CardActionArea onClick={function () {router.push(`/triggers/${trigger.id}`)}}>
                  <div>ici seront disposées les images</div>
                  <CardContent>
                    <Typography gutterBottom variant="h1" className="text-xl font-bold text-blue-400">
                      {trigger?.title}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      Dernières modifications : {trigger?.updated_at}
                    </Typography>
                  </CardContent>
                </CardActionArea>
                <CardActions className="flex w-full justify-between">
                  <IconButton aria-label="delete" onClick={handleDelete}>
                    <DeleteIcon />
                  </IconButton>
                  <Switch/>
                </CardActions>
              </Card>
          )})}
        </div>
        <Button variant="contained" color="success" startIcon={<AddIcon />} onClick={handleCreate}>Create Trigger</Button>
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
  return {
    props: { session }
  }
}
