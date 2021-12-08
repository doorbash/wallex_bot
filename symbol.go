package main

import "fmt"

type Symbol struct {
	Symbol             string  `json:"symbol"`
	BaseAsset          string  `json:"baseAsset"`
	BaseAssetPrecision int     `json:"baseAssetPrecision"`
	QuoteAsset         string  `json:"quoteAsset"`
	QuotePrecision     int     `json:"quotePrecision"`
	FaName             string  `json:"faName"`
	FaBaseAsset        string  `json:"faBaseAsset"`
	FaQuoteAsset       string  `json:"faQuoteAsset"`
	StepSize           int     `json:"stepSize"`
	TickSize           int     `json:"tickSize"`
	MinQty             float64 `json:"minQty"`
	MinNotional        float64 `json:"minNotional"`
	Stats              struct {
		BidPrice         string  `json:"bidPrice"`
		AskPrice         string  `json:"askPrice"`
		Two4HCh          float64 `json:"24h_ch"`
		Two4HVolume      string  `json:"24h_volume"`
		Two4HQuoteVolume string  `json:"24h_quoteVolume"`
		Two4HHighPrice   string  `json:"24h_highPrice"`
		Two4HLowPrice    string  `json:"24h_lowPrice"`
		LastPrice        string  `json:"lastPrice"`
		LastQty          string  `json:"lastQty"`
		BidVolume        string  `json:"bidVolume"`
		AskVolume        string  `json:"askVolume"`
		BidCount         int     `json:"bidCount"`
		AskCount         int     `json:"askCount"`
	} `json:"stats"`
}

func (s Symbol) GetBid() string {
	return ParsePrice(s.Stats.BidPrice, s.QuotePrecision)
}

func (s Symbol) GetAsk() string {
	return ParsePrice(s.Stats.AskPrice, s.QuotePrecision)
}

func (s Symbol) GetPricesTxt() string {
	return fmt.Sprintf(`
%s:
قیمت فروش: %s %s
قیمت خرید: %s %s
`, s.FaBaseAsset, s.GetAsk(), s.FaQuoteAsset, s.GetBid(), s.FaQuoteAsset)
}

func (s Symbol) GetPricesWith24chTxt() string {
	return fmt.Sprintf(`
%s: (%.2f%s)
قیمت فروش: %s %s
قیمت خرید: %s %s
`, s.FaBaseAsset, s.Stats.Two4HCh, "%", s.GetAsk(), s.FaQuoteAsset, s.GetBid(), s.FaQuoteAsset)
}
