package helper

import (
	cof "gayandn/configration"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetBusinessIdClaim(c *fiber.Ctx) string {

	claims := claims(c)
	businessId, ok := claims["businessId"].(string)
	if !ok {
		panic("Couldn't parse userid as string")
	}
	return businessId
}

func claims(c *fiber.Ctx) jwt.MapClaims {
	//Get Token
	reqToken := c.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(reqToken, claims, keyFunc)
	if err != nil {
		panic(err)
	}
	return claims
}

func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(cof.Config("SECRET")), nil
}
