package main

import (
	"fmt"
	"log"
	"main/routes"

	"github.com/andrefsilveira1/LoadEnv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Server started!")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	routes.Setup(app)
	port := LoadEnv.LoadEnv("API_PORT")
	fmt.Println("Listening on Port:", port)
	log.Fatal(app.Listen(":" + port))
}
