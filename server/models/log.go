package models

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Count uint16 `json:"count"`
}

