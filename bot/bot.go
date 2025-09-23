package bot

import (
	"log"
	"time"

	"github.com/abdigaliarslan/crypto-bot/config"
	"github.com/abdigaliarslan/crypto-bot/crypto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string) {
	cfg := config.LoadConfig()
	cryptoAPI := crypto.NewClient(cfg.CryptoAPI)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		time.Sleep(100 * time.Millisecond)
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		switch update.Message.Command() {
		case "start":
			handleStart(update, bot)
		case "price":
			handlePrice(update, bot, cryptoAPI)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command. Please use /start or /price <symbol>.")
			bot.Send(msg)
		}
	}

}
