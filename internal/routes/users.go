package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yendelevium/crypTracker/internal/database"
	"github.com/yendelevium/crypTracker/middleware"
	"github.com/yendelevium/crypTracker/models"
	"golang.org/x/crypto/bcrypt"
)

// Send a JWT, stored in the cookies
// MapClaims is the payload in the JWT, which will be used for authorization
func CreateJWT(userData models.User) (*fiber.Cookie, error) {
	// Create a token with the signing method and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userData.UserID,
		"username": userData.Username,
	})

	// Create the JWT by adding and signing it with the SECRET
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Printf("Couldn't generate JWT: %s", err)
		return &fiber.Cookie{}, err
	}

	// Create a fiber.Cookie to store the JWT in the client side
	jwtCookie := fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		HTTPOnly: true,
		SameSite: "Lax", //So the client properly stores the cookie, and also something else, google it ig
		// Set Secure : true ONLY IN PRODUCTION over https
		// Secure:   true,
		// Make the cookie expire a week from now
		Expires: time.Now().Add(time.Duration(time.Hour * 24 * 7)),
	}
	return &jwtCookie, nil
}

// TODO:
// Update all errors to be of type models.Error instead of just sending a message
func UserRouter(dbClient *database.DBClient) *fiber.App {
	userRouter := fiber.New()

	// Add user route
	userRouter.Add(http.MethodPost, "/", func(c *fiber.Ctx) error {
		var userData models.User
		err := json.Unmarshal(c.Body(), &userData)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fmt.Sprintf("Error unmarshalling request body: %s", err))
		}

		if userData.Username == "" || userData.Password == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: "Username and/or Password fields can't be NIL",
				Status:  http.StatusBadRequest,
			})
		}

		// Hashing the password and giving a UUID to the user
		userData.UserID = uuid.New().String()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("Couldn't hash password :%s", err),
				Status:  http.StatusBadRequest,
			})
		}
		userData.Password = string(hashedPassword)
		result := dbClient.Client.Create(&userData)
		if result.Error != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("Couldn't add user : %s", result.Error),
				Status:  http.StatusBadRequest,
			})
		}

		// Creating a JWT(which will be in a *fiber.Cookie)
		// This is so the user is automatically logged in after registration
		jwtCookie, err := CreateJWT(userData)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.Error{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}

		c.Cookie(jwtCookie)
		c.Status(http.StatusCreated)
		return c.JSON(struct {
			Message string `json:"message"`
		}{
			Message: "201 : Successfully registered user",
		})
	})

	userRouter.Add(http.MethodPost, "/login", func(c *fiber.Ctx) error {
		var user models.User
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.Error{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}

		// Fetching the user with the given username
		var fetchedUser models.User
		result := dbClient.Client.First(&fetchedUser, "username = ?", user.Username)
		if result.Error != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("User doesn't exist: %s", result.Error.Error()),
				Status:  http.StatusBadRequest,
			})
		}
		if fetchedUser.UserID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: "Incorrect username or password",
				Status:  http.StatusBadRequest,
			})
		}

		err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(user.Password))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: "Incorrect username or password",
				Status:  http.StatusBadRequest,
			})
		}

		jwtCookie, err := CreateJWT(fetchedUser)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.Error{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}

		c.Cookie(jwtCookie)
		c.Status(http.StatusOK)
		return c.JSON(struct {
			Message string `json:"message"`
		}{
			Message: fmt.Sprintf("Successfully logged in user : %s!", user.Username),
		})
	})

	// TODO :
	// Get a particular user, including their details and/or watchlist
	userRouter.Add(http.MethodGet, "/:userId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		var user models.User
		result := dbClient.Client.First(&user, "user_id = ?", userId)
		if result.Error != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("User '%s' doesn't exist", userId),
				Status:  http.StatusBadRequest,
			})
		}

		// Removing hashed password and update time
		c.Status(http.StatusOK)
		return c.JSON(models.User{
			UserID:       user.UserID,
			Username:     user.Username,
			ProfileImage: user.ProfileImage,
			CreatedAt:    user.CreatedAt,
		})
	})

	userRouter.Add(http.MethodGet, "/:userId/watchlist", middleware.AuthorizeUser, func(c *fiber.Ctx) error {
		// Get watchlist from the DB and send it
		return c.JSON("Watchlist")
	})
	return userRouter
}
