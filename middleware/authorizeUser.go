package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yendelevium/crypTracker/models"
)

// Middleware to authorize user
func AuthorizeUser(c *fiber.Ctx) error {
	tokenString := c.Cookies("Authorization")
	log.Println(tokenString)
	// From the docs
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.Status(http.StatusUnauthorized)
		// If you get the error
		// token is malformed: token contains an invalid number of segments
		// This can also mean that the cookie itself if not there bro( along with the obv malformed token)
		return c.JSON(models.Error{
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		})
	}

	// Authorization logic
	claims, ok := token.Claims.(jwt.MapClaims)
	// If payload doesn't exit or token is invalid, user is Unauthorized
	if !(ok && token.Valid) {
		c.Status(http.StatusUnauthorized)
		return c.JSON(models.Error{
			Message: "Payload doesn't exist, JWT has been modified",
			Status:  http.StatusUnauthorized,
		})
	}

	// Checking if the cookie's userId and the actual userId are same
	// No need to check for token expiration as the cookie automatically has an expiry date
	// Should i be worried about this? Idk, maybe i'll add an expiration in the JWT payload later
	// And check if current time is> than expiration or not
	log.Println(claims["user_id"])
	if claims["user_id"] != c.Params("userId") {
		c.Status(http.StatusUnauthorized)
		return c.JSON(models.Error{
			Message: "Unauthorized, this is not your account",
			Status:  http.StatusUnauthorized,
		})
	}
	return c.Next()

}
