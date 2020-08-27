package model

import (
	"database/sql"
	"log"
)

type User struct {
	User_id      int
	Name         string
	Password     string
	Token        string
	Feed_num     int
	Rgister_date string
}

func SelectUsersData(user_id int) (*User, error) {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM users WHERE user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	user := User{}
	err = rows.Scan(&user.User_id, &user.Name, &user.Name, &user.Password, &user.Token, &user.Feed_num, &user.Rgister_date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
