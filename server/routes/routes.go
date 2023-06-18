package routes

import (
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("api/data/:id", controllers.GetData)
	app.Post("/api/data/create", controllers.CreateData)
	app.Patch("/api/data/complete/:id", controllers.CompleteTask)
	app.Patch("/api/data/undo/:id", controllers.UndoTask)
	app.Delete("/api/data/delete/:id", controllers.DeleteTask)

}
