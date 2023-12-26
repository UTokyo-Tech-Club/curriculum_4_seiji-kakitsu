package controller

import (
	"database/sql"
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func RegisterUserController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model.RequestData); err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := usecase.RegisterUser(db)
	if err != nil {
		log.Printf("fail: %n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 成功した場合のレスポンス
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"id": id}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
