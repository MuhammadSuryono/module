package database

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

var GormDb *gorm.DB

func (dbConfig DBConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

// Init database
func Init(config DBConfig)  {
	db, err := gorm.Open("mysql", config.GetConnectionString())
	if err != nil {
		fmt.Println("error ", err)
	}
	GormDb = db
	pingTicker := time.NewTicker(15 * time.Second)
	pingDone := make(chan bool)
	go func() {
		for {
			select {
			case <-pingDone:
				return
			case <-pingTicker.C:
				b := pingDb(db.DB())
				if !b {
					pingDone <- true
				}
			}
		}
	}()
}

// Ping to DB
func pingDb(db *sql.DB) bool {
	er := db.Ping()
	if er != nil {
		log.Print(string("\033[31m"), "mysql error ping", er)
		return false
	} else {
		log.Print(string("\033[32m"), "mysql success ping")
		return true
	}
}

// Create new Connection
func CreateConnection() {

	port, err := strconv.Atoi(os.Getenv("DATABASE_POST"))

	if err != nil {
		fmt.Println(string("\033[33m"), err)
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     port,
		User:     os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBName:   os.Getenv("DATABASE_NAME"),
	}

	Init(dbConfig)
}