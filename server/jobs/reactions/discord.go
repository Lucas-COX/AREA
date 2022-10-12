package reactions

import (
	"Area/database/models"
	"Area/lib"
	"bytes"
	"encoding/gob"

	"github.com/gtuk/discordwebhook"
)

func React(reaction models.Reaction, trigger models.Trigger, user models.User) {
	var storedData models.TriggerData
	var buf bytes.Buffer
	var embeds []discordwebhook.Embed
	var fields []discordwebhook.Field

	buf.Write(trigger.Data)
	err := gob.NewDecoder(&buf).Decode(&storedData)

	lib.LogError(err)
	var username = "Area"
	var title = "New event in Area"
	var color = "1668818"
	var timestamp = storedData.Timestamp.Format("January 2, 2006") + " at " + storedData.Timestamp.Format("15:04:05")
	var footer discordwebhook.Footer = discordwebhook.Footer{
		Text: &timestamp,
	}

	url := storedData.ReactionToken
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
	switch reaction.Action {
	case "send":
		message := discordwebhook.Message{
			Username: &username,
			Embeds:   &embeds,
		}
		err := discordwebhook.SendMessage(url, message)
		lib.LogError(err)
	}
}
