package dao

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// ①-1
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// ①-2
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlUserPwd, mysqlDatabase))
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func SearchUsers() {
	rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
}

func InsertUsers() {
	_, err := db.Exec("INSERT INTO user (id, name, age) VALUES (?,?, ?)", id.String(), requestData.Name, requestData.Age)
}
