package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
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

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			log.Printf("fail: db.Begin, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()
		// ②-1
		name := r.URL.Query().Get("name") // To be filled
		if name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// ②-2
		rows, err := tx.Query("SELECT id, name, age FROM user WHERE name = ?", name)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// ②-3
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}
		// トランザクションをコミット
		if err := tx.Commit(); err != nil {
			log.Printf("fail: tx.Commit, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// ②-4
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	case http.MethodPost:
		//トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			log.Printf("fail: db.Begin, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		// POSTメソッドの処理
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		id := ulid.MustNew(ulid.Timestamp(t), entropy)

		var requestData struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		// HTTPリクエストボディからJSONデータを読み取る
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestData); err != nil {
			log.Printf("fail: json.Decode, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if requestData.Name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(requestData.Name) > 50 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if requestData.Age < 20 || requestData.Age > 80 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// データベースにINSERT
		_, err = tx.Exec("INSERT INTO user (id, name, age) VALUES (?,?, ?)", id.String(), requestData.Name, requestData.Age)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			if err := tx.Rollback(); err != nil {
				log.Printf("fail: tx.Rollback, %v\n", err)
			}
			return
		}

		if err := tx.Commit(); err != nil {
			log.Printf("fail: tx.Commit, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 成功した場合のレスポンス
		w.WriteHeader(http.StatusOK)
		response := map[string]string{"id": id.String()}
		bytes, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
