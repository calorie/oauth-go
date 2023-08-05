package repository

import (
	"os"

	"github.com/calorie/oauth-go/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.AuthorizationCode{}, &domain.CodeChallenge{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
