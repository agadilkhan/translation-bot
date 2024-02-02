package bot

import (
	"github.com/agadilkhan/translator-bot/internal/translation"
	"sync"
)

type Bot struct {
	Token             string
	TranslationWebApi *translation.TranslationWebAPI
	mu                sync.Mutex
	pending           map[string]struct{}
}

func New(token string) *Bot {
	t := translation.New()

	return &Bot{
		Token:             token,
		TranslationWebApi: t,
		pending:           make(map[string]struct{}),
	}
}
