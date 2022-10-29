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
		case "outlook":
			triggered = services.Outlook.Check(v.Action, v)
		case "github":
			triggered = services.Github.Check(v.Action, v)
		case "notion":
			triggered = services.Notion.Check(v.Action, v)
		case "discord":
			triggered = services.Discord.Check(v.Action, v)
		default:
			triggered = false
		}
		if triggered {
			updated, _ := database.Trigger.GetById(v.ID, v.UserID)
			switch v.ReactionService {
			case "gmail":
				services.Gmail.React(v.Reaction, *updated)
			case "outlook":
				services.Outlook.React(v.Reaction, *updated)
			case "github":
				services.Github.Check(v.Reaction, *updated)
			case "notion":
				services.Notion.React(v.Reaction, *updated)
			case "discord":
				services.Discord.React(v.Reaction, *updated)
			}
		}
	}
}

func (j jobsManager) RunAsync() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(j.Do)
	s.StartAsync()
}
