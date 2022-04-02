package main

import (
	"fmt"

	"github.com/fullaheadcoco/gin-todos/Config"
	"github.com/fullaheadcoco/gin-todos/Models"
	"github.com/fullaheadcoco/gin-todos/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// Creating a connection to the database
	Config.DB, err = gorm.Open("mysql", Config.DbURL((Config.BuildDBConfig())))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()

	// run the migration: todo struct
	Config.DB.AutoMigrate(&Models.Todo{})

	// setup routes
	r := Routes.SetupRouter()

	// running
	r.Run()
}
