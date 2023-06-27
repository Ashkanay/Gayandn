package handler

import (
	"errors"
	"fmt"
	"gayandn/configration"
	"gayandn/contract/request"
	"gayandn/contract/response"
	"gayandn/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// https://gowebexamples.com/password-hashing/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUsername(u string) (*model.User, error) {
	db := configration.DB
	var user model.User

	err := db.Where(&model.User{LoginName: u}).Find(&user).Error
	fmt.Println("err: ", err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("ErrRecordNotFound:", err)
			return nil, err
		}
		return nil, err
	}
	if user.LoginName == "" {
		fmt.Println("LoginName is null:", err)
		return nil, err
	}

	return &user, err
}

// Login get user and password
func Login(c *fiber.Ctx) error {
	input := new(request.LoginRequest)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	fmt.Println("LoginName:  ", input.Identity)
	fmt.Println("Password:  ", input.Password)

	//identity := input.Identity
	pass := input.Password
	user, err := new(model.User), *new(error)

	user, err = getUserByUsername(input.Identity)
	fmt.Println("getUserByUsername: ", err)
	fmt.Println("user: ", user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

	if !CheckPasswordHash(pass, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	//claims["businessId"] = user.BusinessId
	claims["loginname"] = user.LoginName
	claims["userid"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(configration.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var ud response.LoginResultResponse
	if user != nil {
		ud = response.LoginResultResponse{
			Data:    response.AuthneticationResponse{Token: t},
			UserId:  user.ID,
			Name:    user.LoginName,
			Success: true,
			Code:    200,
			Message: "Success login",
		}
	}

	return c.JSON(fiber.Map{"status": "success", "Code": "200", "message": "Success login", "LoginResultResponse": ud})
}
