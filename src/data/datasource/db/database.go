package db

import "gorm.io/gorm"

type DatabaseInterface interface {
	CreateTable()
	Insert() error
	Find()
}

func NewDatabase(db *gorm.DB) database {
	return database{
		dbClient: db,
	}
}

type database struct {
	dbClient *gorm.DB
}
