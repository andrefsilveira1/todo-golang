package controllers

import (
	"fmt"
	"main/db"
	"main/models"
	"main/repositories"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

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

	if err := bcrypt.CompareHashAndPassword(result.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		fmt.Println("error:", err)
		return c.JSON(fiber.Map{
			"message": "Incorrent password",
		})

	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(result.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "login invalid",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	db, erro := db.Connect()
	if erro != nil {
		panic(erro)
	}
	defer db.Close()

	repositorie := repositories.NewRepository(db)
	result, err := repositorie.User(claims.Issuer)

	return c.JSON(result)

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
