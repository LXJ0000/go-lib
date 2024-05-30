package orm

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewOrmDatabase() Database {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go-backend?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(); err != nil {
		log.Fatal(err)
	}
	database := NewDatabase(db)

	return database
}
