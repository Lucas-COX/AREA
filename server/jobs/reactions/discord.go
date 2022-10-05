package reactions

import "Area/database/models"

func React(reaction models.Reaction) {
	switch reaction.Type {
	case "send":
		//TODO send discord message
	}
}
