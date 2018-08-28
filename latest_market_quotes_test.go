package coinmarketcap

import (
	"encoding/json"
	"testing"
)

func TestLatestMarketQuotesJson(t *testing.T) {
	data := []byte(`
{
    "data":{
        "1":{
            "id":1,
            "name":"Bitcoin",
            "symbol":"BTC",
            "slug":"bitcoin",
            "circulating_supply":17199862,
            "total_supply":17199862,
            "max_supply":21000000,
            "date_added":"2013-04-28T00:00:00.000Z",
            "num_market_pairs":331,
            "cmc_rank":1,
            "last_updated":"2018-08-09T21:56:28.000Z",
            "quote":{
                "USD":{
                    "price":6602.60701122,
                    "volume_24h":4314444687.5194,
                    "percent_change_1h":0.988615,
                    "percent_change_24h":4.37185,
                    "percent_change_7d":-12.1352,
                    "market_cap":113563929433.21645,
                    "last_updated":"2018-08-09T21:56:28.000Z"
                }
            }
        }
    },
    "status":{
        "timestamp":"2018-06-02T22:51:28.209Z",
        "error_code":0,
        "error_message":"",
        "elapsed":10,
        "credit_count":1
    }
}`)

	r := &GetLatestMarketQuotesResponse{}
	err := json.Unmarshal(data, r)
	if err != nil {
		t.Fatal(err)
	}
}
