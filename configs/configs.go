package configs

import (
	"github.com/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func ConnectDatabase() error {
	var err error
	DataBase, err = gorm.Open(sqlite.Open("credManger.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DataBase.AutoMigrate(&models.UserCredentials{}, &models.CredentialsIDs{})

	return nil
}
