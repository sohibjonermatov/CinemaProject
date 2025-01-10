package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"secondProject/models"
	"secondProject/pkg/settings"
)

var db *gorm.DB

func Setup() {
	var err error

	dbConfig := settings.Config.DB

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("db.Setup err:%w", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)

	authoMigration()
	log.Println("DB successfully connected")
}

func CloseDB() {
	sqlDB, err := db.DB()
	sqlDB.Close()
	if err != nil {
		log.Fatalf("Error on closing the DB: %w", err)
	}
}

func GetDb() *gorm.DB {
	return db
}

func authoMigration() {
	for _, model := range []interface{}{
		(*models.Movie)(nil),
	} {
		dbSilent := db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
		err := dbSilent.AutoMigrate(model)
		if err != nil {
			log.Fatalf("create model %s : %s", model, err)
		}
	}
}
