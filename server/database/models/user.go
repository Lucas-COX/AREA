package models

import "gorm.io/gorm"

type UserRole string

const (
	MemberRole UserRole = "member"
	AdminRole  UserRole = "admin"
)

type User struct {
	gorm.Model
	Username    string   `gorm:"not null;uniqueIndex"`
	Password    string   `gorm:"not null"`
	Role        UserRole `gorm:"default:member"`
	Triggers    []Trigger
	GoogleToken string
}

func (u *User) TableName() string { return "users" }
