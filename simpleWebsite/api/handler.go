package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simpleWebsite/models"
)

func LoginHandler(c *gin.Context) {
	var person models.Person
	err := c.Bind(&person)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"errcode":     400,
			"description": "Post Data Error",
		})
		return
	}

	res := person.QueryPerson()
	if res {
		c.JSON(http.StatusOK, gin.H{
			"success": res,
			"id":      person.Id,
			"name":    person.Name,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": res,
		})
	}
}

func RegisterHandler(c *gin.Context) {
	var person models.Person
	err := c.Bind(&person)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"errcode":     400,
			"description": "Post Data Error",
		})
		return
	}

	res := person.AddPerson()
	if res {
		c.JSON(http.StatusOK, gin.H{
			"success": res,
			"id":      person.Id,
			"name":    person.Name,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": res,
		})
	}
}
