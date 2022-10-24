package main

import (
	"fmt"
	"go-gorm/gormsevice/connection"
)

type Dog struct {
	ID     int
	Name   string
	GirlID int
}
type Girl struct {
	ID   int
	Name string
	Dog  Dog
}

func main() {
	db := connection.DB()

	//db.AutoMigrate(&Girl{}, &Dog{})
	//
	//dog := Dog{
	//	Name: "dog1",
	//}
	//gril2 := Girl{
	//	ID:   2,
	//	Name: "girl2",
	//	Dog:  dog,
	//}
	//
	//db.Create(&gril2)
	var gril Girl
	db.Preload("Dog").First(&gril, 2)
	fmt.Println(gril)

}
