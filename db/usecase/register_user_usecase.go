package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"fmt"
	"log"
)

func Validation() error {
	if model.RequestData.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if len(model.RequestData.Name) > 50 {
		return fmt.Errorf("name is too long")
	}

	if model.RequestData.Age < 20 || model.RequestData.Age > 80 {
		return fmt.Errorf("invalid age")
	}

	return nil
}

func RegisterUser(db *sql.DB) (string, error) {
	if err := Validation(); err != nil {
		log.Printf("fail: Validation failed, %v\n", err)
		return "", err
	}
	id, err := dao.CreateUser(db)
	if err != nil {
		log.Printf("fail: %v\n", err)
		return "", err
	}
	return id, nil
}
