package routers

import (
	"github.com/gin-gonic/gin"
	. "github.com/chen247407970/simpleWebsite/api"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", LoginHandler)
	router.POST("/register", RegisterHandler)

	return router
}