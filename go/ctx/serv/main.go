package main

import (
	"serv/server"
)

func main() {
	srv := server.DefaultServer
	srv.ListenAndServe()
}
