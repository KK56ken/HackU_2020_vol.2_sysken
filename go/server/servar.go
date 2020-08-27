package server

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Users struct {
	user_id       int
	name          string
	password      string
	token         string
	feed_num      int
	register_date string
}

func Serve(addr string)  {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	// http.GET("/poke", poke)
	http.ListenAndServe(":3001", nil)

}






