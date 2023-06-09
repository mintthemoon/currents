package trading

import (
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/mintthemoon/currents/math"
)

type Ticker struct {
	BaseAsset string `json:"base_asset"`
	QuoteAsset string `json:"quote_asset"`
	BaseVolume decimal.Big `json:"base_volume"`
	QuoteVolume decimal.Big `json:"quote_volume"`
	Price decimal.Big `json:"price"`
	Time time.Time `json:"time"`
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
		r.Price.Quo(math.One, &t.Price)
	}
	return r
}