package handlers

import (
	"Area/database"
	"Area/lib"
	"net/http"

	"github.com/jinzhu/copier"
)

type reactionsResponse struct {
	Reactions []ReactionResponseBody `json:"reactions"`
}

func GetReactions(w http.ResponseWriter, r *http.Request) {
	var resp reactionsResponse
	_, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	reactions, err := database.Reaction.Get()
	lib.CheckError(err)

	copier.Copy(&resp.Reactions, &reactions)
	lib.SendJson(w, resp)
}
