package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //переменная для работы с БД

func InitDB() (*gorm.DB, error) {
	// Функция инициализации БД
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable" //Data Source Name
	var err error

	// Подключаемся к БД, если не удалось выдаем fatal
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v ", err)
	}

	return db, nil
}
