package bot

import (
	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!help translate" {
		s.ChannelMessageSend(m.ChannelID, "Usage: !translate  <target_language_code>  <text>")
		return
	}
}
