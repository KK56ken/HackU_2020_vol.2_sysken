package model

import "database/sql"

func Db_access() []string {

	var s []string

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err.Error())
		}
		s = append(s, name)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	return s
}
