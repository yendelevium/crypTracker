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
	"gorm.io/gorm/clause"
)

// TODO :
// Create an error handler to handle errors coz it's getting way too fucking redundant

// TODO:
// Create functions for redundant code like signout and password hashing

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
			UserID  string `json:"user_id"`
		}{
			Message: fmt.Sprintf("Successfully logged in user : %s!", user.Username),
			UserID:  fetchedUser.UserID,
		})
	})

	// Should use DELETE as I'm technically deleting the JWT
	// But I'll be redirecting to this after deleting user, and I can only redirect GET methods
	userRouter.Add(http.MethodGet, "/signout", func(c *fiber.Ctx) error {
		// Doesn't work
		// c.ClearCookie("Authorization")

		// So replace the Auth cookie so it expires immediately(as its expiry date is 1hr before time.Now())
		deleteJWT := &fiber.Cookie{
			Name:    "Authorization",
			Value:   "",
			Expires: time.Now().Add(-time.Hour),
			Path:    "/", // To match the original Path
		}
		c.Status(http.StatusOK)
		c.Cookie(deleteJWT)
		return c.JSON(struct {
			Message string `json:"message"`
		}{
			Message: "Successfully signed out",
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

	userRouter.Add(http.MethodDelete, "/:userId", middleware.AuthorizeUser, func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		// Deleting the user
		// Select(clause.Associations) for cascading delete
		result := dbClient.Client.Select(clause.Associations).Where("user_id = ?", userId).Delete(&models.User{})
		if result.RowsAffected == 0 {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: "User doesn't exist",
				Status:  http.StatusBadRequest,
			})
		}

		// Signing out
		// Logic is redundant with /users/signout
		// I tried c.Redirect(), it works, but the message is "successfully signed out" instead of "successfull deleted"
		// CAN pass a query param to /users/signout, but I need a better way
		deleteJWT := &fiber.Cookie{
			Name:    "Authorization",
			Value:   "",
			Expires: time.Now().Add(-time.Hour),
			Path:    "/", // To match the original Path
		}
		c.Status(http.StatusOK)
		c.Cookie(deleteJWT)
		c.Status(http.StatusOK)
		return c.JSON(struct {
			Message string `json:"message"`
		}{
			Message: fmt.Sprintf("Successfully deleted user : %s", userId),
		})

	})

	userRouter.Add(http.MethodPut, "/:userId", middleware.AuthorizeUser, func(c *fiber.Ctx) error {
		// TODO:
		// This doesn't check for changing user_id or created_at fields
		// Do that

		// Send the wholeass struct which u need to change(Created time also coz otherwise it will set it to 0)
		// Ik fr a FACT that I won't allow them to changethe userId
		// So idts I need ON UPDATE CASCADE?!
		var newUser models.User
		err := json.Unmarshal(c.Body(), &newUser)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: "Cannot unmarshall request body",
				Status:  http.StatusBadRequest,
			})
		}

		// Redundant password hashing logic, create new function for this?
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("Couldn't hash password :%s", err),
				Status:  http.StatusBadRequest,
			})
		}
		newUser.Password = string(hashedPassword)
		result := dbClient.Client.Save(&newUser)
		if result.Error != nil {
			c.Status(http.StatusInternalServerError)
			c.JSON(models.Error{
				Message: fmt.Sprintf("Couldn't update user %s : %s", newUser.UserID, result.Error),
				Status:  http.StatusInternalServerError,
			})
		}

		return c.JSON(newUser)

	})

	userRouter.Add(http.MethodGet, "/:userId/watchlist", middleware.AuthorizeUser, func(c *fiber.Ctx) error {
		userId := c.Params("userID")
		var watchlistCoins []models.Coin
		// Retreiveing all watched coins via the database using JOINS
		result := dbClient.Client.
			Table("watchlists").
			// Select("watchlists.coin_gecko_id,coins.symbol,coins.name,coins.current_price,coins.market_cap,coins.updated_at,coins.image").
			Joins("inner join coins ON watchlists.coin_gecko_id = coins.coin_gecko_id AND watchlists.user_id = ?", userId).
			Scan(&watchlistCoins)
		if result.Error != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: result.Error.Error(),
				Status:  http.StatusBadRequest,
			})
		}

		c.Status(http.StatusOK)
		return c.JSON(watchlistCoins)
	})

	userRouter.Add(http.MethodPost, "/:userId/watchlist", middleware.AuthorizeUser, func(c *fiber.Ctx) error {
		// Idts i need the user as I'm already verifying in the the middleware and i have the id in the path
		// Req body must be
		// {
		// 	"coin_id":"bitcoin"
		// }
		var watchlist models.Watchlist
		userId := c.Params("userId")

		err := json.Unmarshal(c.Body(), &watchlist)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("Failed to unmarshall request body: %s", err),
				Status:  http.StatusBadGateway,
			})
		}

		// Checking if the coin exists in the DB
		result := dbClient.Client.First(&models.Coin{}, "coin_gecko_id = ?", watchlist.CoinGeckoID)
		if result.Error != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: result.Error.Error(),
				Status:  http.StatusBadGateway,
			})
		}

		// If it exists, add the relation to watchlists schema
		watchlist.UserID = userId
		log.Println(watchlist)
		result = dbClient.Client.Create(&watchlist)
		if result.Error != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("Couldn't add to watchlist, %s", result.Error),
				Status:  http.StatusBadGateway,
			})
		}

		c.Status(http.StatusCreated)
		return c.JSON(watchlist)

	})

	userRouter.Add(http.MethodDelete, ":userId/watchlist", middleware.AuthorizeUser, func(c *fiber.Ctx) error {
		// Req body must be
		// {
		// 	"coin_id":"bitcoin"
		// }
		userId := c.Params("userId")
		var watchlist models.Watchlist
		err := json.Unmarshal(c.Body(), &watchlist)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("Couldn't unmarshall request body : %s", err),
				Status:  http.StatusBadRequest,
			})
		}

		// Deleting the watched coin from the watchlist
		result := dbClient.Client.
			Where("user_id = ? AND coin_gecko_id = ? ", userId, watchlist.CoinGeckoID).
			Delete(&models.Watchlist{})
		if result.Error != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.Error{
				Message: fmt.Sprintf("Couldn't delete record : %s", result.Error),
				Status:  http.StatusInternalServerError,
			})
		}

		c.Status(http.StatusOK)
		return c.JSON(struct {
			Message string `json:"message"`
		}{
			Message: fmt.Sprintf("Successfully deleted %s from watchlist", watchlist.CoinGeckoID),
		})

	})

	return userRouter
}
