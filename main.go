package main

import (
	"crypto-bot/bot"
	"crypto-bot/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	log.Println("Bot is starting...")
	bot.StartBot(cfg.TelegramAPI)

}
