import { withSession } from '../../config/withs';
import type { GetServerSidePropsContext } from 'next';
import { useRouter } from 'next/router';
import { getSession } from '../../lib/session';
import AppLayout from '../../components/AppLayout';
import React from 'react'
import { useState } from 'react'
import { Paper, TextField, Button, Switch } from '@mui/material';
import { toast } from 'react-toastify';
import axios from 'axios'
import icons from '../../lib/icons';
import Image from 'next/image';
import InputAdornment from "@mui/material/InputAdornment";
import { TrendingFlatOutlined } from '@mui/icons-material';
import URLSafeBase64 from 'urlsafe-base64'


export default function TriggerPage({ session }: TriggerProps) {
    const router = useRouter()
    const { id } = router.query
    const trigger = session?.user?.triggers?.find((t) => {return t.id === Number(id)})
    const [state, setState] = useState({
      trigger: trigger,
    })

    if (trigger === undefined) {
      return router.push("/");
    }

    const handleTitleChange = (e: any) => {
      if (state.trigger !== undefined) {
        setState({
          trigger: {
            ...state.trigger,
            title: e.target.value
          }
        })
      }
    }
    const handleDescriptionChange = (e: any) => {
      if (state.trigger !== undefined) {
        setState({
          trigger: {
            ...state.trigger,
            description: e.target.value
          }
        })
      }
    }
    const handleReactionTokenChange = (e: any) => {
      if (state.trigger !== undefined) {
        setState({
          trigger: {
            ...state.trigger,
            reaction: {
              ...state.trigger.reaction,
              token: e.target.value,
            }
          }
        })
      }
    }
    const handleToggle = async (e: React.ChangeEvent<HTMLInputElement>) => {
      if (state.trigger !== undefined) {
        setState({ trigger: {
          ...state.trigger,
          active: !state.trigger.active
        }})
      }
    }
    const handleApply = async (e: any) => {
      try {
        await toast.promise(axios.put(`${process.env.NEXT_PUBLIC_API_URL}/triggers/${trigger.id}`, {
          ...state.trigger
        }, {
          headers: { 'Authorization': 'Bearer ' + (session.token as string), 'Content-Type': 'application/json' }
        }), {
          pending: "Loading...",
          error: "An error occured while updating trigger.",
          success: "Trigger successfully updated !"
        })
        router.push("/")
      } catch (e) {
        console.error(e)
      }
    }
    const handleGoogleLogin = async () => {
      try {
        const location = "http://localhost:3000" + router.asPath
        const url = URLSafeBase64.encode(Buffer.from(location));

        const response = axios.get(`${process.env.NEXT_PUBLIC_API_URL}/providers/google/auth?callback=${url}`, {
          headers: { 'Authorization': 'Bearer ' + session.token }
        })
        router.push((await response).data.url)
      } catch (e) {
        toast.error("Failed to redirect to Google authentication page.")
      }
    }

    const ActionIcon = <Image src={icons[trigger.action.type]} alt={`${trigger.action.type} icon`} width="15" height="15" />
    const ReactionIcon = <Image src={icons[trigger.reaction.type]} alt={`${trigger.reaction.type} icon`} width="15" height="15" />

    return (
      <AppLayout
        type="centered"
        className="flex flex-col items-center justify-center bg-blue-50/50"
        loggedIn={true}
      >
        <Paper className="w-2/3 h-2/3 flex flex-col justify-between p-10">
          <div className="flex flex-col space-y-6">
            <TextField
              label="Title"
              variant="outlined"
              defaultValue={trigger?.title}
              onChange={handleTitleChange}
            />
            <TextField
              label="Description"
              multiline
              rows={4}
              defaultValue={trigger?.description}
              onChange={handleDescriptionChange}
            />
          </div>
          <div className="h-full flex items-center space-x-6">
            <div className={"flex flex-col space-y-4 items-center justify-evenly w-full h-1/2 border rounded-lg bg-primary/5"}>
              <div>Action</div>
              <Button
                variant="outlined"
                color="primary"
                startIcon={ActionIcon}
                className="h-10"
                disabled={session.user?.google_logged}
                onClick={handleGoogleLogin}
              >
                {session.user?.google_logged ? "Logged in" : "Login with Google"}
              </Button>
            </div>
            <TrendingFlatOutlined fontSize='large' color="secondary" />
            <div className={"flex flex-col space-y-4 items-center justify-evenly w-full h-1/2 border rounded-lg bg-secondary/5"}>
              <div>Reaction</div>
              <TextField
                label={"Discord webhook url"}
                className="bg-white"
                value={state.trigger && state.trigger.reaction.token}
                onChange={handleReactionTokenChange}
                InputProps={{
                  className: "h-10",
                  startAdornment: (
                    <InputAdornment position="start">
                      {ReactionIcon}
                    </InputAdornment>
                  ),
                }}
              />
            </div>
          </div>
          <div className="justify-self-end space-x-4">
            <Button variant="outlined" onClick={handleApply}>
              Apply changes
            </Button>
            <Button variant="outlined" color={"error"} onClick={() => router.push("/")}>
              Cancel
            </Button>
            <Switch color={"secondary"} checked={state.trigger && state.trigger.active} onChange={handleToggle} />
          </div>
        </Paper>
      </AppLayout>
    );
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
