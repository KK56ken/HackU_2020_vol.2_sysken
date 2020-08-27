package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	user_id       int
	name          string
	password      string
	token         string
	feed_num      int
	register_date string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var values []string

	values = db_access()

	for i := 0; i < 1; i++ {
		fmt.Fprintln(w, values[i])
	}
}
func db_access() []string {

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

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	// http.GET("/poke", poke)
	http.ListenAndServe(":3001", nil)

	// values := make([]sql.RawBytes, len(columns))

	// fmt.Println("values=", values)
	// //  rows.Scan は引数に `[]interface{}`が必要.

	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	// for rows.Next() {
	// 	err = rows.Scan(scanArgs...)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	var value string
	// 	for i, col := range values {
	// 		// Here we can check if the value is nil (NULL value)
	// 		if col == nil {
	// 			value = "NULL"
	// 		} else {
	// 			value = string(col)
	// 		}
	// 		fmt.Println(columns[i], ": ", value)
	// 	}
	// 	fmt.Println("-----------------------------------")
	// }

}
