package bot

import (
	"fmt"
	"github.com/agadilkhan/translator-bot/internal/translation"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (b *Bot) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from the bot itself
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

	// parse the translation command arguments
	args, err := parseTranslateCommand(msg)
	if err != nil {
		log.Printf("parseTranslateCommand err: %v", err)
		s.ChannelMessageSend(m.ChannelID, "`Invalid command. Use '!translate help' for usage.`")
		return
	}

	targetLang, text := args[0], args[1]

	// launch a goroutine to handle the translation asynchronously
	go func() {
		res, err := b.TranslationWebApi.Translate(translation.Translation{
			Source:      "auto",
			Destination: targetLang,
			Original:    text,
		})

		b.mu.Lock()
		// remove the user ID from the pending map after translation is done
		delete(b.pending, m.Author.ID)
		b.mu.Unlock()

		if err != nil {
			log.Printf("Translate err: %v", err)
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`Ooops... Some error.`"))
			return
		}

		s.ChannelMessageSend(m.ChannelID, res.Translation)
	}()

}

func parseTranslateCommand(message string) ([]string, error) {
	words := strings.Fields(message)

	// check if the command has at least three words
	if len(words) < 3 {
		return nil, fmt.Errorf("the command must have at least three words")
	}

	targetLang := words[1]

	text := strings.Join(words[2:], " ")

	return []string{targetLang, text}, nil
}
