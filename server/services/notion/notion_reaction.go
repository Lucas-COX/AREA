package notion

import (
	"Area/database/models"
	"Area/lib"
	"context"
	"strings"

	"github.com/jomei/notionapi"
)

func comment(client *notionapi.Client, storedData models.TriggerData, trigger *models.Trigger) {
	var comment notionapi.CommentCreateRequest
	pageUrl := strings.Split(storedData.ReactionData, "-")
	pageId := pageUrl[len(pageUrl)-1]

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

func create_page(client *notionapi.Client, storedData models.TriggerData, trigger *models.Trigger) {
	var page notionapi.PageCreateRequest
	pageUrl := strings.Split(storedData.ReactionData, "-")
	pageId := pageUrl[len(pageUrl)-1]

	basicBlock := notionapi.BasicBlock{
		Object: "block",
		Type:   notionapi.BlockTypeParagraph,
	}

	title := notionapi.RichText{
		Text:        &notionapi.Text{Content: storedData.Title},
		Mention:     nil,
		Equation:    nil,
		Annotations: &notionapi.Annotations{Bold: true, Italic: false, Strikethrough: false, Underline: false, Code: false, Color: "default"},
		PlainText:   storedData.Title,
	}

	description := notionapi.RichText{
		Text:      &notionapi.Text{Content: storedData.Description},
		Mention:   nil,
		Equation:  nil,
		PlainText: storedData.Description,
	}

	paragraph := notionapi.Paragraph{
		RichText: []notionapi.RichText{title, description},
	}

	page.Parent = notionapi.Parent{
		Type:   notionapi.ParentTypePageID,
		PageID: notionapi.PageID(pageId),
	}
	page.Children = []notionapi.Block{notionapi.ParagraphBlock{
		BasicBlock: basicBlock,
		Paragraph:  paragraph,
	}}
	page.Properties = notionapi.Properties{"title": notionapi.TitleProperty{Title: []notionapi.RichText{title}}}
	_, err := client.Page.Create(context.Background(), &page)
	lib.LogError(err)
}
