package notion

import (
	"Area/database/models"
	"context"

	"github.com/jomei/notionapi"
)

func comment(client *notionapi.Client, pageId string, storedData models.TriggerData, trigger *models.Trigger) {
	var comment notionapi.CommentCreateRequest

	comment.Parent = notionapi.Parent{
		Type:   notionapi.ParentTypePageID,
		PageID: notionapi.PageID(pageId),
	}

	comment.RichText = []notionapi.RichText{
		{
			Type: notionapi.ObjectTypeText,
			Text: &notionapi.Text{Content: storedData.Description},
		},
	}
	client.Comment.Create(context.Background(), &comment)
}
