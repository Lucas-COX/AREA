import type { GetServerSidePropsContext } from 'next'
import { useState } from 'react';
import AppLayout from '../components/AppLayout'
import { getSession } from '../lib/session'
import { withSession } from '../config/withs'
import { IconButton, Card, Switch, CardActions, CardActionArea, CardContent, Typography } from '@mui/material'
import DeleteIcon from '@mui/icons-material/Delete'
import axios from 'axios'

export default function Home({ session }: HomeProps) {

  const trigger = session.user?.triggers ? session.user?.triggers[0] : undefined;
  
  const handleDelete = (e: any) => {
    axios.delete(`${process.env.NEXT_PUBLIC_API_URL}/triggers/${trigger?.id}`)
  }

  console.log(trigger);
  return (
    <AppLayout type="centered" className="flex flex-col">
        <div>Hello {session.user?.username}</div>
        <Card className="flex flex-col items-center">
          <CardActionArea>
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
          <CardActions className="space-x-60">
            <IconButton aria-label="delete" onClick={handleDelete}>
              <DeleteIcon />
            </IconButton>
            <Switch/>
          </CardActions>
        </Card>
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
