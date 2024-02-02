package bot

import (
	"github.com/agadilkhan/translation-bot/internal/translation"
)

type Bot struct {
	Token             string
	TranslationWebApi *translation.TranslationWebApi
}

func New(token string) *Bot {
	t := translation.New()

	return &Bot{
		Token:             token,
		TranslationWebApi: t,
	}
}
