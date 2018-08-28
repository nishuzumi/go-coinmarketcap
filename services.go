package coinmarketcap

func (c *Client) NewGetLatestMarketQuotesService() *GetLatestMarketQuotesService {
	return &GetLatestMarketQuotesService{c: c}
}
