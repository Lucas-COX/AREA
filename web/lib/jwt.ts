export function getToken(): string | null {
    return localStorage.getItem("epytodo_token")
}

export function setToken(token: string) {
    localStorage.setItem("epytodo_token", token)
}
