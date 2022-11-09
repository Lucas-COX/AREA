package handlers

import (
	"Area/lib"
	"Area/services"
	"Area/services/types"
	"net"
	"net/http"
	"time"
)

type server struct {
	CurrentTime int64               `json:"current_time"`
	Services    []types.JsonService `json:"services"`
}

type client struct {
	Host string `json:"host"`
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	var resp aboutResponse
	var err error
	resp.Client.Host, _, err = net.SplitHostPort(lib.ReadUserIP(r))
	if err != nil {
		resp.Client.Host = lib.ReadUserIP(r)
	}
	resp.Server.CurrentTime = time.Now().UTC().UnixMilli()
	resp.Server.Services = []types.JsonService{
		services.Discord.ToJson(),
		services.Timer.ToJson(),
		services.Google.ToJson(),
		services.Github.ToJson(),
		services.Microsoft.ToJson(),
		services.Notion.ToJson(),
	}
	lib.SendJson(w, resp)
}
