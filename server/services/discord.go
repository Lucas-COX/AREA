package services

import (
	"Area/database/models"
	"Area/lib"
	"Area/services/reactions"
	"bytes"
	"encoding/gob"
)

type discordService struct {
	actions   []Action
	reactions []Reaction
}

func (*discordService) Authenticate(callback string, userId uint) string {
	return ""
}

func (*discordService) AuthenticateCallback(base64State string, code string) (string, error) {
	return "", nil
}

func (discord *discordService) GetActions() []Action {
	return discord.actions
}

func (discord *discordService) GetReactions() []Reaction {
	return discord.reactions
}

func (discord *discordService) GetName() string {
	return "discord"
}

func (*discordService) Check(action string, trigger models.Trigger) bool {
	return false
}

func (*discordService) React(reaction string, trigger models.Trigger) {
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

func (discord *discordService) ToJson() JsonService {
	return JsonService{
		Name:      discord.GetName(),
		Actions:   discord.GetActions(),
		Reactions: discord.GetReactions(),
	}
}

func NewDiscordService() *discordService {
	return &discordService{
		actions: []Action{},
		reactions: []Reaction{
			{Name: "send", Description: "Sends a message through a webhook url"},
		},
	}
}
