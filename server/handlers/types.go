package handlers

type ErrorBody struct {
	Message string `json:"message"`
}

type AuthBody struct {
	Token string `json:"token"`
}
