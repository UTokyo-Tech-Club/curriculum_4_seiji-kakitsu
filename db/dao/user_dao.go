package dao

import (
	"database/sql"
	"db/model"
	"fmt"
	"github.com/oklog/ulid/v2"
	"log"
	"math/rand"
	"os"
	"time"
)

var db *sql.DB

func InitDB() {
	// ①-1
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

func HandleTransaction(tx *sql.Tx, err error) {
	// 関数が終了する際にトランザクションをコミットまたはロールバック
	if r := recover(); r != nil {
		// パニックが発生した場合はロールバック
		tx.Rollback()
		log.Printf("fail: Transaction rolled back due to panic: %v\n", r)
	} else if err != nil {
		// エラーが発生した場合はロールバック
		tx.Rollback()
		log.Printf("fail: Transaction rolled back due to error: %v\n", err)
	} else {
		// 成功した場合はトランザクションをコミット
		if err := tx.Commit(); err != nil {
			log.Printf("fail: tx.Commit, %v\n", err)
			return
		}
	}
}

func GetUserByName(name string) (*sql.Rows, error) {
	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return nil, err
	}
	defer HandleTransaction(tx, err)

	rows, err := tx.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		log.Printf("fail: tx.Query, %v\n", err)
		return nil, err
	}
	return rows, err
}

func CreateUser() (string, error) {
	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return "", err
	}
	defer HandleTransaction(tx, err)

	// IDの生成
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	_, err = tx.Exec("INSERT INTO user (id, name, age) VALUES (?,?, ?)", id, model.RequestData.Name, model.RequestData.Age)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}

	return id, nil
}
