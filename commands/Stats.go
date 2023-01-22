package commands

import (
	"ds-bot/tmp"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
)

// https://media.giphy.com/media/3y0oCOkdKKRi0/giphy.gif loading gif
//func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
//	if m.Author.Bot {
//		return
//	}
//
//	args := strings.Split(strings.TrimPrefix(m.Content, botPrefix), " ")
//	command := args[0]
//
//	// get all messages
//	// channel ranking
//	// user ranking
//	// compare
//
//	if command == "stats" {
//		if args[1] == "ever" {
//			s.ChannelMessageSend(m.ChannelID, "Drumroll...")
//			guild, err := s.State.Guild(m.GuildID)
//			if err != nil {
//				fmt.Println(err)
//			}
//
//			currentChannel := guild.Channels[0]
//			allChannelMessages := currentChannel.Messages
//			stats := make(map[int]int)
//
//			for i := 0; i < currentChannel.MessageCount; i++ {
//				currentMessage := allChannelMessages[i]
//				currentMessageAuthorID, _ := strconv.Atoi(currentMessage.ID)
//				stats[currentMessageAuthorID]++
//			}
//			var messageInfo, userID int
//			for i := 0; i < len(guild.Members); i++ {
//				messageInfo, userID = tmp.Max(stats, guild.Members)
//			}
//			userId := strconv.Itoa(userID)
//			u, _ := s.User(userId)
//			s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("The most effective user of all times and peoples: ", fmt.Sprintf("%d messages for all time! %v", messageInfo, u.Mention())).Return())
//
//		} else {
//
//		}
//	}
//}

// WordStats is func to find all word appearing in given range
// .word x "n" --> find all appearing of word x in n last messages
func WordStats(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	args := strings.Split(strings.TrimPrefix(m.Content, botPrefix), " ")
	command := args[0]

	if command == "word" {
		if args[1] != "ever" {
			guild, err := s.State.Guild(m.GuildID)
			if err != nil {
				log.Println(err)
			}
			channel := guild.Channels[0]

			// User -> times args[1] is in messages from that user
			arr := make(map[*discordgo.User]int)
			var users []*discordgo.User

			ran, _ := strconv.Atoi(args[1])

			for i := 0; i < ran; i++ {
				if channel.MessageCount >= ran {
					mes := channel.Messages[channel.MessageCount-i]
					if strings.Contains(args[1], mes.Content) {
						arr[mes.Author]++
						users = append(users, mes.Author)
					}
				}
			}

		} else {
			guild, err := s.State.Guild(m.GuildID)
			if err != nil {
				log.Println(err)
			}
			channel := guild.Channels[0]

			// User -> times args[1] is in messages from that user
			arr := make(map[*discordgo.User]int)
			var users []*discordgo.User

			// find all mentions of that word
			for i := 0; i < channel.MessageCount; i++ {
				mes := channel.Messages[channel.MessageCount-i]
				if strings.Contains(args[1], mes.Content) {
					arr[mes.Author]++
					users = append(users, mes.Author)
				}
			}

			// find 5 max mentions of this word
			// user -> mention
			// max mentions
			max := arr[users[0]]
			var winUsers []*discordgo.User
			var countRep []int
			var user *discordgo.User
			var ind int

			for i := 0; i < 5; i++ {
				for a, v := range users {
					if arr[users[a]] > max {
						user = v
						max = arr[users[a]]
						ind = a
					}
				}
				winUsers = append(winUsers, user)
				countRep = append(countRep, max)
				arr[users[ind]] = 0
			}

			s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedInfoMessage(fmt.Sprintf("Word {%s} usage graph", args[1]), "").Return())
			var link = fmt.Sprintf("<img src=`https://quickchart.io/chart?c={type:'radar',data:{labels:['%v','%v','%v','%v','%v'],datasets:[{label:'%s',data:[%d,%d,%d,%d,%d]}`>", users[0], users[1], users[2], users[3], users[4], args[1], countRep[0], countRep[1], countRep[2], countRep[3], countRep[4])
			s.ChannelMessageSend(m.ChannelID, link)
		}
	}
}

/*
	guildId, _ := strconv.Atoi(m.GuildID)
	guild := s.State.Guilds[guildId]
	roles := guild.Roles

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("This server has roles", fmt.Sprintf("%v", roles)).Return())
*/
