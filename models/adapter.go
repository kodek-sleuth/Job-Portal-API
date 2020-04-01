package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/lib/pq"
	"github.com/satori/go.uuid"
	"time"

	//"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "kevina52"
	dbname   = "testing"
)

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

func SQLConnection() (*gorm.DB, interface{}) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	// Migrate the schema
	db.AutoMigrate(&Job{}, &Users{})
	return db, nil
}
