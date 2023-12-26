package controller

import (
	"database/sql"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func SearchUserController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	name := r.URL.Query().Get("name") // To be filled
	if name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := usecase.SearchUser(db, name)
	if err != nil {
		log.Printf("fail: usecase.SearchUser, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
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
}
