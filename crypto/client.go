package crypto

import (
	"crypto-bot/config"
	"log"
	"net/http"
	"os"
)

func GetCoin() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)

		req.Header.Set("Accepts", "application/json")
		req.Header.Add("X-CMC_PRO_API_KEY", config.LoadConfig().CryptoAPI)
	}
}
