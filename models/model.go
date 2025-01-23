package models

type CoinData struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	CurrentPrice float64 `json:"current_price"`
	MarketCap    float64 `json:"market_cap"`
}
