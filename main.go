package main

import (
	"./db"
	logger "./log"
)

func main() {
	db.HelloDB()
	db.HelloDataB()
	logger.HelloLog()
}
