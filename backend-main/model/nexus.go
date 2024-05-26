package model

import "gorm.io/gorm"

type Nexus struct {
	gorm.Model
	MacAddress string      `gorm:"uniqueIndex" json:"macAddress"`
	PairSecret string      `json:"pairSecret"`
	Users      []*User     `gorm:"many2many:user_nexuses;" json:"users"`
	Config     NexusConfig `gorm:"serializer:json" json:"config"`
	NickName   string      `json:"nickName"`
}

type NexusConfig struct {
	Sensitivity float64        `json:"sensitivity"`
	TitanW      []TitanWConfig `json:"titanW"`
}

type TitanWConfig struct {
	UUID     string `json:"uuid"`
	NickName string `json:"nickName"`
	Enabled  bool   `json:"enabled"`
}

func (n *Nexus) AfterFind(_ *gorm.DB) (err error) {
	if n.Users == nil {
		n.Users = make([]*User, 0)
	}
	if n.Config.TitanW == nil {
		n.Config.TitanW = make([]TitanWConfig, 0)
	}
	return
}

func (n *Nexus) AfterCreate(_ *gorm.DB) (err error) {
	if n.Users == nil {
		n.Users = make([]*User, 0)
	}
	if n.Config.TitanW == nil {
		n.Config.TitanW = make([]TitanWConfig, 0)
	}
	return
}
