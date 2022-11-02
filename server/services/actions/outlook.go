package actions

// import (
// 	"Area/database"
// 	m "Area/database/models"
// 	"Area/lib"
// 	"bytes"
// 	"context"
// 	"encoding/gob"
// 	"os"

// 	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
// 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
// 	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
// 	"github.com/microsoftgraph/msgraph-sdk-go/me/messages"
// 	"github.com/microsoftgraph/msgraph-sdk-go/me/messages/item"
// 	"github.com/microsoftgraph/msgraph-sdk-go/models"
// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/microsoft"
// )

// type tokenProvider struct {
// 	token azcore.AccessToken
// }

// func (t tokenProvider) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
// 	return t.token, nil
// }

// func CreateOutlookConnection(refresh_token string) *msgraphsdkgo.GraphServiceClient {
// 	conf := &oauth2.Config{
// 		ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
// 		ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
// 		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL") + "/providers/microsoft/callback",
// 		Scopes: []string{
// 			"https://graph.microsoft.com/Mail.Read",
// 			"https://graph.microsoft.com/User.Read",
// 			"offline_access",
// 		},
// 		Endpoint: microsoft.AzureADEndpoint("consumers"),
// 	}
// 	var token oauth2.Token = oauth2.Token{
// 		RefreshToken: refresh_token,
// 	}
// 	if refresh_token == "" {
// 		return nil
// 	}

// 	tmp, _ := conf.TokenSource(context.Background(), &token).Token()
// 	provider := tokenProvider{token: azcore.AccessToken{
// 		Token: tmp.AccessToken,
// 	}}
// 	srv, _ := msgraphsdkgo.NewGraphServiceClientWithCredentials(provider, []string{})
// 	return srv
// }

// func fetchLastOutlookReceive(srv *msgraphsdkgo.GraphServiceClient) *models.Messageable {
// 	options := messages.MessagesRequestBuilderGetRequestConfiguration{
// 		Headers: map[string]string{
// 			"Prefer": "outlook.body-content-type=\"text\"",
// 		},
// 		QueryParameters: &messages.MessagesRequestBuilderGetQueryParameters{
// 			Select: []string{"body", "sender", "subject"},
// 		},
// 	}
// 	itemOptions := item.MessageItemRequestBuilderGetRequestConfiguration{
// 		Headers: map[string]string{
// 			"Prefer": "outlook.body-content-type=\"text\"",
// 		},
// 		QueryParameters: &item.MessageItemRequestBuilderGetQueryParameters{
// 			Select: []string{"body", "sender", "subject", "bodyPreview", "receivedDateTime"},
// 		},
// 	}

// 	result, err := srv.Me().Messages().Get(context.Background(), &options)
// 	if err != nil {
// 		lib.LogError(err)
// 		return nil
// 	}
// 	id := result.GetValue()[0].GetId()
// 	mail, err := srv.Me().MessagesById(*id).Get(context.Background(), &itemOptions)

// 	if err != nil {
// 		lib.LogError(err)
// 		return nil
// 	}
// 	return &mail
// }

// func compareOutlookData(newData m.TriggerData, oldData m.TriggerData, mail models.Messageable, trigger *m.Trigger) bool {
// 	newData.Timestamp = *mail.GetReceivedDateTime()
// 	if trigger.Data == nil || oldData.Timestamp.Before(newData.Timestamp) {
// 		newData.Description = *mail.GetBodyPreview()
// 		newData.Author = *mail.GetSender().GetEmailAddress().GetAddress()
// 		newData.Title = *mail.GetSubject()
// 		newData.ActionData = oldData.ActionData
// 		newData.ReactionData = oldData.ReactionData
// 		if trigger.Data != nil {
// 			trigger.Data = lib.EncodeToBytes(newData)
// 			database.Trigger.Update(trigger)
// 			return true
// 		}
// 		trigger.Data = lib.EncodeToBytes(newData)
// 		database.Trigger.Update(trigger)
// 	}
// 	return false
// }

// func OutlookReceive(srv *msgraphsdkgo.GraphServiceClient, triggerId uint, userId uint) bool {
// 	var newData m.TriggerData
// 	var storedData m.TriggerData
// 	var mail = fetchLastOutlookReceive(srv)
// 	var buf bytes.Buffer

// 	trigger, err := database.Trigger.GetById(triggerId, userId)
// 	lib.LogError(err)
// 	buf.Write(trigger.Data)

// 	gob.NewDecoder(&buf).Decode(&storedData)
// 	if mail == nil {
// 		return false
// 	}
// 	return compareOutlookData(newData, storedData, *mail, trigger)
// }
