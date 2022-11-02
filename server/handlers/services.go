package handlers

import (
	"Area/lib"
	"Area/services"
	"net/http"
)

type servicesResponse struct {
	Services []services.JsonService `json:"services"`
}

func GetServices(w http.ResponseWriter, r *http.Request) {
	var resp servicesResponse
	resp.Services = []services.JsonService{
		services.Google.ToJson(),
		services.Discord.ToJson(),
		services.Microsoft.ToJson(),
		services.Github.ToJson(),
		services.Notion.ToJson(),
	}

	lib.SendJson(w, resp)
}
