package main

import (
	_ "simpleWebsite/models"
	router "simpleWebsite/routers"
)

func main() {
	router := router.InitRouter()
	router.Run(":33323")
}
