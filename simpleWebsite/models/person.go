package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"simpleWebsite/setting"
)

var db *gorm.DB

type Person struct {
	Id   int    `gorm:"primary key" json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type RegisterRsp struct {
	Success bool `json:"success"`
	Person
}

func init() {
	var dbType, dbName, user, password, host string

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&Person{})
}

func closeDB() {
	defer db.Close()
}

func (p *Person) AddPerson() bool {
	/*
		rs, err := db.SqlDB.Exec("insert into persons(id, name) values (?, ?)", p.Id, p.Name)
		if err != nil {
			fmt.Println("exec failed,", err)
			return false
		}
		id, err := rs.LastInsertId()
		if err != nil {
			fmt.Println("exec failed,", err)
			return false
		}
		// db.SqlDB.Query("select sleep(10)")
	*/

	err := db.Create(p)
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, "insert failed:", err)
		return false
	}
	fmt.Println("InsertedPerson={\"id\":", p.Id, ", \"name\": \"", p.Name, "\"}")
	if db.NewRecord(*p) {
		fmt.Fprintf(gin.DefaultWriter, "insert persons(id, name) values (%d, %s) failed!", p.Id, p.Name)
		return false
	}

	fmt.Fprintln(gin.DefaultWriter, "insert success:", p.Id)
	return true
}

func (p *Person) QueryPerson() bool {
	/*
		rows, err := db.SqlDB.Query("SELECT name FROM persons WHERE id=?", p.Id)
		defer rows.Close()
		if err != nil {
			return false
		}
	*/
	isEmpty, isCorrect := true, false
	/*
		for rows.Next() {
			var rowName string
			isEmpty = false
			if err := rows.Scan(&rowName); err == nil && p.Name == rowName {
				isCorrect = true
			}
		}
	*/

	var dbPerson Person
	db.Where("id = ?", p.Id).First(&dbPerson)
	/*
		if dbPerson != nil {
			isEmpty = false
			if dbPerson.Name == p.Name {
				isCorrect = true
			}
		}
	*/

	fmt.Fprintln(gin.DefaultWriter, "dbPerson={\"id\":", dbPerson.Id, ", \"name\": \"", dbPerson.Name, "\"}")
	if dbPerson.Id != 0 {
		isEmpty = false
		if dbPerson.Name == p.Name {
			isCorrect = true
		}
	}

	if isEmpty {
		reqMap := make(map[string]interface{})
		reqMap["id"] = p.Id
		reqMap["name"] = p.Name
		bytesData, err := json.Marshal(reqMap)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		reader := bytes.NewReader(bytesData)
		url := "http://0.0.0.0:33323/register"
		request, err := http.NewRequest("POST", url, reader)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
		client := http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		var rspJson RegisterRsp
		err = json.Unmarshal(respBytes, &rspJson)
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, "Json unmarshal failed,", err.Error())
		}

		if rspJson.Success {
			return true
		}
		return false
	}
	if isCorrect {
		fmt.Fprintln(gin.DefaultWriter, "Get data, id: ", p.Id, " name: ", p.Name)
		return true
	}
	return false
}
