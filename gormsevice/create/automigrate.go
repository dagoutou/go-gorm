package main

import (
	"encoding/json"
	"go-gorm/gormsevice/connection"
	"gorm.io/datatypes"
)

type Product struct {
	Id     int            `json:"id" gorm:"primary key"`
	Name   string         `json:"name" gorm:"type:varchar(25);not null"`
	Script datatypes.JSON `json:"script" gorm:"json;not null"`
}

func main() {
	var (
		mp = make(map[string]string)
		pr Product
	)
	mp["age"] = "222"
	mp["sql"] = "select * from aaaa"
	pr.Name = "gouzi"
	marshal, err := json.Marshal(&mp)
	if err != nil {
		return
	}
	pr.Script = marshal
	db := connection.DB()
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}
	db.Create(&pr)
}
