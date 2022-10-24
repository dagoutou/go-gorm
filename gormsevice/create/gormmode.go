package main

import (
	"fmt"
	"go-gorm/gormsevice/connection"
	"gorm.io/gorm"
)

type user struct {
	Name string `json:"name" gorm:"type:varchar(20)"`
	gorm.Model
}

func main() {
	var us []user
	//var a string
	db := connection.DB()
	//db.Create(&us)
	//db.Model(&user{}).Find(&us)
	db.Limit(1).Offset(1).Find(&us)
	fmt.Println(us)

}
