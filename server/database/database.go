package database

import (
	"fmt"

	c "Area/config"
	"Area/lib"

	"gorm.io/driver/mysql"
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
	mysqlConfig := mysql.Config{
		DSN: fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=%s&loc=Local",
			config.Database.User, config.Database.Password, config.Database.Protocol,
			config.Database.Address, config.Database.Name, config.Database.ParseTime),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	lib.CheckError(err)

	return db
}
