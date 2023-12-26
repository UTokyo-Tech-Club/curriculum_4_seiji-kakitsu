package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"log"
)

func SearchUser(db *sql.DB, name string) ([]model.UserResForHTTPGet, error) {
	users, err := dao.GetUserByName(db, name)
	if err != nil {
		log.Printf("fail: %v\n", err)
		return nil, err
	}
	return users, nil
}
