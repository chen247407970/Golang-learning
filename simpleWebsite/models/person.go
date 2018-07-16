package models

import (
	"fmt"
	db "github.com/chen247407970/simpleWebsite/database"
)

type Person struct {
	Id		int		`json:"id" form:"id"`
	Name	string	`json:"name" form:"name"`
}

func (p *Person) AddPerson() bool {
	rs, err := db.SqlDB.Exec("INSERT INTO person(id, name) VALUES (?, ?)", p.Id, p.Name)
	if err != nil {
		return false
	}
	id, err := rs.LastInsertId()
	fmt.Println("Last insterted id:", id)
	if err != nil {
		return false
	}
	return true
}

func (p *Person) QueryPerson() bool {
	rows, err := db.SqlDB.Query("SELECT id, name FROM person WHERE id=? and name=?", p.Id, p.Name)
	defer rows.Close()
	if err != nil {
		return false
	}
	cnt := 0
	for rows.Next() {
		cnt++
	}
	if cnt == 1 {
		fmt.Println("Get data, id: ", p.Id, " name: ", p.Name)
		return true
	}
	return false
}