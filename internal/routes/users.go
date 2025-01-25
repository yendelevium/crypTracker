package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yendelevium/crypTracker/models"
)

// TODO:
// Update all errors to be of type models.Error instead of just sending a message
func UserRouter() *fiber.App {
	userRouter := fiber.New()
	// TODO:
	// Store the data in the DB, salt and hash the password
	userRouter.Add(http.MethodPost, "/", func(c *fiber.Ctx) error {
		var userData models.User
		err := json.Unmarshal(c.Body(), &userData)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fmt.Sprintf("Error unmarshalling request body: %s", err))
		}
		userData.UserID = uuid.New().String()
		c.Status(http.StatusOK)
		return c.JSON(userData)
	})

	// TODO :
	// Get a particular user, including their details and/or watchlist
	userRouter.Add(http.MethodGet, "/:userId", func(c *fiber.Ctx) error {
		return c.JSON("WIP")
	})

	userRouter.Add(http.MethodGet, "/:userId/watchlist", func(c *fiber.Ctx) error {
		// Get watchlist from the DB and send it
		return c.JSON("Watchlist")
	})

	// INCOMPLETE :
	// I'll have 2 ways to edit the watchlist
	// One is to add more elements to the watchlist, the other is to subtract
	// I'll have a field which tells me which is it, add/subtract
	// Based on that I'll edit the DB
	// Also, AUTH to check if the delete msg is from the user who is the owner of the watchlist?
	// The userId in the req should match the userId of the watchlists
	userRouter.Add(http.MethodPut, "/:userId/watchlist", func(c *fiber.Ctx) error {
		var watchlist models.Watchlist
		err := json.Unmarshal(c.Body(), &watchlist)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fmt.Sprintf("Error unmarshalling request body: %s", err))
		}

		// Maybe have the insert/delete DBhandler as the value instead of bool?!
		allowedMethods := map[string]bool{
			"insert": true,
			"remove": true,
		}
		// handler, ok := allowedMethods[watchlist.Method]
		_, ok := allowedMethods[watchlist.Method]
		if !ok {
			c.Status(http.StatusBadRequest)
			return c.JSON("method should be either 'insert' or 'remove'")
		}
		// Do the Auth, you can get the userId that all the coins are supposed to be related to by the queryparams
		// Also check the userid of the sender and make sure all the user_ids of the watchlist are the same
		// If any user_id is different, DON'T commit the changes to the DB and return an error pointing out which UID is different

		// Ok user_id won't work as we need to check if they are logged in or not
		// We do by doing JWT to check if the logged in user is the same user sending the request

		for _, entry := range watchlist.Watching {
			// Update db
			// Idk how but like (watchlist.method, DELETE/INSERT from watchlist where userId = "" AND coinId = "")
			// I'll prolly have the corresponding method in the handler, just need to pass the params to it

			// Authorize the userID
			// Call the handler
			// If the AUTH fails at any point, discard all changes and send an error
			// THIS IS THE WRONG LOGIC, JUST USING THIS TO GET AN IDEA OF WHAT I'M SUPPOSED TO DO
			if watchlist.JWT != entry.UserID {
				c.Status(http.StatusUnauthorized)
				errorResponse := models.Error{
					Status:  http.StatusUnauthorized,
					Message: fmt.Sprintf("Unauthorized access to user: %s", entry.UserID),
				}
				return c.JSON(errorResponse)
			}

		}
		// Update the DB
		c.Status(http.StatusAccepted)
		return c.JSON(watchlist)
	})

	return userRouter
}
