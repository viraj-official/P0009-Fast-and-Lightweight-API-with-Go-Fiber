package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var userArray []User

func getUserHandler(c *fiber.Ctx) error {
	return c.JSON(userArray)
}

func createUserHandler(c *fiber.Ctx) error {
	var createUserReq CreateUserRequest

	if err := c.BodyParser(&createUserReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	newUser := User{
		ID:   len(userArray) + 1,
		Name: createUserReq.Name,
		Age:  createUserReq.Age,
	}

	userArray = append(userArray, newUser)

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func main() {
	app := fiber.New()

	// GET API endpoint
	app.Get("/users", getUserHandler)

	// POST API endpoint
	app.Post("/users/create", createUserHandler)
	log.Fatal(app.Listen(":8080"))
}
