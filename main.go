package main

import (
	"github.com/agadilkhan/translation-bot/bot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	discordToken := os.Getenv("DISCORD_TOKEN")

	b := bot.New(discordToken)
	b.Run()
}
