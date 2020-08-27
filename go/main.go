package main

import (
	"flag"
	"github.com/HackU_2020_vol.2_sysken/go/server"
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