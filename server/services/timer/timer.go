package timer

import (
	"Area/database/models"
	"Area/services/types"
	"bytes"
	"encoding/gob"
	"time"
)

type timerService struct {
	actions   []types.Action
	reactions []types.Reaction
}

func (*timerService) Authenticate(callback string, userId uint) string {
	return ""
}

func (*timerService) AuthenticateCallback(base64State string, code string) (string, error) {
	return "", nil
}

func (timer *timerService) GetActions() []types.Action {
	return timer.actions
}

func (timer *timerService) GetReactions() []types.Reaction {
	return nil
}

func (timer *timerService) GetName() string {
	return "timer"
}

func (*timerService) Check(action string, trigger models.Trigger) bool {

	currentTime := time.Now()

	var storedData models.TriggerData
	var buf bytes.Buffer

	buf.Write(trigger.Data)

	gob.NewDecoder(&buf).Decode(&storedData)

	switch action {
	case "every x minutes":
		return checkTimeInterval(currentTime, storedData, &trigger)
	case "everyday at":
		return checkEveryDayTime(currentTime, storedData, &trigger)
	case "single time":
		return checkSingleTime(currentTime, storedData, &trigger)
	}
	return false
}

func (*timerService) React(reaction string, trigger models.Trigger) {

}

func (timer *timerService) ToJson() types.JsonService {
	return types.JsonService{
		Name:      timer.GetName(),
		Actions:   timer.GetActions(),
		Reactions: timer.GetReactions(),
	}
}

func NewTimerService() *timerService {
	return &timerService{
		actions: []types.Action{
			{Name: "every x minutes", Description: "When x minutes has passed"},
			{Name: "everyday at", Description: "every day at x time"},
			{Name: "single time", Description: "When it's x time"},
		},
	}
}

func New() *timerService {
	return &timerService{
		actions:   []types.Action{},
		reactions: []types.Reaction{},
	}
}
