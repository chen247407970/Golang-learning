package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
	"github.com/chen247407970/simpleWebsite/models"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	id_str := c.PostForm("id")
	name := c.PostForm("name")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}

	ps := models.Person{
		Id:		id,
		Name:	name,
	}

	ra := ps.QueryPerson()
	if ra {
		c.JSON(http.StatusOK, gin.H {
			"success":	ra,
			"id":		id,
			"name":		name,
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"success":	ra,
		})
	}
}

func RegisterHandler(c *gin.Context) {
	id_str := c.PostForm("id")
	name := c.PostForm("name")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}

	ps := models.Person{
		Id:		id,
		Name:	name,
	}

	ra := ps.AddPerson()
	if ra {
		c.JSON(http.StatusOK, gin.H {
			"success":	ra,
			"id":		id,
			"name":		name,
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"success":	ra,
		})
	}
}