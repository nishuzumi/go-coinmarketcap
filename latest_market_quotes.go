package coinmarketcap

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type GetLatestMarketQuotesService struct {
	c       *Client
	id      string
	symbol  string
	convert string
}

func (s *GetLatestMarketQuotesService) ID(ids ...int) *GetLatestMarketQuotesService {
	s.id = strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1)
	return s
}

func (s *GetLatestMarketQuotesService) Symbol(symbols ...string) *GetLatestMarketQuotesService {
	s.symbol = strings.Replace(strings.Trim(fmt.Sprint(symbols), "[]"), " ", ",", -1)
	return s
}

func (s *GetLatestMarketQuotesService) Convert(currencies ...Currency) *GetLatestMarketQuotesService {
	s.convert = strings.Replace(strings.Trim(fmt.Sprint(currencies), "[]"), " ", ",", -1)
	return s
}

func (s *GetLatestMarketQuotesService) Do(ctx context.Context, opts ...RequestOption) (res *GetLatestMarketQuotesResponse, err error) {
	r := &Request{
		Method:   "GET",
		Endpoint: "/v1/cryptocurrency/quotes/latest",
	}

	m := Params{
		"id":      s.id,
		"symbol":  s.symbol,
		"convert": s.convert,
	}
	r.SetParams(m)

	data, err := s.c.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetLatestMarketQuotesResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return
}

type GetLatestMarketQuotesResponse struct {
	Data map[string]struct {
		ID                int       `json:"id"`
		Name              string    `json:"name"`
		Symbol            string    `json:"symbol"`
		Slug              string    `json:"slug"`
		CirculatingSupply int       `json:"circulating_supply"`
		TotalSupply       int       `json:"total_supply"`
		MaxSupply         int       `json:"max_supply"`
		DateAdded         time.Time `json:"date_added"`
		NumMarketPairs    int       `json:"num_market_pairs"`
		CmcRank           int       `json:"cmc_rank"`
		LastUpdated       time.Time `json:"last_updated"`
		Quote             map[string]struct {
			Price            decimal.Decimal
			Volume24h        decimal.Decimal `json:"volume_24h"`
			PercentChange1h  decimal.Decimal `json:"percent_change_1h"`
			PercentChange24h decimal.Decimal `json:"percent_change_24h"`
			PercentChange7d  decimal.Decimal `json:"percent_change_7d"`
			MarketCap        decimal.Decimal `json:"market_cap"`
			LastUpdated      time.Time       `json:"last_updated"`
		} `json:"quote"`
	} `json:"data"`
	Status Status `json:"status"`
}
