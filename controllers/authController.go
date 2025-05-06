package controllers

import (
	"blog/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		Id:        1111,
		FirstName: "Jake",
		Email:     "jhond@gmail.com",
		Password:  "1979",
	}
	user.LastName = "Doe"

	// user:=models.User{11,"Jhon","Doe","jhond@gmail.com","1979"}
	return c.JSON(user)
}
