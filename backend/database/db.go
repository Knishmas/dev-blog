package database 

import (
	"database/sql"
	"log"
	"os"
	//"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB(){
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./blog.DB"
	}

	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Open database error", err)
	}
	
	RunMigrations(DB)
}

func CloseDB(){
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Fatal("Failed to close database", err)
		}
	} else {
		log.Println("Database already closed")
	}

}
