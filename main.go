package main

import (
	"fmt"
	"gayandn/configration"
	"gayandn/controler"
	"gayandn/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	configration.LoadEnvVariables()
	configration.ConnectDB()
	middleware.Protected()

	fmt.Println("first: init")
}

func main() {

	app := fiber.New()
	app.Use(cors.New())

	//handler.InsertUser()
	controler.SetupRoutes(app)
	log.Fatal(app.Listen(":8050"))
	fmt.Println("Hello Gayandn")
}
