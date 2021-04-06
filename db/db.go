package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"github.com/go-sql-driver/mysql"
	"github.com/r4lly99/echo-snack/model"
)

func New() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 os.Getenv("DB_HOST" + ":" + "DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.DB().SetConnMaxLifetime(time.Minute * 3)
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@/dbname")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../realworld_test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Follow{},
		&model.Snack{},
	)
}
