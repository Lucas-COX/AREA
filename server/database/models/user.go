package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"not null;uniqueIndex"`
	Password    string `gorm:"not null"`
	Triggers    []Trigger
	GoogleToken string
}

func (u *User) TableName() string { return "users" }
