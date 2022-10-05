package actions

import "Area/database/models"

func CheckGmailAction(action models.Action) bool {
	switch action.Event {
	case "receive":

		return true
	}
	return false
}
