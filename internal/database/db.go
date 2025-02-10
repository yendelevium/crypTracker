package database

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/yendelevium/crypTracker/internal/api"
	"github.com/yendelevium/crypTracker/models"
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

func (dbClient *DBClient) ScrapeData() {
	data, err := api.FetchCoinData()
	if err != nil {
		log.Println(err)
	}

	// If I run EVERY update concurrently, due to write-write conflict, the DB kindof implements a lock
	// This makes it counter productive as this makes it take MORE TIME PER QUERY
	// But, we can limit the concurrent routines to 5 using a buffered channel

	// If multiple goroutines are updating the same table concurrently, lock contention could be slowing things down.
	// PostgreSQL enforces row-level locks for updates, and if many updates are happening at once,
	// transactions might queue up, increasing execution time.

	// This works
	var wg sync.WaitGroup
	slots := make(chan struct{}, 5)

	for _, ele := range data {
		wg.Add(1)
		go func(e models.Coin) {
			defer wg.Done()
			slots <- struct{}{} // Acquire a concurrency slot. IF it's full, the channel blocks

			result := dbClient.Client.Model(&e).Updates(e)
			if result.Error != nil {
				log.Println(result)
			}

			<-slots // Release the slot so the next routine can run
		}(ele)
	}

	wg.Wait()
}

func (dbClient *DBClient) StartScraping() {
	ticker := time.NewTicker(15 * time.Second)
	// To immediately start scraping when the function is called, we use <-ticker.C
	// If we use for range ticker.C, it will first wait for 15s, and THEN start the first scrapw
	for ; ; <-ticker.C {
		log.Println("Scraping Data")
		go dbClient.ScrapeData() // No blocking, allows overlap if needed
	}
}
