package database

import (
	"fmt"
	"os"

	c "Area/config"
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
		dsn = fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=%s&loc=Local",
			config.Database.User, config.Database.Password, config.Database.Protocol,
			config.Database.Address, config.Database.Name, config.Database.ParseTime)
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
