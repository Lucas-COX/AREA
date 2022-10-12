import React, { useState } from 'react';
import type { GetServerSidePropsContext } from 'next'
import { useRouter } from 'next/router';
import AppLayout from '../components/AppLayout'
import { getSession } from '../lib/session'
import { withSession } from '../config/withs'
import { IconButton, Switch, CardActions, CardActionArea, CardContent, Typography, Button } from '@mui/material'
import DeleteIcon from '@mui/icons-material/Delete'
import AddIcon from '@mui/icons-material/Add';
import axios from 'axios'
import { toast } from 'react-toastify'
import Image from 'next/image';
import { TrendingFlatOutlined } from '@mui/icons-material';
import { gmail, discord } from '../lib/icons'
import Card from '../components/Card';

const icons = {
  "gmail": gmail,
  "discord": discord,
}

export default function Home({ session }: HomeProps) {

  const router = useRouter();
  const [state, setState] = useState({ triggers: session.user?.triggers ? session.user.triggers.sort((a, b) => (a.updated_at < b.updated_at) ? 1 : -1) : []});

  const handleCreate = async () => {
    const newTrigger = {
      title: "New trigger",
      user_id: session.user?.id,
      action: {
        type: 'gmail',
        event: 'receive',
        token: 'google'
      },
      reaction: {
        type: 'discord',
        action: 'send'
      }
    }
    try {
      const response = await toast.promise(axios.post(`${process.env.NEXT_PUBLIC_API_URL}/triggers`, {
        ...newTrigger
      }, {
          headers: { 'Authorization': 'Bearer ' + (session.token as string), 'Content-Type': 'application/json' }
      }), {
        pending: "Loading...",
        error: "An error occured while creating trigger.",
        success: "Trigger successfully created !"
      })
      router.push(`/triggers/${response.data.trigger.id}`);
    } catch (e) {
      console.error(e)
    }
  }
  const computeLastModified = (d: Date) => {
    const now = Date.now()
    const date = Date.parse(d.toString())
    const diff = Math.ceil(Math.abs(now.valueOf() - date.valueOf()) / (1000 * 60 * 60 * 24));
    return String(diff) + (diff < 1 ? " days" : " day") +  " ago";
  }
  // Todo: make trigger list scrollable

  return (
    <AppLayout type="centered" className="flex flex-col space-y-4 pb-10 pt-32 bg-blue-50/50" loggedIn={true}>
        <Button className="bg-white" variant="outlined" color="primary" startIcon={<AddIcon />} onClick={handleCreate}>Create Trigger</Button>
        <div className="flex flex-col space-y-4 p-4 w-full max-h-fit sm:w-3/4 overflow-y-scroll">
          {state.triggers.map(function (trigger) {
            const handleDelete = async () => {
              try {
                await toast.promise(axios.delete(`${process.env.NEXT_PUBLIC_API_URL}/triggers/${trigger?.id}`, {
                  headers: { 'Authorization': 'Bearer ' + session.token }
                }), {
                  pending: 'Loading...',
                  error: 'An error occured while deleting trigger.',
                  success: 'Trigger successfully deleted !'
                })
                setState({ triggers: state.triggers.filter((t) => t.id !== trigger.id )});
              } catch (e) {
                console.error(e);
              }
            }
            const handleToggle = async (e: React.ChangeEvent<HTMLInputElement>) => {
              try {
                await toast.promise(axios.put(`${process.env.NEXT_PUBLIC_API_URL}/triggers/${trigger?.id}`, {
                  active: !trigger.active,
                }, {
                  headers: { 'Authorization': 'Bearer ' + session.token, 'Content-Type': 'application/json' }
                }), {
                  pending: 'Loading...',
                  error: 'An error occured while turning trigger ' + (e.target.checked ? 'on.' : 'off.'),
                  success: 'Successfully turned trigger ' + (e.target.checked ? 'on.' : 'off.'),
                })
                setState({ triggers: state.triggers.map((t) => {
                  if (t.id !== trigger.id)
                    return (t)
                  else
                    return ({ ...t, active: !t.active })
                })})
              } catch (e) {
                console.error(e);
              }
            }

            return (
              <Card key={`trigger_${trigger.id}`} className="flex flex-col items-center h-64 bg-white">
                <CardActionArea onClick={function () {router.push(`/triggers/${trigger.id}`)}}>
                  <div className='flex items-center justify-evenly p-4'>
                    <div className="w-20 h-20">
                      <Image src={icons[trigger.action.type]} layout="responsive" />
                    </div>
                    <TrendingFlatOutlined fontSize='large' color="secondary" />
                    <div className="w-20 h-20">
                      <Image src={icons[trigger.reaction.type]} layout="responsive" />
                    </div>
                  </div>
                  <CardContent>
                    <div className="flex justify-between items-center">
                      <Typography gutterBottom className="text-xl font-bold text-blue-400">
                        {trigger?.title}
                      </Typography>
                      <Typography variant="body2" color="text.secondary">
                        {`Last edited : ${trigger.updated_at && computeLastModified(trigger.updated_at)}`}
                      </Typography>
                    </div>
                    <Typography gutterBottom variant="body2" color="text.secondary">
                      {trigger?.description}
                    </Typography>
                  </CardContent>
                </CardActionArea>
                <CardActions className="flex w-full justify-between">
                  <IconButton aria-label="delete" onClick={handleDelete}>
                    <DeleteIcon color="error" />
                  </IconButton>
                  <Switch color="secondary" onChange={handleToggle} checked={trigger.active} />
                </CardActions>
              </Card>
          )})}
        </div>
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
