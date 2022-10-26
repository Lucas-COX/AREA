package jobs

import (
	"Area/database"
	"Area/lib"
	"Area/services"
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
	triggers, err := database.Trigger.GetActive()
	lib.LogError(err)

	for _, v := range triggers {
		switch v.ActionService {
		case "gmail":
			triggered = services.Gmail.Check(v.Action, v)
		default:
			triggered = false
		}
		if triggered {
			switch v.ReactionService {
			case "discord":
				services.Discord.React(v.Reaction, v)
			}
		}
	}
}

func (j jobsManager) RunAsync() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(j.Do)
	s.StartAsync()
}
