package main

import (
	"fmt"
	"go-gorm/gormsevice/connection"
)

func main() {
	db := connection.DB()
	fmt.Println(db)
}
