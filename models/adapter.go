package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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

func SQLConnection() (*gorm.DB, error) {
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

//// Config structure
//type Config struct {
//	DB *DBConfig
//}
//// DBConfig structure
//type DBConfig struct {
//	Dialect    string
//	Host       string
//	Port       string
//	Username   string
//	Password   string
//	DBName     string
//	TestDBName string
//	SSLMode    string
//}
//// GetConfig function
//func GetConfig() *Config {
//	return &Config{
//		DB: &DBConfig{
//			Dialect:    helpers.GetEnv("DB_DIALECT", ""),
//			Host:       helpers.GetEnv("DB_HOST", ""),
//			Port:       helpers.GetEnv("DB_PORT", ""),
//			Username:   helpers.GetEnv("DB_USERNAME", ""),
//			Password:   helpers.GetEnv("DB_PASSWORD", ""),
//			DBName:     helpers.GetEnv("DB_NAME", ""),
//			TestDBName: helpers.GetEnv("TEST_DB_NAME", ""),
//			SSLMode:    helpers.GetEnv("DB_SSL_MODE", ""),
//		},
//	}
//}
//func dbConnect(environment string) *gorm.DB {
//	dbConfig := GetConfig()
//	db, err := gorm.Open(dbConfig.DB.Dialect, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
//		dbConfig.DB.Host, dbConfig.DB.Port, dbConfig.DB.Username, databaseName, dbConfig.DB.Password, dbConfig.DB.SSLMode))
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	return db
//}
//func Adapter() {
//	db := dbConnect()
//	db.AutoMigrate(
//		User{},
//	)
//	defer db.Close()
//	return
//}
