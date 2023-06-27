package handler

import (
	"fmt"
	"gayandn/configration"
	"log"

	"gayandn/model"

	"github.com/gofiber/fiber/v2"
)

func validUser(id string, p string) bool {
	db := configration.DB
	var user model.User
	db.First(&user, id)
	if user.LoginName == "" {
		return false
	}
	// if !CheckPasswordHash(p, user.Password) {
	// 	return false
	// }
	return CheckPasswordHash(p, user.Password)
}

// GetAllUsers query all users
func GetAllUsers(c *fiber.Ctx) error {
	db := configration.DB
	var usres []model.User

	//Get All User by Businesses Filter
	db.Find(&usres)
	return c.JSON(fiber.Map{"status": "success", "message": "All usres", "data": usres})
}

// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := configration.DB
	var user model.User
	db.Find(&user, "id = ?", id)
	if user.LoginName == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": user})
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {

	input := new(model.User)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	user, err := getUserByUsername(input.LoginName)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}

	if user != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "This username already you have it", "data": err})
	}

	db := configration.DB

	password, err := HashPassword(input.Password)
	if err != nil {
		log.Fatalf("error while generating random string: %s", err)
	}
	fmt.Println(password)
	input.Password = password
	if err := db.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": input.LoginName})
}

// UpdateUser update user
func UpdateUser(c *fiber.Ctx) error {
	input := new(model.User)

	//var uui UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")

	db := configration.DB
	var user model.User

	db.First(&user, id)
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "Users successfully updated", "data": user})
}

// Soft Delete User
func DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")

	if !validUser(id, pi.Password) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
	}

	db := configration.DB
	var user model.User

	db.First(&user, id)
	db.Delete(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "Users successfully deleted", "data": nil})
}

func InsertUser() {
	password := "123"
	hash, _ := HashPassword(password)

	password2 := "321"
	hash2, _ := HashPassword(password2)

	db := configration.DB
	//var users []*model.User
	users := []*model.User{
		{FirstName: "Barez", MiddleName: "Hussen", LastName: "Mohemmed", LoginName: "a", Password: string(hash), Gender: "Male", Mobil: 0750, Email: "barez@24group.net"},

		{FirstName: "Barez1", MiddleName: "Hussen1", LastName: "Mohemmed1", LoginName: "b", Password: string(hash2), Gender: "Male", Mobil: 0750, Email: "barez1@24group.net"},
	}
	// users = []*model.User{
	// 	{FirstName: "Barez", MiddleName: "Hussen", LastName: "Mohemmed", LoginName: "a", Password: string(hash), Gender: "Male", Mobil: 0750, Email: "barez@24group.net",
	// 		Inheritance: model.Inheritance{
	// 			Id:        1,
	// 			IsActive:  true,
	// 			IsDeleted: false,
	// 			Deleted:   false,
	// 		}},

	// 	{FirstName: "Barez1", MiddleName: "Hussen1", LastName: "Mohemmed1", LoginName: "b", Password: string(hash2), Gender: "Male", Mobil: 0750, Email: "barez1@24group.net",
	// 		Inheritance: model.Inheritance{
	// 			Id:        2,
	// 			IsActive:  true,
	// 			IsDeleted: false,
	// 			Deleted:   false,
	// 		}},
	// }
	fmt.Println("Before crate")

	result := db.Create(users) // pass a slice to insert multiple row

	if result.Error != nil {
		log.Fatalf("Failed insert *************************************************")
		panic(result.Error)
	} // returns error

	fmt.Println(result.RowsAffected) // returns inserted records count
}
