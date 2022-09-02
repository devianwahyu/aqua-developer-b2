package database

import (
	"fmt"
	"log"

	"github.com/devianwahyu/efishery/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(conf *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", conf.Host, conf.User, conf.Password, conf.DBName, conf.Port, conf.SSLMode, conf.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
		return db, err
	}
	db.AutoMigrate(&model.Customer{})

	return db, nil
}
