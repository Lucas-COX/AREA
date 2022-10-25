package jobs

import (
	"Area/database"
	"Area/database/models"
	"Area/jobs/actions"
	"Area/jobs/reactions"
	"Area/lib"
	"time"

	"github.com/go-co-op/gocron"
)

// classe JobsManager
type jobsManager struct{}

type JobsManager interface {
	RunAsync()
	RunSync()
	Do()
}

func NewManager() jobsManager {
	return jobsManager{}
}

func (j jobsManager) RunSync() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(j.Do)
	s.StartBlocking()
}

func (jobsManager) Do() {
	var triggered bool
	triggers, err := database.Trigger.GetActive(true)
	lib.LogError(err)

	for _, v := range triggers {
		switch v.Action.Type {
		case models.GmailAction:
			triggered = actions.CheckGmailAction(v.Action, v, v.User)
		default:
			triggered = false
		}
		if triggered {
			switch v.Reaction.Type {
			case models.DiscordReaction:
				reactions.Discord(v.Reaction, v, v.User)
			}
		}
	}
}

func (j jobsManager) RunAsync() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(j.Do)
	s.StartAsync()
}
