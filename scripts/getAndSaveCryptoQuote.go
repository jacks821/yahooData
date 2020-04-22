package yahooScraper

import (
	"github.com/piquette/finance-go/crypto"
	"log"
	"time"
)

func getCrypto(symbol string) {
	q, err := crypto.Get(symbol)
	if err != nil {
		log.Fatal(err)
	}

	if q == nil {
		return
	}

	t := time.Now()

	quote := Quote{Symbol: q.Quote.Symbol,
		ShortName:                  q.Quote.ShortName,
		RegularMarketChangePercent: q.Quote.RegularMarketChangePercent,
		RegularMarketPreviousClose: q.Quote.RegularMarketPreviousClose,
		RegularMarketPrice:         q.Quote.RegularMarketPrice,
		RegularMarketTime:          q.Quote.RegularMarketTime,
		RegularMarketChange:        q.Quote.RegularMarketChange,
		RegularMarketOpen:          q.Quote.RegularMarketOpen,
		RegularMarketDayHigh:       q.Quote.RegularMarketDayHigh,
		RegularMarketDayLow:        q.Quote.RegularMarketDayLow,
		RegularMarketVolume:        string(q.Quote.RegularMarketVolume),
		Bid:                        q.Quote.Bid,
		Ask:                        q.Quote.Ask,
		BidSize:                    q.Quote.BidSize,
		AskSize:                    q.Quote.AskSize,
		FiftyDayAverage:            q.Quote.FiftyDayAverage,
		TwoHundredDayAverage:       q.Quote.TwoHundredDayAverage,
		AverageDailyVolume3Month:   string(q.Quote.AverageDailyVolume3Month),
		AverageDailyVolume10Day:    string(q.Quote.AverageDailyVolume10Day),
		QuoteTime:                  t,
	}

	if q.Quote.RegularMarketVolume == 0 {
		quote.RegularMarketVolume = "Zero"
	}

	if q.Quote.AverageDailyVolume10Day == 0 {
		quote.AverageDailyVolume10Day = "Zero"
	}

	if q.Quote.AverageDailyVolume3Month == 0 {
		quote.AverageDailyVolume3Month = "Zero"
	}

	crypto := CryptoPair{Quote: quote,
		CirculatingSupply:   string(q.CirculatingSupply),
		VolumeLastDay:       string(q.VolumeLastDay),
		VolumeAllCurrencies: string(q.VolumeAllCurrencies),
	}

	if q.CirculatingSupply == 0 {
		crypto.CirculatingSupply = "Zero"
	}

	if q.VolumeLastDay == 0 {
		crypto.VolumeLastDay = "Zero"
	}

	if q.VolumeAllCurrencies == 0 {
		crypto.VolumeAllCurrencies = "Zero"
	}
	_ = SaveCrypto(crypto)
}
