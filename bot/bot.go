package bot

import (
	"fmt"
	"github.com/agadilkhan/translation-bot/webapi"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

type Bot struct {
	Token             string
	TranslationWebApi *webapi.TranslationWebApi
}

func New(token string) *Bot {
	t := webapi.New()

	return &Bot{
		Token:             token,
		TranslationWebApi: t,
	}
}

func (b *Bot) Run() {
	// create a new Discord bot session using the provided bot token
	dg, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// register an event handler
	dg.AddHandler(messageHandler)

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
