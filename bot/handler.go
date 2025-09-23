package bot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/abdigaliarslan/crypto-bot/crypto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleStart(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := "Hello! Welcome to the Crypto Bot. Use /price <symbol> to get the latest price of a cryptocurrency."
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
}

func handlePrice(update tgbotapi.Update, b *tgbotapi.BotAPI, cryptoAPI *crypto.Client) {

	arg := update.Message.CommandArguments()
	if strings.TrimSpace(arg) == "" {
		b.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide a cryptocurrency symbol. Usage: /price <symbol>"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coin, err := cryptoAPI.GetCryptoPrice(ctx, arg)
	if err != nil {
		b.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Error fetching price. Please ensure the symbol is correct."))
		return
	}

	usd, ok := coin.Quote["USD"]
	if !ok {
		b.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "USD quote not available"))
		return
	}

	reply := fmt.Sprintf("The current price of %s is $%.2f (last updated: %s)", coin.Name, usd.Price, usd.LastUpdated)
	b.Send(tgbotapi.NewMessage(update.Message.Chat.ID, reply))
}
