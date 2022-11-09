import type { GetServerSidePropsContext } from 'next';
import type { Trigger, Service } from '../../config/types';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import {
  Paper, TextField, Button, Switch, Select, InputLabel, MenuItem, SelectChangeEvent
} from '@mui/material';
import { toast } from 'react-toastify';
import axios from 'axios';
import { TrendingFlatOutlined } from '@mui/icons-material';
import AppLayout from '../../components/AppLayout';
import { getSession } from '../../lib/session';
import { withSession } from '../../config/withs';
import useServices from '../../hooks/useServices';
import Spinner from '../../components/Spinner';

interface TriggerPageState {
  trigger: Trigger;
}


export default function TriggerPage({ session }: TriggerProps) {
  const router = useRouter();
  const { id } = router.query;
  const trigger = session?.user?.triggers?.find((t) => t.id === Number(id));
  const [state, setState] = useState<TriggerPageState>({
    trigger: trigger as Trigger,
  });
  const {services, setServices, loading, error} = useServices(session.token as string)

  if (trigger === undefined) {
    return router.push('/');
  }

  const handleTitleChange = (e: any) => {
    if (state.trigger !== undefined) {
      setState({
        ...state,
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
        ...state,
        trigger: {
          ...state.trigger,
          description: e.target.value,
        },
      });
    }
  };

  const handleToggle = async (e: React.ChangeEvent<HTMLInputElement>) => {
    if (state.trigger !== undefined) {
      setState({
        ...state,
        trigger: {
          ...state.trigger,
          active: !state.trigger.active,
        },
      });
    }
  };
  const handleApply = async (e: any) => {
    try {
      console.log(state.trigger);
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

  const handleActionServiceChange = (e: SelectChangeEvent) => {
    var action = state.trigger.action
    if (e.target.value !== "undefined") {
      action = (services.find((s) => s.name === e.target.value) as Service).actions[0].name;
    }
    setState({ ...state, trigger: { ...state.trigger, action_service: e.target.value, action }});
  }
  const handleActionEventChange = (e: SelectChangeEvent) => {
    setState({ ...state,
      trigger: {
        ...state.trigger,
        action: e.target.value,
      }
    })
  }
  const handleReactionServiceChange = (e: SelectChangeEvent) => {
    var reaction = state.trigger.reaction
    if (e.target.value !== "undefined") {
      reaction = (services.find((s) => s.name === e.target.value) as Service).reactions[0].name;
    }
    setState({ ...state, trigger: { ...state.trigger, reaction_service: e.target.value, reaction }})
  }
  const handleReactionActionChange = (e: SelectChangeEvent) => {
    setState({ ...state,
      trigger: {
        ...state.trigger,
        reaction: e.target.value,
      }
    })
  }

  if (loading) {
    return (
      <AppLayout
        type="centered"
        loggedIn
      >
        <Spinner />
      </AppLayout>
    )
  }

  const filteredActionServices = services.filter((service) => session?.user?.services.includes(service.name) && service.actions.length !== 0)
  const filteredReactionServices = services.filter((service) => session?.user?.services.includes(service.name) && service.reactions.length !== 0)
  const filteredActions = filteredActionServices.filter((service) => service.name === state.trigger.action_service).map((service) => service.actions).flat()
  const filteredReactions = filteredReactionServices.filter((service) => service.name === state.trigger.reaction_service).map((service) => service.reactions).flat()
  filteredActionServices.push({ name: 'undefined', actions: [], reactions: [] })
  filteredReactionServices.push({ name: 'undefined', actions: [], reactions: [] })

  return (
    <AppLayout
      type="centered"
      className="flex flex-col items-center justify-center bg-blue-50/50"
      loggedIn
    >
      <Paper className="w-full h-full  sm:w-4/5 sm:h-4/5 md:w-3/4 md:h-3/4 flex flex-col justify-between p-10">
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
            <div className="flex w-full justify-evenly">
              <Select
                value={state.trigger.action_service ? state.trigger.action_service.toString() : "undefined"}
                label="Service"
                onChange={handleActionServiceChange}
              >
                {filteredActionServices.map((service) => (
                  <MenuItem key={`service-pick-${service.name}`} value={service.name}>{service.name.toUpperCase()}</MenuItem>
                ))}
              </Select>
              {filteredActions.length !== 0 && <Select
                value={state.trigger?.action || (filteredActions[0] && filteredActions[0].name)}
                label="Action"
                onChange={handleActionEventChange}
              >
              {filteredActions.map((action) => (
                  <MenuItem key={`action-${action.name}`} value={action.name}>{action.name}</MenuItem>
              ))}
              </Select>}
            </div>
          </div>
          <TrendingFlatOutlined fontSize="large" color="secondary" />
          <div className="flex flex-col space-y-4 items-center justify-evenly w-full h-1/2 border rounded-lg bg-secondary/5">
            <div>Reaction</div>
            <div className="flex w-full justify-evenly">
              <Select
                value={state.trigger.reaction_service ? state.trigger.reaction_service.toString() : "undefined"}
                label="Service"
                onChange={handleReactionServiceChange}
              >
                {filteredReactionServices.map((service) => (
                  <MenuItem key={`service-pick-${service.name}`} value={service.name}>{service.name.toUpperCase()}</MenuItem>
                ))}
              </Select>
              {filteredReactions.length !== 0 && <Select
                value={state.trigger?.reaction || (filteredReactions[0] && filteredReactions[0].name)}
                label="Action"
                onChange={handleReactionActionChange}
              >
              {filteredReactions.map((reaction) => (
                  <MenuItem key={`action-${reaction.name}`} value={reaction.name}>{reaction.name}</MenuItem>
              ))}
              </Select>}
            </div>
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
