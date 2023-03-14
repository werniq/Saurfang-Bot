package commands

import "github.com/bwmarrin/discordgo"

func ChangeStatus(s *discordgo.Session, event *discordgo.Event) {
	s.UpdateListeningStatus("Horde theme song")
}
