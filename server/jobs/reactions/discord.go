package reactions

import (
	"Area/database/models"
	"Area/lib"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func React(reaction models.Reaction) {
	url := "https://discord.com/api/webhooks/1028682673511727164/alasz7LSPL1mgLHarHssiAbSZtm-pK2KSkTwROdMOROhE903UrKsYWAgRI17h7-TzHFP"
	switch reaction.Action {
	case "send":
		payload := new(bytes.Buffer)

		err := json.NewEncoder(payload).Encode("message")
		lib.CheckError(err)

		resp, err := http.Post(url, "application/json", payload)
		lib.CheckError(err)

		log.Printf(resp.Status)
	}
}
