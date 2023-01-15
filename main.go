package main

import (
	"fmt"

	"github.com/syed-ashraf123/go-crud/database"
	"github.com/syed-ashraf123/go-crud/models"
	"github.com/syed-ashraf123/go-crud/router"
)

func main() {
	database.Connect()
	models.Migrate()
	fmt.Println("Successfully Connected")

	r := router.SetupRouter()
	r.Run()
}
