package services

import (
	"Area/services/discord"
	"Area/services/github"
	"Area/services/google"
	"Area/services/microsoft"
	"Area/services/notion"
)

var Google = google.New()
var Microsoft = microsoft.New()
var Github = github.New()
var Discord = discord.New()
var Notion = notion.New()
