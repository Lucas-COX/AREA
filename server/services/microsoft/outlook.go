package microsoft

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

type Sender struct {
	ID           string `json:"id"`
	EmailAddress struct {
		Address string `json:"address"`
	} `json:"emailAddress"`
}

type Message struct {
	ID               string    `json:"id"`
	Subject          string    `json:"subject"`
	BodyPreview      string    `json:"bodyPreview"`
	Sender           Sender    `json:"sender"`
	ReceivedDateTime time.Time `json:"receivedDateTime"`
}

type microsoftListResponse struct {
	Value []Message `json:"value"`
}

func createMicrosoftConnection(refresh_token string) *http.Client {
	var conf = &oauth2.Config{
		ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
		ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
		Scopes: []string{
			"https://graph.microsoft.com/Mail.Read",
			"https://graph.microsoft.com/User.Read",
			"offline_access",
		},
		Endpoint: microsoft.AzureADEndpoint("consumers"),
	}
	var token oauth2.Token = oauth2.Token{
		RefreshToken: refresh_token,
	}
	if refresh_token == "" {
		return nil
	}
	return conf.Client(context.Background(), &token)
}

func fetchLastOutlookMessage(srv *http.Client) *Message {
	var list microsoftListResponse
	r := http.Request{
		Header: map[string][]string{
			"Prefer": {"microsoft.body-content-type=\"text\""},
		},
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme:   "https",
			Host:     "graph.microsoft.com",
			Path:     "/v1.0/me/messages",
			RawQuery: "$select=sender,subject,receivedDateTime,bodyPreview&$top=1&$orderby=receivedDateTime%20desc",
		},
	}
	res, err := srv.Do(&r)

	if res.StatusCode != http.StatusOK {
		lib.LogError(err)
		return nil
	}

	err = json.NewDecoder(res.Body).Decode(&list)
	if err != nil {
		lib.LogError(err)
		return nil
	}
	if len(list.Value) == 0 {
		return nil
	} else {
		return &list.Value[0]
	}
}

func compareOutlookData(newData models.TriggerData, oldData models.TriggerData, mail Message, trigger *models.Trigger) bool {
	newData.Timestamp = mail.ReceivedDateTime
	if trigger.Data == nil || oldData.Timestamp.Before(newData.Timestamp) {
		newData.Description = mail.BodyPreview
		newData.Author = mail.Sender.EmailAddress.Address
		newData.Title = mail.Subject
		newData.ActionData = oldData.ActionData
		newData.ReactionData = oldData.ReactionData
		if trigger.Data != nil {
			trigger.Data = lib.EncodeToBytes(newData)
			database.Trigger.Update(trigger)
			return true
		}
		trigger.Data = lib.EncodeToBytes(newData)
		database.Trigger.Update(trigger)
	}
	return false
}

func checkOutlookReceive(srv *http.Client, triggerId uint, userId uint) bool {
	var newData models.TriggerData
	var storedData models.TriggerData
	var mail = fetchLastOutlookMessage(srv)
	var buf bytes.Buffer

	trigger, err := database.Trigger.GetById(triggerId, userId)
	lib.LogError(err)
	buf.Write(trigger.Data)

	gob.NewDecoder(&buf).Decode(&storedData)
	if mail == nil {
		return false
	}
	return compareOutlookData(newData, storedData, *mail, trigger)
}
