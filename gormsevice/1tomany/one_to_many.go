package main

import (
	"fmt"
	"go-gorm/gormsevice/connection"
	"gorm.io/gorm"
)

type Dog struct {
	ID     int
	Name   string
	GirlID int
}
type Girl struct {
	ID   int
	Name string
	Dogs []Dog
}

func main() {
	db := connection.DB()
	//db.AutoMigrate(&Girl{}, &Dog{})
	//d1 := Dog{
	//	ID:   1,
	//	Name: "gou1",
	//}
	//d2 := Dog{
	//	ID:   2,
	//	Name: "gou2",
	//}
	//g := Girl{
	//	ID:   0,
	//	Name: "",
	//	Dogs: []Dog{d1, d2},
	//}
	//db.Create(&g)
	var g Girl
	var g2 Girl
	var g3 Girl
	db.Preload("Dogs").Find(&g)
	db.Preload("Dogs", "name = ?", "gou1").Find(&g2)
	fmt.Println(g)
	fmt.Println(g2)
	db.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", "gou2")
	}).Find(&g3)
	fmt.Println(g3)
}
