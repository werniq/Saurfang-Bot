package commands

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"strconv"
	"strings"
)

func Spam(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "" {
		return
	}

	args := strings.Split(m.Content, " ")
	command := args[0]
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = nil
	}

	if command == botPrefix+"spam" && args != nil {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		str := ""
		for i := 1; i < len(args); i++ {
			str += args[i] + " "
		}
		bebra := os.Getenv("BEBRA")
		channel, err := s.UserChannelCreate(bebra)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "failed creating new channel")
			return
		}

		msg := "I love you, my day and night, my heart and blood, my water and my sun. You are the center of my universe, the reason for my existence"
		for i := 0; i < num; i++ {
			_, err = s.ChannelMessageSend(channel.ID, msg)
		}

	}
}
