// Endpoint to get coin data
// https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&x_cg_demo_api_key=YOUR_API_KEY
// See the JSON result to create corresponding structs and to parse the data

package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yendelevium/crypTracker/models"
)

// Return value/ function arguments?!
func FetchCoinData() ([]models.Coin, error) {
	err := godotenv.Load()
	if err != nil {
		// log.Fatalf("Error loading .env file: %s", err)
		return []models.Coin{}, err
	}
	apiKey := os.Getenv("COINGECKO_API_KEY")

	// Creating the request for the coinGecko endpoint
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.coingecko.com/api/v3/coins/markets", nil)
	if err != nil {
		// log.Fatalf("Failed to create http request to /coins/markets: %s", err)
		return []models.Coin{}, err
	}

	req.Header.Add("x_cg_demo_api_key", apiKey)
	req.Header.Add("accept", "application/json")

	q := req.URL.Query()
	q.Add("vs_currency", "usd")
	req.URL.RawQuery = q.Encode()
	// log.Println(req.URL.String())

	// Handling the response of the API
	resp, err := client.Do(req)
	if err != nil {
		// log.Fatalf("Error getting a response: %s", err)
		return []models.Coin{}, err
	}
	defer resp.Body.Close()
	// log.Println(resp.StatusCode)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		// log.Printf("Error reading the response body: %s", err)
		return []models.Coin{}, err
	}

	var respData []models.Coin
	err = json.Unmarshal(data, &respData)
	if err != nil {
		// log.Printf("Failed to unmarshal data: %s", err)
		return []models.Coin{}, err
	}

	return respData, nil
}
