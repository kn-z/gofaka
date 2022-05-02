package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Cid  int    `gorm:"type:int;primary_key;not null" json:"cid"`
	Name string `gorm:"type:varchar(20);not null" json:"username"`
}
