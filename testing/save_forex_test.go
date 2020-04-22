package yahooData

import (
	"testing"
	"time"

	scripts "../scripts"
)

func TestSaveForex(t *testing.T) {

	a := time.Now()

	quote := scripts.Quote{Symbol: "Fake",
		ShortName:                  "fakeShortName",
		RegularMarketChangePercent: 1.0,
		RegularMarketPreviousClose: 1.0,
		RegularMarketPrice:         1.0,
		RegularMarketTime:          1,
		RegularMarketChange:        1.0,
		RegularMarketOpen:          1.0,
		RegularMarketDayHigh:       1.0,
		RegularMarketDayLow:        1.0,
		RegularMarketVolume:        1,
		Bid:                        10.0,
		Ask:                        12.0,
		BidSize:                    1,
		AskSize:                    1,
		PreMarketPrice:             1.0,
		PreMarketChange:            1.0,
		PreMarketChangePercent:     1.0,
		PreMarketTime:              1,
		PostMarketPrice:            1.0,
		PostMarketChange:           1.0,
		PostMarketChangePercent:    1.0,
		PostMarketTime:             1,
		FiftyTwoWeekLow:            1.0,
		FiftyTwoWeekHigh:           2.0,
		FiftyDayAverage:            1.0,
		TwoHundredDayAverage:       1.0,
		AverageDailyVolume3Month:   100,
		AverageDailyVolume10Day:    120,
		QuoteSource:                "source",
		CurrencyID:                 "currencyId",
		IsTradeable:                true,
		QuoteTime:                  a,
	}

	forex := scripts.ForexPair{Quote: quote}

	resp := scripts.saveForex(forex)

	if resp.Quote.Symbol != quote.Symbol {
		t.Error("expected", quote.Symbol, "got", resp.Quote.Symbol)
	}

	if resp.Quote.ShortName != quote.ShortName {
		t.Error("expected", quote.ShortName, "got", resp.Quote.ShortName)
	}

	if resp.Quote.RegularMarketChangePercent != quote.RegularMarketChangePercent {
		t.Error("expected", quote.RegularMarketChangePercent, "got", resp.Quote.RegularMarketChangePercent)
	}

	if resp.Quote.RegularMarketPrice != quote.RegularMarketPrice {
		t.Error("expected", quote.RegularMarketPrice, "got", resp.Quote.RegularMarketPrice)
	}

	if resp.Quote.RegularMarketTime != quote.RegularMarketTime {
		t.Error("expected", quote.RegularMarketTime, "got", resp.Quote.RegularMarketTime)
	}

	if resp.Quote.RegularMarketChange != quote.RegularMarketChange {
		t.Error("expected", quote.RegularMarketChange, "got", resp.Quote.RegularMarketChange)
	}

	if resp.Quote.RegularMarketOpen != quote.RegularMarketOpen {
		t.Error("expected", quote.RegularMarketOpen, "got", resp.Quote.RegularMarketOpen)
	}

	if resp.Quote.RegularMarketDayHigh != quote.RegularMarketDayHigh {
		t.Error("expected", quote.RegularMarketDayHigh, "got", resp.Quote.RegularMarketDayHigh)
	}

	if resp.Quote.RegularMarketDayLow != quote.RegularMarketDayLow {
		t.Error("expected", quote.RegularMarketDayLow, "got", resp.Quote.RegularMarketDayLow)
	}

	if resp.Quote.RegularMarketVolume != quote.RegularMarketVolume {
		t.Error("expected", quote.RegularMarketVolume, "got", resp.Quote.RegularMarketVolume)
	}

	if resp.Quote.Bid != quote.Bid {
		t.Error("expected", quote.Bid, "got", resp.Quote.Bid)
	}

	if resp.Quote.Ask != quote.Ask {

		t.Error("expected", quote.Ask, "got", resp.Quote.Ask)
	}

	if resp.Quote.BidSize != quote.BidSize {
		t.Error("expected", quote.BidSize, "got", resp.Quote.BidSize)
	}

	if resp.Quote.AskSize != quote.AskSize {
		t.Error("expected", quote.AskSize, "got", resp.Quote.AskSize)
	}

	if resp.Quote.PreMarketPrice != quote.PreMarketPrice {
		t.Error("expected", quote.PreMarketPrice, "got", resp.Quote.PreMarketPrice)
	}

	if resp.Quote.PreMarketChange != quote.PreMarketChange {
		t.Error("expected", quote.PreMarketChange, "got", resp.Quote.PreMarketChange)
	}

	if resp.Quote.PostMarketChangePercent != quote.PostMarketChangePercent {
		t.Error("expected", quote.RegularMarketDayLow, "got", resp.Quote.RegularMarketDayLow)
	}

	if resp.Quote.PostMarketTime != quote.PostMarketTime {
		t.Error("expected", quote.PostMarketTime, "got", resp.Quote.PostMarketTime)
	}
	if resp.Quote.FiftyTwoWeekLow != quote.FiftyTwoWeekLow {
		t.Error("expected", quote.FiftyTwoWeekLow, "got", resp.Quote.FiftyTwoWeekLow)
	}

	if resp.Quote.FiftyTwoWeekHigh != quote.FiftyTwoWeekHigh {
		t.Error("expected", quote.FiftyTwoWeekHigh, "got", resp.Quote.FiftyTwoWeekHigh)
	}
	if resp.Quote.FiftyDayAverage != quote.FiftyDayAverage {
		t.Error("expected", quote.FiftyDayAverage, "got", resp.Quote.FiftyDayAverage)
	}
	if resp.Quote.TwoHundredDayAverage != quote.TwoHundredDayAverage {
		t.Error("expected", quote.TwoHundredDayAverage, "got", resp.Quote.TwoHundredDayAverage)
	}
	if resp.Quote.AverageDailyVolume3Month != quote.AverageDailyVolume3Month {
		t.Error("expected", quote.AverageDailyVolume3Month, "got", resp.Quote.AverageDailyVolume3Month)
	}
	if resp.Quote.AverageDailyVolume10Day != quote.AverageDailyVolume10Day {
		t.Error("expected", quote.AverageDailyVolume10Day, "got", resp.Quote.AverageDailyVolume10Day)
	}
	if resp.Quote.QuoteSource != quote.QuoteSource {
		t.Error("expected", quote.QuoteSource, "got", resp.Quote.QuoteSource)
	}
	if resp.Quote.CurrencyID != quote.CurrencyID {
		t.Error("expected", quote.CurrencyID, "got", resp.Quote.CurrencyID)
	}
	if resp.Quote.IsTradeable != quote.IsTradeable {
		t.Error("expected", quote.IsTradeable, "got", resp.Quote.IsTradeable)
	}
	if resp.Quote.QuoteTime != quote.QuoteTime {
		t.Error("expected", quote.QuoteTime, "got", resp.Quote.QuoteTime)
	}
}
