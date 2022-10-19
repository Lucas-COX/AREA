import type { GetServerSidePropsContext } from 'next';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import {
  Paper, TextField, Button, Switch,
} from '@mui/material';
import { toast } from 'react-toastify';
import axios from 'axios';
import Image from 'next/image';
import InputAdornment from '@mui/material/InputAdornment';
import { TrendingFlatOutlined } from '@mui/icons-material';
import URLSafeBase64 from 'urlsafe-base64';
import icons from '../../lib/icons';
import AppLayout from '../../components/AppLayout';
import { getSession } from '../../lib/session';
import { withSession } from '../../config/withs';
import useActions from '../../hooks/useActions';
import useReactions from '../../hooks/useReactions';
import Spinner from '../../components/Spinner';


export default function TriggerPage({ session }: TriggerProps) {
  const router = useRouter();
  const { id } = router.query;
  const trigger = session?.user?.triggers?.find((t) => t.id === Number(id));
  const [state, setState] = useState({
    trigger,
  });
  const actionsState = useActions(session.token as string)
  const reactionsState = useReactions(session.token as string)

  if (trigger === undefined) {
    return router.push('/');
  }

  const handleTitleChange = (e: any) => {
    if (state.trigger !== undefined) {
      setState({
        trigger: {
          ...state.trigger,
          title: e.target.value,
        },
      });
    }
  };
  const handleDescriptionChange = (e: any) => {
    if (state.trigger !== undefined) {
      setState({
        trigger: {
          ...state.trigger,
          description: e.target.value,
        },
      });
    }
  };
  // const handleReactionTokenChange = (e: any) => {
  //   if (state.trigger !== undefined) {
  //     setState({
  //       trigger: {
  //         ...state.trigger,
  //         reaction: {
  //           ...state.trigger.reaction,
  //           token: e.target.value,
  //         },
  //       },
  //     });
  //   }
  // };
  const handleToggle = async (e: React.ChangeEvent<HTMLInputElement>) => {
    if (state.trigger !== undefined) {
      setState({
        trigger: {
          ...state.trigger,
          active: !state.trigger.active,
        },
      });
    }
  };
  const handleApply = async (e: any) => {
    try {
      await toast.promise(axios.put(`${process.env.NEXT_PUBLIC_API_URL}/triggers/${trigger.id}`, {
        ...state.trigger,
      }, {
        headers: { Authorization: `Bearer ${session.token as string}`, 'Content-Type': 'application/json' },
      }), {
        pending: 'Loading...',
        error: 'An error occured while updating trigger.',
        success: 'Trigger successfully updated !',
      });
      router.push('/');
    } catch (e) {
      console.error(e);
    }
  };


  const handleGoogleLogin = async () => {
    try {
      const location = `http://localhost:3000${router.asPath}`;
      const url = URLSafeBase64.encode(Buffer.from(location));

      const response = axios.get(`${process.env.NEXT_PUBLIC_API_URL}/providers/google/auth?callback=${url}`, {
        headers: { Authorization: `Bearer ${session.token}` },
      });
      router.push((await response).data.url);
    } catch (e) {
      toast.error('Failed to redirect to Google authentication page.');
    }
  };

  if (actionsState.loading || reactionsState.loading) {
    return (
      <AppLayout
        type="centered"
        loggedIn
      >
        <Spinner />
      </AppLayout>
    )
  }

  // const ActionIcon = <Image src={icons[actionsState.actions[trigger.action_id].type]} alt={`${actionsState.actions[trigger.action_id].type} icon`} width="15" height="15" />;
  // const ReactionIcon = <Image src={icons[reactionsState.reactions[trigger.reaction_id].type]} alt={`${reactionsState.reactions[trigger.action_id].type} icon`} width="15" height="15" />;

  return (
    <AppLayout
      type="centered"
      className="flex flex-col items-center justify-center bg-blue-50/50"
      loggedIn
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
          <div className="flex flex-col space-y-4 items-center justify-evenly w-full h-1/2 border rounded-lg bg-primary/5">
            <div>Action</div>
            {trigger.action_id && <Button
              variant="outlined"
              color="primary"
              // startIcon={ActionIcon}
              className="h-10"
              disabled={session.user?.google_logged}
            >
              {session.user?.google_logged ? 'Logged in' : `Login with ${actionsState.actions[trigger.action_id]}`}
            </Button>}
          </div>
          <TrendingFlatOutlined fontSize="large" color="secondary" />
          <div className="flex flex-col space-y-4 items-center justify-evenly w-full h-1/2 border rounded-lg bg-secondary/5">
            <div>Reaction</div>
            {/* <TextField
              label="Discord webhook url"
              className="bg-white"
              value={state.trigger && state.trigger.reaction.token}
              onChange={handleReactionTokenChange}
              InputProps={{
                className: 'h-10',
                startAdornment: (
                  <InputAdornment position="start">
                    {ReactionIcon}
                  </InputAdornment>
                ),
              }}
            /> */}
          </div>
        </div>
        <div className="justify-self-end space-x-4">
          <Button variant="outlined" onClick={handleApply}>
            Apply changes
          </Button>
          <Button variant="outlined" color="error" onClick={() => router.push('/')}>
            Cancel
          </Button>
          <Switch color="secondary" checked={state.trigger && state.trigger.active} onChange={handleToggle} />
        </div>
      </Paper>
    </AppLayout>
  );
}

export type HomeProps = withSession

export async function getServerSideProps(context: GetServerSidePropsContext) {
  const session = await getSession(context);
  if (session.authenticated == false) {
    return {
      redirect: {
        destination: '/login',
        permanent: false,
      },
    };
  }
  const triggers = session.user?.triggers;
  if (triggers) {
    for (let i = 0; i < triggers.length; i++) {
      if (triggers[i].id === Number(context.query.id)) {
        return { props: { session } };
      }
    }
  }
  return {
    redirect: {
      destination: '/',
      permanent: false,
    },
  };
}

export type TriggerProps = withSession
