package controllers

import (
	"main/db"
	"main/models"
	"main/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateData(c *fiber.Ctx) error {
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	localData := models.Data{
		Title:       body["title"],
		Completed:   body["completed"],
		Description: body["description"],
	}

	db, erro := db.Connect()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repositorie := repositories.NewRepository(db)
	result, err := repositorie.CreateData(localData)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(result)
}

func GetData(c *fiber.Ctx) error {
	db, erro := db.Connect()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repositorie := repositories.NewRepository(db)
	result, err := repositorie.FindAll()
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(result)
}

func CompleteTask(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	db, erro := db.Connect()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repositorie := repositories.NewRepository(db)
	result, err := repositorie.CompleteTask(id)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(result)

}

func UndoTask(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	db, erro := db.Connect()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repositorie := repositories.NewRepository(db)
	result, err := repositorie.UndoTask(id)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(result)

}
