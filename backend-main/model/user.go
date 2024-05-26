package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex"`
	Password       string
	Nexuses        []*Nexus `gorm:"many2many:user_nexuses;"`
	PushoverKey    string
	TelegramUserID string
}

func (u *User) AfterFind(_ *gorm.DB) (err error) {
	if u.Nexuses == nil {
		u.Nexuses = make([]*Nexus, 0)
	}
	return
}

func (u *User) AfterCreate(_ *gorm.DB) (err error) {
	if u.Nexuses == nil {
		u.Nexuses = make([]*Nexus, 0)
	}
	return
}
