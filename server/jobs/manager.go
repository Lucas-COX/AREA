package jobs

import (
	"Area/database"
	"Area/database/models"
	"Area/jobs/actions"
	"Area/jobs/reactions"
	"Area/lib"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-co-op/gocron"
)

// classe JobsManager
type jobsManager struct{}

type JobsManager interface {
	RunAsync()
	RunSync()
	Do()
}

// méthode New() -> Manager
// méthode RunAsync() -> Lance une go-routine de la méthode RunSync() -> check toutes les minutes les triggers qui sont activés et
// exécuter les tâches correspondantes
// Quand tu lances le traitement d'un nouveau trigger utilise des goroutines pour accélérer le process

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
	lib.CheckError(err)
	spew.Dump(triggers)
	for _, v := range triggers {
		switch v.Action.Type {
		case models.GmailAction:
			triggered = actions.CheckGmailAction(v.Action, v.User)
		default:
			triggered = false
		}
		if triggered {
			switch v.Reaction.Type {
			case models.DiscordReaction:
				reactions.React(v.Reaction)
			}
		}
	}
}

func (j jobsManager) RunAsync() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(j.Do)
	s.StartAsync()
}
