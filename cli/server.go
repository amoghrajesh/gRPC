package main

import "gRPC/db"

func main() {
	serverStub()
}

func serverStub() {
	db.Run()
}
