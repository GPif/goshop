package main

import (
	"fmt"
	"goshop/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = config.InitDB("goshop.db")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Starting...")
}
