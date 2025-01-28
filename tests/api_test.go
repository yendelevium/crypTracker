package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yendelevium/crypTracker/internal/database"
	"github.com/yendelevium/crypTracker/internal/routes"
	"gorm.io/gorm"
)

// DOESN'T WORK
func TestApi(t *testing.T) {
	app := fiber.New()
	app.Mount("/api", routes.CoinRouter(&database.DBClient{Client: &gorm.DB{}}))

	req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	// This is for internal debugguing. app.Test() will make the request u pass through it
	resp, _ := app.Test(req)
	// This is for testing the output using testify
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
