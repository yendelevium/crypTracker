package database

import (
	"log"
	"os"

	"github.com/yendelevium/crypTracker/internal/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	Client *gorm.DB
}

func ConnectPostgres() (*DBClient, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	return &DBClient{Client: db}, err
}

// Change to dbClient.Client.Save() instead? so it can update AND seed?!
// This can be my bg goroutine for fecthing data every 10s
func (dbClient *DBClient) Seed() error {
	data, err := api.FetchCoinData()
	if err != nil {
		return err
	}
	// Inserting each element
	// Can't directly do dbClient.Create(data)
	// As data must be []*models.CoinData, but we have []models.CoinData
	for _, ele := range data {
		result := dbClient.Client.Create(&ele)
		if result.Error != nil {
			log.Println(result)
		}
	}

	return nil
}
