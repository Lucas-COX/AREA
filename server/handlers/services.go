package handlers

import (
	"Area/lib"
	"Area/services"
	"Area/services/types"
	"net/http"
)

type servicesResponse struct {
	Services []types.JsonService `json:"services"`
}

func GetServices(w http.ResponseWriter, r *http.Request) {
	var resp servicesResponse
	resp.Services = []types.JsonService{
		services.Google.ToJson(),
		services.Discord.ToJson(),
		services.Microsoft.ToJson(),
		services.Github.ToJson(),
		services.Notion.ToJson(),
	}

	lib.SendJson(w, resp)
}
