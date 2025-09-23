package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	Key string
}

func NewClient(key string) *Client {
	return &Client{Key: key}
}

func (c *Client) GetCryptoPrice(ctx context.Context, symbol string) (*CoinData, error) {
	url := fmt.Sprintf(
		"https://pro-api.coinmarketcap.com/v1/tools/price-conversion?symbol=%s&amount=1&convert=USD",
		symbol,
	)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-CMC_PRO_API_KEY", c.Key)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data PriceConverion

	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return &data.Data, nil

}
