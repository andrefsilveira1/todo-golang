package controllers

import (
	"fmt"
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

	passwd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

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

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	db, erro := db.Connect()
	if erro != nil {
		panic(erro)
	}
	defer db.Close()

	user := models.User{
		Email:    data["email"],
		Password: []byte("$2a$14$unumYyVv69bqWSpzki1UoutvYO4aM2rbON0./qvDZdKAA7Zf.X6NC"),
	}
	repositorie := repositories.NewRepository(db)
	result, err := repositorie.Login(user)
	if err != nil {
		c.JSON(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	passwd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	fmt.Println("Crpy:", passwd)
	if err := bcrypt.CompareHashAndPassword(result.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		fmt.Println("error:", err)
		return c.JSON(fiber.Map{
			"message": "Incorrent password",
		})

	}
	return c.JSON(result)

}
