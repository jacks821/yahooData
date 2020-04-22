package yahooScraper

import (
	"github.com/piquette/finance-go/forex"
	"log"
	"time"
)

func getForex(symbol string) {
	q, err := forex.Get(symbol)
	if err != nil {
		log.Fatal(err)
	}

	if q == nil {
		return
	}

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
		QuoteTime:                  time.Now(),
	}

	if q.Quote.AverageDailyVolume10Day == 0 {
		quote.AverageDailyVolume10Day = "Zero"
	}

	if q.Quote.AverageDailyVolume3Month == 0 {
		quote.AverageDailyVolume3Month = "Zero"
	}

	forex := ForexPair{Quote: quote}

	forex = SaveForex(forex)
}
