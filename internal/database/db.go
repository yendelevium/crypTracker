package database

import (
	"log"
	"os"
	"sync"
	"time"

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

func (dbClient *DBClient) ScrapeData(wg *sync.WaitGroup) error {
	defer wg.Done()
	data, err := api.FetchCoinData()
	if err != nil {
		log.Println(err)
		return err
	}
	// Inserting each element
	for _, ele := range data {
		result := dbClient.Client.Model(&ele).Updates(ele)
		if result.Error != nil {
			log.Println(result)
		}
	}

	return nil
}

func (dbCLient *DBClient) StartScraping() {
	ticker := time.NewTicker(15 * time.Second)
	for ; ; <-ticker.C {
		wg := &sync.WaitGroup{}
		log.Println("Scraping Data")
		wg.Add(1)
		go dbCLient.ScrapeData(wg)
		wg.Wait()
	}
}
