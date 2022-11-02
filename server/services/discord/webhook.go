package discord

import (
	"Area/database/models"
	"Area/lib"

	"github.com/gtuk/discordwebhook"
)

func sendMessage(storedData models.TriggerData, action string, service string) {
	var username = "Area"
	var title = "New " + string(action) + " in " + string(service)
	var color = "1668818"
	var embeds []discordwebhook.Embed
	var fields []discordwebhook.Field
	var timestamp = storedData.Timestamp.Format("January 2, 2006") + " at " + storedData.Timestamp.Format("15:04:05")
	var footer discordwebhook.Footer = discordwebhook.Footer{
		Text: &timestamp,
	}

	url := storedData.ReactionData

	fields = append(fields, discordwebhook.Field{
		Name:  &storedData.Title,
		Value: &storedData.Description,
	})
	embeds = append(embeds, discordwebhook.Embed{
		Title:       &title,
		Description: &storedData.Author,
		Fields:      &fields,
		Color:       &color,
		Footer:      &footer,
	})
	message := discordwebhook.Message{
		Username: &username,
		Embeds:   &embeds,
	}
	err := discordwebhook.SendMessage(url, message)
	lib.LogError(err)
}
