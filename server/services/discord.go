package services

import (
	"Area/database/models"
	"Area/lib"
	"Area/services/reactions"
	"bytes"
	"encoding/gob"
)

type DiscordService struct {
	Actions   []Action
	Reactions []Reaction
}

func (discord *DiscordService) Authenticate(redirect string, callback string, userId uint) string {
	return ""
}

func (discord *DiscordService) AuthenticateCallback(base64State string, code string) (string, error) {
	return "", nil
}

func (discord *DiscordService) GetActions() []Action {
	return discord.Actions
}

func (discord *DiscordService) GetReactions() []Reaction {
	return discord.Reactions
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

func NewDiscordService() *DiscordService {
	return &DiscordService{
		Actions: []Action{},
		Reactions: []Reaction{
			{Name: "send", Description: "Sends a message through a webhook url"},
		},
	}
}
