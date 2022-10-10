package reactions

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/gtuk/discordwebhook"
)

func React(reaction models.Reaction, user models.User) {
	trigger, err := database.Trigger.GetById(reaction.TriggerID, user.ID)
	var storedData models.TriggerData
	var buf bytes.Buffer

	lib.CheckError(err)
	buf.Write(trigger.Data)
	err = gob.NewDecoder(&buf).Decode(&storedData)

	lib.LogError(err)
	var username = "Area"
	var content = fmt.Sprintf("New Event : %s\t\n %s\t\n %s\t\n", storedData.Author, storedData.Title, storedData.Description)
	url := reaction.Token
	switch reaction.Action {
	case "send":
		message := discordwebhook.Message{
			Username: &username,
			Content:  &content,
		}
		err := discordwebhook.SendMessage(url, message)
		lib.LogError(err)
	}
}
