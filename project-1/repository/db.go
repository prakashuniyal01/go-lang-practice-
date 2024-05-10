package repository

import (
	"fmt"
	"log"
	"trademarkia/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPsqlDB(cfg *common.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DbHost, cfg.DbUser, cfg.DbName, cfg.DbPort, cfg.DbPassword)
	Db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	Db.AutoMigrate(UserData{})
	return Db, nil
}

func SeedDatabase() error {
	users := []UserData{
		{FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 30, Married: true},
		{FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Age: 25, Married: false},
	}

	for _, user := range users {
		if err := DB.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
