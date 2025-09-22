package crypto

type PriceConverion struct {
	Data CoinData `json:"data"`
}

type CoinData struct {
	Symbol      string           `json:"symbol"`
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Amount      int              `json:"amount"`
	LastUpdated string           `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

type Quote struct {
	Price       float64 `json:"price"`
	LastUpdated string  `json:"last_updated"`
}
