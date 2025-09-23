package main

import (
	"log"

	"github.com/abdigaliarslan/crypto-bot/bot"
	"github.com/abdigaliarslan/crypto-bot/config"
)

func main() {
	cfg := config.LoadConfig()
	log.Println("Bot is starting...")

	bot.StartBot(cfg.TelegramAPI)
}
