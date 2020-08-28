package main

import (
	"flag"
	"hacku_vol2/server/server"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Serve(addr)
}
