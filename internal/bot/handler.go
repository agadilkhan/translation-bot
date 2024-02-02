package bot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!translate help" {
		s.ChannelMessageSend(m.ChannelID, "`Usage: !translate <target_language_code> <text>`")
		return
	}
}
