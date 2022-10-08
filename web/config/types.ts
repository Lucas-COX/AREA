export type Action = {
    id: Number,
    type: String,
    event: String,
    token: String,
    created_at: Date,
    updated_at: Date,
    trigger_id: Number,
}

export type Reaction = {
    id: Number,
    type: String,
    action: String,
    token: String,
    created_at: Date,
    updated_at: Date,
    trigger_id: Number,
}

export type Trigger = {
    id: Number,
    title: String,
    description?: String,
    active: boolean,
    created_at: Date,
    updated_at: Date,
    user_id: Number,
    action: Action,
    reaction: Reaction,
}

export type User = {
    id: Number,
    username: String,
    created_at: Date,
    updated_at: Date,
    triggers?: Trigger[],
}

export type Session = {
    user?: User,
    authenticated: boolean,
    token?: String | null,
}
