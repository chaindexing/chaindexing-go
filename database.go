package chaindexing

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB Database
)

type Database struct {
	DB *gorm.DB
}

func Setup() {
	var db *gorm.DB
	var err error

	config := GetConfig()

	database := config.Database.Dbname
	username := config.Database.Username
	password := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port
	sslMode := config.Database.SSlMode

	switch config.Database.Driver {
	case POSTGRES:
		db, err = gorm.Open(
			postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s", host, port, username, database, sslMode, password)),
		)
	}

	if err != nil {
		panic("Chaindexing -> Cannot connect to Provided DB")
	}

	DB = Database{db}
}

// Create inserts a record in the Database
func (d Database) Create(value interface{}) error {
	return d.DB.Create(value).Error
}

// WithTrx returns a transaction
func (d Database) WithTrx() *gorm.DB {
	return d.DB.Begin()
}

// Count counts all records in a table
func (d Database) Count(model interface{}, count *int64) {
	d.DB.Model(model).Count(count)
}

func GetDB() Database {
	return DB
}
