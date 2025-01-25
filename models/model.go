package models

type CoinData struct {
	CoinGeckoID  string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	CurrentPrice float64 `json:"current_price"`
	MarketCap    float64 `json:"market_cap"`
}

// DB Models
type User struct {
	// MVP
	// Is it better for Id to be stored as string or uuid.UUID?
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	// ProfileImage string `json:"profileImage"` //Take the img, store in teh cloud, and put the cloud link here
}

type Coin struct {
	CoinID string `json:"coin_id"`
	CoinData
}

type WatchlistEntry struct {
	// watchlist_user_id(w_user_id) or just user_id?!
	UserID string `json:"user_id"`
	CoinID string `json:"coin_id"`
}

type Watchlist struct {
	Method   string           `json:"method"`
	JWT      string           `json:"jwt"`
	Watching []WatchlistEntry `json:"watching"`
}

type Error struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
}
