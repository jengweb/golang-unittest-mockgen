package core

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func SetupDatabase() *gorm.DB {
	dbConfig := DBConfig{
		Host:   "127.0.0.1",
		Port:   "3306",
		DBName: "data_test",
		User:   "root",
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	defer sqlDB.Close()
	return db
}
