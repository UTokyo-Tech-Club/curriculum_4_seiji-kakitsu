package dao

import (
	"database/sql"
	"db/model"
	"github.com/oklog/ulid/v2"
	"log"
	"math/rand"
	"time"
)

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

type UserDAO struct {
	DB *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{DB: db}
}

func GetUserByName(db *sql.DB, name string) ([]model.UserResForHTTPGet, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin(), %v\n", err)
		return nil, err
	}
	defer HandleTransaction(tx, err)
	rows, err := tx.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		log.Printf("fail: tx.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]model.UserResForHTTPGet, 0)
	for rows.Next() {
		var u model.UserResForHTTPGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func CreateUser(db *sql.DB) (string, error) {
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
