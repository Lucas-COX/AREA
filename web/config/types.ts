export type Service = {
    name: string,
    actions: Action[],
    reactions: Reaction[],
}

export type Action = {
    name: string,
    description: string,
}

export type Reaction = {
    name: string,
    description: string,
}

export type Trigger = {
    id: number,
    title: string,
    description?: string,
    active: boolean,
    created_at: Date,
    updated_at: Date,
    user_id: number,
    action_service: string,
    reaction_service: string,
    action: string,
    reaction: string,
}

export type User = {
    id: number,
    username: string,
    created_at: Date,
    updated_at: Date,
    triggers?: Trigger[],
    services: string[],
}

export type Session = {
    user?: User,
    authenticated: boolean,
    token?: string | null,
}
