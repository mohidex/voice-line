package database

import "gorm.io/gorm"

type Database interface {
	Transaction(txFunc func(tx *gorm.DB) error) error
	GetDB() *gorm.DB
	Close() error
}
