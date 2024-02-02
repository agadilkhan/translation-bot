package app

import (
	"fmt"
	"github.com/agadilkhan/translation-bot/internal/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	discordToken := os.Getenv("DISCORD_TOKEN")

	b := bot.New(discordToken)

	// create a new Discord bot session using the provided bot token
	dg, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// register an event handler
	dg.AddHandler(b.MessageHandler)

	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// open a websocket connection
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer dg.Close()

	// wait here until CTRL+C or other term signal is received
	fmt.Println("Bot is now running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-c
}
