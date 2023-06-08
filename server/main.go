package main

import (
	"fmt"
	"log"
	"main/config"
	"main/db"
	"main/repositories"

	"github.com/andrefsilveira1/LoadEnv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Server started!")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		config.Config()
		db, erro := db.Connect()
		if erro != nil {
			panic(erro)
		}
		defer db.Close()

		repositorie := repositories.NewRepository(db)
		datas, err := repositorie.FindAll("")
		if err != nil {
			panic(err)
		}
		return c.JSON(datas)
	})

	// app.Get("/api/todos", func(c *fiber.Ctx) error {
	// 	return c.JSON(todos)
	// })

	// app.Post("/api/todos", func(c *fiber.Ctx) error {
	// 	todo := &Todo{}

	// 	if err := c.BodyParser(todo); err != nil {
	// 		return err
	// 	}

	// 	todo.ID = len(todos) + 1
	// 	todos = append(todos, *todo)

	// 	return c.JSON(todo)
	// })

	// app.Patch("/api/todos/:id/completed", func(c *fiber.Ctx) error {
	// 	id, err := c.ParamsInt("id")
	// 	if err != nil {
	// 		return c.Status(401).SendString("Invalid ID:")
	// 	}

	// 	for i, t := range todos {
	// 		if t.ID == id {
	// 			todos[i].Completed = true
	// 			break
	// 		}
	// 	}
	// 	return c.JSON(todos)
	// })

	// app.Patch("/api/todos/:id/uncompleted", func(c *fiber.Ctx) error {
	// 	id, err := c.ParamsInt("id")
	// 	if err != nil {
	// 		return c.Status(401).SendString("INVALID ID")
	// 	}
	// 	for i, t := range todos {
	// 		if t.ID == id {
	// 			todos[i].Completed = false
	// 			break
	// 		}
	// 	}
	// 	return c.JSON(todos)
	// })

	port := LoadEnv.LoadEnv("API_PORT")
	fmt.Println("Listening on Port:", port)
	log.Fatal(app.Listen(":" + port))
}
