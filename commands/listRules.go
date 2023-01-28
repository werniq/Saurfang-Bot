package commands

import (
	"ds-bot/tmp"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// listRules shows reasons, which can be used to ban the user
func ListRules(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	args := strings.Split(strings.TrimPrefix(m.Content, botPrefix), " ")
	command := args[0]

	if command == "listrules" {
		s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("Rules", rules, 3).Build())
	}
}
