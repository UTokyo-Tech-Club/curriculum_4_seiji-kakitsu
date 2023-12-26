package usecase

import (
	"database/sql"
	"db/dao"
	"log"
)

func SearchUser(name string) *sql.Rows {
	rows, err := dao.GetUserByName(name)
	if err != nil {
		log.Printf("fail: %v\n", err)
		return nil
	}
	return rows
}
