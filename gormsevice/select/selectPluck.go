package main

import (
	"go-gorm/gormsevice/connection"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	Id     int            `json:"id" gorm:"primary key"`
	Name   string         `json:"name" gorm:"type:varchar(25);not null"`
	Script datatypes.JSON `json:"script" gorm:"json;not null"`
}
type user struct {
	Name string `json:"name" gorm:"type:varchar(20)"`
	gorm.Model
}

func main() {
	//var pl []string

	db := connection.DB()
	//db.Model(&Product{}).Pluck("name", &pl)
	//fmt.Println(pl)
	//db.Model(&Product{}).Where(1).Update("name", "狗子")
	//db.Where(1).Find(&pl1)
	//
	//db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Product{})
	//db.Exec("DELETE FROM products")
	//fmt.Println(pl1)
	db.Where(1).Delete(&user{})
}
