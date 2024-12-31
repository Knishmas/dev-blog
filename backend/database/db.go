package database 

import (
	"database/sql"
	"log"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(){
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./blog.db"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Open database error", err)
	}
	
	RunMigrations()
}