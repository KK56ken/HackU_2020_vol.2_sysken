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
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	// http.GET("/poke", poke)
	http.ListenAndServe(":3001", nil)

	db, err := sql.Open("mysql", "root:@/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db=", db)

	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM users") //
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("rows=", rows)

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	fmt.Println("values=", values)
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
