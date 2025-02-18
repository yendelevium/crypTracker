package models

import "time"

// DB Models
type User struct {
	// MVP
	// Is it better for Id to be stored as string or uuid.UUID?
	UserID       string      `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	Username     string      `json:"username" gorm:"not null;unique"`
	Password     string      `json:"password" gorm:"not null"`
	ProfileImage string      `json:"profile_image" gorm:"default:https://plus.unsplash.com/premium_vector-1725479330926-997591b3ca07?q=80&w=2450&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"` //Take the img, store in teh cloud, and put the cloud link here
	CreatedAt    time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	Watchlist    []Watchlist `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Relationship to Watchlist
}

type Coin struct {
	CoinGeckoID  string      `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Symbol       string      `json:"symbol"`
	Name         string      `json:"name"`
	Image        string      `json:"image"`
	CurrentPrice float64     `json:"current_price"`
	MarketCap    float64     `json:"market_cap"`
	CreatedAt    time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	Watchlist    []Watchlist `gorm:"foreignKey:CoinGeckoID;constraint:OnDelete:CASCADE"` // Relationship to Watchlist
}

type Watchlist struct {
	// watchlist_user_id(w_user_id) or just user_id?!
	UserID      string    `json:"user_id" gorm:"primaryKey;autoIncrement:false;constraint:OnDelete:CASCADE"`
	CoinGeckoID string    `json:"coin_id" gorm:"primaryKey;autoIncrement:false;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships (optional but recommended for clarity)
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Coin *Coin `gorm:"foreignKey:CoinGeckoID;constraint:OnDelete:CASCADE"`
}

type Error struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
}
