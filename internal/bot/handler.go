package bot

import (
	"fmt"
	"github.com/agadilkhan/translation-bot/internal/translation"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (b *Bot) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := m.Content

	if !strings.Contains(msg, "!translate") {
		s.ChannelMessageSend(m.ChannelID, "`Invalid command. Use '!translate help' for usage.`")
		return
	}

	if msg == "!translate help" {
		s.ChannelMessageSend(m.ChannelID, "`Usage: !translate <target_language_code> <text>`")
		return
	}

	args, err := parseTranslateCommand(msg)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "`Invalid command. Use '!translate help' for usage.`")
		return
	}

	targetLang, text := args[0], args[1]

	trans := translation.Translation{
		Source:      "auto",
		Destination: targetLang,
		Original:    text,
	}

	res, err := b.TranslationWebApi.Translate(trans)
	if err != nil {
		log.Printf("Translate err: %v", err)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`Ooops... Some error.`"))
		return
	}

	s.ChannelMessageSend(m.ChannelID, res.Translation)
}

func parseTranslateCommand(message string) ([]string, error) {
	words := strings.Fields(message)

	if len(words) < 3 {
		return nil, fmt.Errorf("the command must have at least three words")
	}

	targetLang := words[1]

	text := strings.Join(words[2:], " ")

	return []string{targetLang, text}, nil
}
