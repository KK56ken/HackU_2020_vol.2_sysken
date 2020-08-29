package model

import "database/sql"

type User struct {
	User_id      int
	Name         string
	Password     string
	Token        string
	Feed_num     int
	Rgister_date string
}

func InsertUserr(record *User) error {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	// userテーブルへのレコードの登録を行うSQLを入力する
	stmt, err := db.Prepare("INSERT INTO users (name, password, token , feed_num ) VALUES (?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.Name, record.Password, record.Token, record.Feed_num)
	return err
}
