package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	. "simpleWebsite/api"
)

func InitRouter() *gin.Engine {
	f, err := os.Create("log/gin.log")
	if err != nil {
		fmt.Println(err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.POST("/login", LoginHandler)
	router.POST("/register", RegisterHandler)

	return router
}
