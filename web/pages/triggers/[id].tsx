import { withSession } from '../../config/withs';
import type { GetServerSidePropsContext } from 'next';
import { useRouter } from 'next/router';
import { getSession } from '../../lib/session';
import AppLayout from '../../components/AppLayout';
import React from 'react'
import { Paper, TextField, Button } from '@mui/material';

export default function TriggerPage({ session }) {
    const router = useRouter()
    const { id } = router.query

    return (
        <AppLayout type="centered" className="flex flex-col">
            <div>Trigger {id}</div>
            <Paper>
                <div className="flex flex-col p-10 space-y-4">
                    <TextField label={id} placeholder="New username" variant="outlined"/>
                    <TextField label="Description" multiline rows={4} defaultValue="Default Value"/>
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
