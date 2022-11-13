package notion

import (
	"Area/database/models"
	"Area/lib"
	"context"
	"strings"
	"time"

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
	var location, _ = time.LoadLocation("CET")
	var timestamp = storedData.Timestamp.In(location).Format("January 2, 2006") + " at " + storedData.Timestamp.Local().Format("15:04:05")

	basicBlockDescription := notionapi.BasicBlock{
		Object: "block",
		Type:   notionapi.BlockTypeParagraph,
	}

	basicBlockAuthor := notionapi.BasicBlock{
		Object: "block",
		Type:   notionapi.BlockTypeParagraph,
	}

	basicBlockTimestamp := notionapi.BasicBlock{
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
		Text:        &notionapi.Text{Content: storedData.Description},
		Annotations: &notionapi.Annotations{Bold: true, Italic: false, Strikethrough: false, Underline: false, Code: false, Color: "default"},
		Mention:     nil,
		Equation:    nil,
		PlainText:   storedData.Description,
	}

	author := notionapi.RichText{
		Text:        &notionapi.Text{Content: "By :" + storedData.Author},
		Annotations: &notionapi.Annotations{Bold: false, Italic: true, Strikethrough: false, Underline: false, Code: false, Color: "default"},
		Mention:     nil,
		Equation:    nil,
		PlainText:   storedData.Author,
	}

	timestampText := notionapi.RichText{
		Text:        &notionapi.Text{Content: "At :" + timestamp},
		Annotations: &notionapi.Annotations{Bold: false, Italic: false, Strikethrough: false, Underline: false, Code: false, Color: "default"},
		Mention:     nil,
		Equation:    nil,
		PlainText:   timestamp,
	}

	paragraphDescription := notionapi.Paragraph{
		RichText: []notionapi.RichText{description},
	}

	paragraphAuthor := notionapi.Paragraph{
		RichText: []notionapi.RichText{author},
	}

	paragraphTimestamp := notionapi.Paragraph{
		RichText: []notionapi.RichText{timestampText},
	}

	page.Parent = notionapi.Parent{
		Type:   notionapi.ParentTypePageID,
		PageID: notionapi.PageID(pageId),
	}
	page.Children = []notionapi.Block{
		notionapi.ParagraphBlock{
			BasicBlock: basicBlockDescription,
			Paragraph:  paragraphDescription,
		},
		notionapi.ParagraphBlock{
			BasicBlock: basicBlockAuthor,
			Paragraph:  paragraphAuthor,
		},
		notionapi.ParagraphBlock{
			BasicBlock: basicBlockTimestamp,
			Paragraph:  paragraphTimestamp,
		},
	}
	page.Properties = notionapi.Properties{"title": notionapi.TitleProperty{Title: []notionapi.RichText{title}}}
	_, err := client.Page.Create(context.Background(), &page)
	lib.LogError(err)
}
