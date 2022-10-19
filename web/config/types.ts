export type Action = {
    id: number,
    type: "gmail",
    event: String,
    token: String,
    created_at: Date,
    updated_at: Date,
    trigger_id: Number,
}

export type Reaction = {
    id: number,
    type: "discord",
    action: String,
    created_at: Date,
    updated_at: Date,
}

export type Trigger = {
    id: number,
    title: String,
    description?: String,
    active: boolean,
    created_at: Date,
    updated_at: Date,
    user_id: number,
    action_id: number,
    reaction_id: number,
    action?: Action,
    reaction?: Reaction,
}

export type User = {
    id: number,
    username: String,
    created_at: Date,
    updated_at: Date,
    triggers?: Trigger[],
    google_logged: boolean,
}

export type Session = {
    user?: User,
    authenticated: boolean,
    token?: string | null,
}
