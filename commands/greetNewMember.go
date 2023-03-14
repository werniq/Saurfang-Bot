package commands

import (
	"ds-bot/tmp"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// greetNewMember basically greets new member
func GreetNewMember(s *discordgo.Session, mes *discordgo.MessageCreate) {
	if mes.Type == 7 {
		s.ChannelMessageSendEmbed(mes.ChannelID, tmp.CreateEmbedMessage("I'm glad to see you here, recruit!", greetMsg, 1).Build())
		s.ChannelMessageSend(mes.ChannelID, fmt.Sprint(forTheHordeLink))
	}
}
