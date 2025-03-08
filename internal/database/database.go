package database

import (
	"github.com/skhanal5/prizepicks/internal/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	session *gorm.DB //potential smell
}

func New(config common.Config) (Database, error) { //TODO: Revisit parameter
	// use env var here
	session, err := gorm.Open(sqlite.Open(config.DBFile), &gorm.Config{})
	if err != nil {
		panic("Failed get a connection with the database") // todo: error handling
	}
	return Database{session: session}, nil
}
