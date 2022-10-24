package main

import (
	"go-gorm/gormsevice/connection"
)

type Dog struct {
	ID     int
	Name   string
	GirlID int
	Girl   Girl
}
type Girl struct {
	ID   int
	Name string
}

func main() {
	db := connection.DB()
	gril := Girl{
		ID:   1,
		Name: "girl1",
	}
	dog1 := Dog{
		ID: 1,
	}
	//db.AutoMigrate(&Dog{})
	//db.Create(&gril)
	//db.Create(&dog1)
	db.Model(&dog1).Association("Girl").Append(&gril)

}
