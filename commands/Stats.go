package commands

import (
	"context"
	"database/sql"
	"ds-bot/tmp"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func WordStats(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	if command == "word" {
		arg := args[0]

		var err error
		//messageCount := 0

		//currentChannel, _ := s.Channel(m.ChannelID)
		db, err := OpenDb()
		if err != nil {
			log.Fatalln(err)
		}

		var auth *discordgo.User
		//var count int
		row := db.QueryRowContext(context.Background(), `SELECT * FROM message`)
		if row.Err() == sql.ErrNoRows {
			s.ChannelMessageSend(m.ChannelID, "No statistic for given word :c")
			return
		}
		var messages []Message
		//messages =
		row.Scan(&messages)
		for ind, val := range messages {
			if !strings.Contains(val.Content, arg) {
				messages[ind] = messages[len(messages)-1]
				messages = messages[:len(messages)-1]
			}
		}
		max := 0
		user_id := "a"
		for i := 0; i < len(messages)-1; i++ {
			author := messages[i].Author
			counter := 0
			for a := 0; a < len(messages)-1; a++ {
				if messages[a].Author == author {
					counter++
				}
			}
			if counter > max {
				max = counter
				user_id = author.UserID
			}
		}
		auth, _ = s.User(user_id)
		fmt.Println(auth, max)
		s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(fmt.Sprintf("Statistic for usage of word %s", arg), fmt.Sprintf(
			`%v User have used word |%s| [%d] times!`, auth.Mention(), arg, max), 1).Build())
	}
}
