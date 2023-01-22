package commands

import (
	"ds-bot/tmp"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Ban(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	args := strings.Split(strings.TrimPrefix(m.Content, botPrefix), " ")
	command := args[0]
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = nil
	}
	if command == "ban" {
		if tmp.HasPerm(s, m.Author, m.ChannelID, discordgo.PermissionBanMembers) {
			if len(args) < 2 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "You should provide reasons for banning: $ban {user} {reason}")
				return
			}
			u := tmp.FindUser(s, m.Mentions, args[0])
			if u == nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "That user is not in the server.")
				return
			}
			err := s.GuildBanCreateWithReason(m.GuildID, u.ID, strings.Join(args[1:], " "), 1)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error banning user: %s", err.Error()))
				return
			}
			title := "User successfully banned"
			desc := fmt.Sprintf("%v has been unbanned. Welcome back!", u.Mention())

			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(title, desc).Return())
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You have no permission")
		}
	}
}
