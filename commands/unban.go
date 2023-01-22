package commands

import (
	"ds-bot/tmp"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Unban(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	args := strings.Split(strings.TrimPrefix(m.Content, botPrefix), " ")
	command := args[0]
	if len(args) <= 2 {
		args = args[1:]
	} else {
		args = nil
	}

	if command == "unban" {
		if tmp.HasPerm(s, m.Author, m.ChannelID, discordgo.PermissionBanMembers) {
			var u *discordgo.User
			if m.ReferencedMessage != nil {
				u = m.ReferencedMessage.Author
			}

			if u == nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "That user was never in the server.")
				return
			}
			err := s.GuildBanDelete(m.GuildID, u.ID)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error unbanning user: %s", err.Error()))
				return
			}
			title := "User successfully unbanned"
			desc := fmt.Sprintf("User %v has unbanned %v! Welcome back!", m.Author, u.Mention())

			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(title, desc).Return())
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You have no permission")
		}
	}
}
