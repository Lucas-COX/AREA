package discord

import (
	"Area/database/models"
	"Area/lib"
	"Area/services/types"
	"bytes"
	"encoding/gob"
)

type discordService struct {
	actions   []types.Action
	reactions []types.Reaction
}

func (*discordService) Authenticate(callback string, userId uint) string {
	return ""
}

func (*discordService) AuthenticateCallback(base64State string, code string) (string, error) {
	return "", nil
}

func (discord *discordService) GetActions() []types.Action {
	return discord.actions
}

func (discord *discordService) GetReactions() []types.Reaction {
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
		sendMessage(storedData, trigger.Action, trigger.ActionService)
	}
}

func (discord *discordService) ToJson() types.JsonService {
	return types.JsonService{
		Name:      discord.GetName(),
		Actions:   discord.GetActions(),
		Reactions: discord.GetReactions(),
	}
}

func New() *discordService {
	return &discordService{
		actions: []types.Action{},
		reactions: []types.Reaction{
			{Name: "send", Description: "Sends a message through a webhook url"},
		},
	}
}
