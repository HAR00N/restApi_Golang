package api

import (
	"fmt"
	"goapi/config"
)

var uid int
var user_id string
var password string

func DbGetDao(query string) {
	database := config.DbConnection()
	defer database.Close()

	rows, err := database.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&uid, &user_id, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println(uid, user_id, password)
	}

	fmt.Println("Connection Success")
}

func MysqlGetDao(query string) {
	database := config.MysqlConnection()
	defer database.Close()

	rows, err := database.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&uid, &user_id, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println(uid, user_id, password)
	}

	fmt.Println("Connection Success")
}
