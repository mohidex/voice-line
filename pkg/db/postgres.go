package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB(host, user, password, dbname string, port int) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get raw database connection: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to PostgreSQL successfully.")
	return &PostgresDB{DB: db}, nil
}

func (p *PostgresDB) Transaction(txFunc func(tx *gorm.DB) error) error {
	return p.DB.Transaction(txFunc)
}

func (p *PostgresDB) GetDB() *gorm.DB {
	return p.DB
}

func (p *PostgresDB) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get raw database connection: %w", err)
	}
	return sqlDB.Close()
}
