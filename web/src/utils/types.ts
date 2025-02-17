// type Coin struct {
// 	CoinGeckoID  string      `json:"id" gorm:"primaryKey;autoIncrement:false"`
// 	Symbol       string      `json:"symbol"`
// 	Name         string      `json:"name"`
// 	Image        string      `json:"image"`
// 	CurrentPrice float64     `json:"current_price"`
// 	MarketCap    float64     `json:"market_cap"`
// 	CreatedAt    time.Time   `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt    time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
// 	Watchlist    []Watchlist `gorm:"foreignKey:CoinGeckoID;constraint:OnDelete:CASCADE"` // Relationship to Watchlist
// }

type TCoin = {
    id: string,
    symbol: string,
    name: string,
    image: string,
    current_price: number,
    market_cap: number,
    created_at: Date,
    updated_at: Date,
}

// type User struct {
// 	// MVP
// 	// Is it better for Id to be stored as string or uuid.UUID?
// 	UserID       string      `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
// 	Username     string      `json:"username" gorm:"not null;unique"`
// 	Password     string      `json:"password" gorm:"not null"`
// 	ProfileImage string      `json:"profile_image" gorm:"default:https://unsplash.com/illustrations/a-colorful-pattern-with-a-green-circle-in-the-middle-h54uX2BEclQ"` //Take the img, store in teh cloud, and put the cloud link here
// 	CreatedAt    time.Time   `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt    time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
// 	Watchlist    []Watchlist `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Relationship to Watchlist
// }

type TUser = {
    user_id: string,
    username: string,
    password?: string,
    profile_image: string,
    created_at?: Date,
    updated_at?: Date,
}

export type{
    
    TCoin,
    TUser,
}