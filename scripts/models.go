package yahooScraper

import (
	"time"
)

type Quote struct {
	// Quote classifying fields.
	Symbol    string `json:"symbol"`
	ShortName string `json:"short_name"`

	// Regular session quote data.
	RegularMarketChangePercent float64 `json:"regular_market_change_percent"`
	RegularMarketPreviousClose float64 `json:"regular_market_previous_close"`
	RegularMarketPrice         float64 `json:"regular_market_price"`
	RegularMarketTime          int     `json:"regular_market_time"`
	RegularMarketChange        float64 `json:"regular_market_change"`
	RegularMarketOpen          float64 `json:"regular_market_open"`
	RegularMarketDayHigh       float64 `json:"regular_market_day_high"`
	RegularMarketDayLow        float64 `json:"regular_market_day_low"`
	RegularMarketVolume        string  `json:"regular_market_volume"`

	// Quote depth.
	Bid     float64 `json:"bid"`
	Ask     float64 `json:"ask"`
	BidSize int     `json:"bid_size"`
	AskSize int     `json:"ask_size"`

	// Averages.
	FiftyDayAverage      float64 `json:"fifty_day_average"`
	TwoHundredDayAverage float64 `json:"two_hundred_day_average"`

	// Volume metrics.
	AverageDailyVolume3Month string    `json:"average_daily_volume_3_month"`
	AverageDailyVolume10Day  string    `json:"average_daily_volume_10_day"`
	QuoteTime                time.Time `json:"quote_time"`
}

type CryptoPair struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Quote
	// Cryptocurrency-only fields.
	CirculatingSupply   string `json:"circulating_supply"`
	VolumeLastDay       string `json:"volume_24hr"`
	VolumeAllCurrencies string `json:"volume_all_currencies"`
}

type Equity struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Quote
	// Equity-only fields.
	LongName                string  `json:"long_name"`
	EpsTrailingTwelveMonths float64 `json:"eps_trailing_twelve_months"`
	EpsForward              float64 `json:"eps_forward"`
	EarningsTimestamp       int     `json:"earnings_timestamp"`
	EarningsTimestampStart  int     `json:"earnings_timestamp_start"`
	EarningsTimestampEnd    int     `json:"earnings_timestamp_end"`
	DividendDate            int     `json:"dividend_date"`
	TrailingPE              float64 `json:"trailing_pe"`
	ForwardPE               float64 `json:"forward_pe"`
	BookValue               float64 `json:"book_value"`
	PriceToBook             float64 `json:"price_to_book"`
	SharesOutstanding       string  `json:"shares_outstanding"`
	MarketCap               string  `json:"marketCap"`
}

type ForexPair struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Quote     `json:"quote"`
}
