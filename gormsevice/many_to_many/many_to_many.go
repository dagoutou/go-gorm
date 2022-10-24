package main

import (
	"fmt"
	"go-gorm/gormsevice/connection"
)

type Dog struct {
	ID    int
	Name  string
	Girls []Girl `gorm:"many2many:dog_girl_relation"`
}
type Girl struct {
	ID   int
	Name string
	Dogs []Dog `gorm:"many2many:dog_girl_relation"`
}

func main() {
	db := connection.DB()
	//db.AutoMigrate(&Dog{}, &Girl{})
	//d1 := Dog{
	//	ID:   2,
	//	Name: "gou3",
	//}
	//d2 := Dog{
	//	ID:   1,
	//	Name: "gou4",
	//}
	//g1 := Girl{
	//	ID:   3,
	//	Name: "girl1",
	//	Dogs: []Dog{d1, d2},
	//}
	//var d = Dog{
	//	ID: 2,
	//}

	var g []Girl
	//g1 := Girl{
	//	Name: "girl1",
	//	Dogs: []Dog{d1, d2},
	//}
	//db.Create(&g1)
	//db.Preload("Dogs", func(ctx *gorm.DB) *gorm.DB {
	//	return ctx.Where("name = ?", "gou3")
	//}).Association("Girls").Find(&g)
	db.Debug().Preload("Dogs", "ID = ?", "1").Find(&g)
	fmt.Println(g)

}
