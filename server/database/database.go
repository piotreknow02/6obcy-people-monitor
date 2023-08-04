package database

import (
	"github.com/glebarez/sqlite"
	"github.com/piotreknow02/6obcy-people-monitor/server/config"
	"github.com/piotreknow02/6obcy-people-monitor/server/models"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Conf.DbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&models.Log{},
	)

	return db, nil
}
