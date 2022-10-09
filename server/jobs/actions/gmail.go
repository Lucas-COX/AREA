package actions

import "Area/database/models"

func CheckGmailAction(action models.Action, user models.User) bool {
	switch action.Event {
	case "receive":

		return true
	}
	return false
}
