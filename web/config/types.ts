export type Trigger = {
    id: Number,
    title: String,
    description?: String,
    created_at: Date,
    updated_at: Date,
    user_id: Number
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
}