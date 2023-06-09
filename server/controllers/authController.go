package controllers

import (
	"main/db"
	"main/models"
	"main/repositories"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	passwd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: passwd,
	}

	db, erro := db.Connect()
	if erro != nil {
		panic(erro)
	}
	defer db.Close()
	repositorie := repositories.NewRepository(db)

	result, err := repositorie.CreateUser(user)
	if err != nil {
		panic(err)
	}

	return c.JSON(result)
}
