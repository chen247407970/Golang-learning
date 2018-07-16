package main

import (
	db "github.com/chen247407970/simpleWebsite/database"
	router "github.com/chen247407970/simpleWebsite/routers"
)

func main() {
	defer db.SqlDB.Close()
	router := router.InitRouter()
	router.Run(":33323")
}