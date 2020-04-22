package yahooScraper

import (
	migrations "./migrations"
	"fmt"
	"log"
	"time"
)

func GetYahooData() {
	migrations.Migrate()
	start := time.Now()
	count := 0

	for _, symbol := range CRYPTO_SYMBOLS {
		fmt.Println("Finding for Crypto ", symbol)
		getCrypto(symbol)
		count++
	}

	for _, symbol := range FOREX_SYMBOLS {
		fmt.Println("Finding for Forex ", symbol)
		getForex(symbol)
		count++
	}
	s := fmt.Sprintf("Finished executing %v symbols in %d time", count, time.Since(start))
	fmt.Println(s)
}
