package services

import (
	"Area/database/models"
	"Area/lib"
	"Area/services/reactions"
	"bytes"
	"encoding/gob"
)

type DiscordService struct {
	actions   []Action
	reactions []Reaction
}

func (discord *DiscordService) Authenticate(callback string, userId uint) string {
	return ""
}

func (discord *DiscordService) AuthenticateCallback(base64State string, code string) (string, error) {
	return "", nil
}

func (discord *DiscordService) GetActions() []Action {
	return discord.actions
}

func (discord *DiscordService) GetReactions() []Reaction {
	return discord.reactions
}

func (discord *DiscordService) GetName() string {
	return "discord"
}

func (discord *DiscordService) Check(action string, trigger models.Trigger) bool {
	return false
}

func (discord *DiscordService) React(reaction string, trigger models.Trigger) {
	var storedData models.TriggerData
	var buf bytes.Buffer

	buf.Write(trigger.Data)
	err := gob.NewDecoder(&buf).Decode(&storedData)
	lib.LogError(err)

	switch reaction {
	case "send":
		reactions.SendDiscordMessage(storedData, trigger.Action, trigger.ActionService)
	}
}

func (discord *DiscordService) ToJson() JsonService {
	return JsonService{
		Name:      discord.GetName(),
		Actions:   discord.GetActions(),
		Reactions: discord.GetReactions(),
	}
}

func NewDiscordService() *DiscordService {
	return &DiscordService{
		actions: []Action{},
		reactions: []Reaction{
			{Name: "send", Description: "Sends a message through a webhook url"},
		},
	}
}
