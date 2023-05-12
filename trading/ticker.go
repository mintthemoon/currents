package trading

import (
	"time"

	"github.com/ericlagergren/decimal"
)

type Ticker struct {
	BaseAsset string `json:"base_asset"`
	QuoteAsset string `json:"quote_asset"`
	BaseVolume decimal.Big `json:"base_volume"`
	QuoteVolume decimal.Big `json:"quote_volume"`
	Price decimal.Big `json:"price"`
	Time time.Time `json:"time"`
}

func TickerFromCandles(candles []*Candle) *Ticker {
	if len(candles) == 0 {
		return &Ticker{}
	}
	ticker := &Ticker{
		BaseAsset: candles[0].BaseAsset,
		QuoteAsset: candles[0].QuoteAsset,
		Time: candles[len(candles) - 1].End,
	}
	i := len(candles) - 1
	for ; i >= 0; i-- {
		if candles[i].Open.Cmp(&decimal.Big{}) != 0 {
			ticker.Price = candles[i].Close
			break
		}
	}
	for ; i >= 0; i-- {
		ticker.BaseVolume.Add(&ticker.BaseVolume, &candles[i].BaseVolume)
		ticker.QuoteVolume.Add(&ticker.QuoteVolume, &candles[i].QuoteVolume)
	}
	return ticker
}

func (t *Ticker) Reversed() *Ticker {
	r := &Ticker{
		BaseAsset: t.QuoteAsset,
		QuoteAsset: t.BaseAsset,
		BaseVolume: t.QuoteVolume,
		QuoteVolume: t.BaseVolume,
		Time: t.Time,
	}
	if t.Price.Cmp(&decimal.Big{}) != 0 {
		one := &decimal.Big{}
		one.SetUint64(1)
		r.Price.Quo(one, &t.Price)
	}
	return r
}