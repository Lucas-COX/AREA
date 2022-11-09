package handlers

import (
	"Area/database"
	"Area/lib"
	"Area/services"
	"net/http"

	"github.com/jinzhu/copier"
)

type userResponse struct {
	User userBody `json:"me"`
}

func Me(w http.ResponseWriter, r *http.Request) {
	var resp userResponse
	user, err := database.User.GetFromContext(r.Context())
	lib.CheckError(err)

	copier.Copy(&resp.User, &user)

	resp.User.Services = []string{}
	if user.GoogleToken != "" {
		resp.User.Services = append(resp.User.Services, services.Google.GetName())
	}
	if user.MicrosoftToken != "" {
		resp.User.Services = append(resp.User.Services, services.Microsoft.GetName())
	}
	if user.GithubToken != "" {
		resp.User.Services = append(resp.User.Services, services.Github.GetName())
	}
	if user.NotionToken != "" {
		resp.User.Services = append(resp.User.Services, services.Notion.GetName())
	}
	if user.DiscordEnabled {
		resp.User.Services = append(resp.User.Services, services.Discord.GetName())
	}
	if user.TimerEnabled {
		resp.User.Services = append(resp.User.Services, services.Timer.GetName())
	}
	triggers, err := database.Trigger.Get(user.ID)
	lib.CheckError(err)

	copier.Copy(&resp.User.Triggers, triggers)
	lib.SendJson(w, resp)
}
