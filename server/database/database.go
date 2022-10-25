package database

import (
	"fmt"
	"os"

	c "Area/config"
	"Area/database/models"
	"Area/lib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var User userController

var Trigger triggerController

var Action actionController

var Reaction reactionController

func New(config *c.Config) *gorm.DB {
	var err error
	var postgresConfig postgres.Config
	var dsn string

	if os.Getenv("DATABASE_URL") != "" {
		dsn = os.Getenv("DATABASE_URL")
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			config.Database.User, config.Database.Password, config.Database.Address, config.Database.Port, config.Database.Name)
	}

	postgresConfig = postgres.Config{
		DSN: dsn,
	}

	db, err = gorm.Open(postgres.New(postgresConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	lib.CheckError(err)

	return db
}

func Seed(db *gorm.DB) {
	var actions = []models.Action{
		{Type: models.NoneAction, Event: models.NoneEvent},
		{Type: models.GmailAction, Event: models.ReceiveEvent},
		{Type: models.GmailAction, Event: models.SendEvent},
	}
	var reactions = []models.Reaction{
		{Type: models.NoneReaction, Action: models.NoneReactionAction},
		{Type: models.DiscordReaction, Action: models.SendReaction},
	}
	db.Exec("DELETE FROM actions")
	db.Exec("DELETE FROM reactions")
	db.Create(actions)
	db.Create(reactions)
}
