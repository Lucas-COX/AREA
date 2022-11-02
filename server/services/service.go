package services

import "Area/database/models"

type Action struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Reaction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OauthState struct {
	Callback string `json:"callback"`
	UserId   uint   `json:"user_id"`
}

type Service interface {
	Authenticate(callback string, userId uint) string                     // returns the url to start the authentication process
	AuthenticateCallback(base64State string, code string) (string, error) // handles the authentication success or failure
	GetActions() []Action                                                 // returns the actions the service handles
	GetReactions() []Reaction                                             // returns the reactions the service handles
	GetName() string                                                      // returns the name of the service
	Check(action string, trigger models.Trigger) bool                     // checks if the action happened
	React(reaction string, trigger models.Trigger)                        // executes the reaction
	ToJson() JsonService                                                  // returns the json representation of the service
}

type JsonService struct {
	Name      string     `json:"name"`
	Actions   []Action   `json:"actions"`
	Reactions []Reaction `json:"reactions"`
}

var Google Service = NewGoogleService()
var Discord Service = NewDiscordService()
var Microsoft Service = NewMicrosoftService()
var Github Service = NewGithubService()
var Notion Service = NewNotionService()

func Get() []Service {
	var result = []Service{
		Google,
		Discord,
	}
	return result
}
