package reactions

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"bytes"
	"encoding/gob"

	"github.com/gtuk/discordwebhook"
)

func React(reaction models.Reaction, user models.User) {
	trigger, err := database.Trigger.GetById(reaction.TriggerID, user.ID)
	var storedData models.TriggerData
	var buf bytes.Buffer
	var embeds []discordwebhook.Embed
	var fields []discordwebhook.Field

	lib.LogError(err)
	buf.Write(trigger.Data)
	err = gob.NewDecoder(&buf).Decode(&storedData)

	lib.LogError(err)
	var username = "Area"
	var title = "New event in Area"
	var color = "1668818"
	var timestamp = storedData.Timestamp.Format("January 2, 2006") + " at " + storedData.Timestamp.Format("15:04:05")
	var footer discordwebhook.Footer = discordwebhook.Footer{
		Text: &timestamp,
	}

	url := reaction.Token
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
