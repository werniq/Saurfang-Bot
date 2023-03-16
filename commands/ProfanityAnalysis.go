package commands

import (
	"bufio"
	"ds-bot/tmp"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
	"time"
)

//type Message struct {
//	ID             int       `json:"id"`
//	Content        string    `json:"content"`
//	AuthorID       string    `json:"author_id"`
//	AuthorUsername string    `json:"author_username"`
//	Timestamp      time.Time `json:"timestamp"`
//	MessageId      string    `json:"message_id"`
//}

func ProfanityAnalysis(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "" {
		return
	}

	args := strings.Split(m.Content, " ")

	if args[0] == botPrefix+"profanity" {
		db, err := OpenDb()
		if err != nil {
			Logger().Printf("Error opening database connection: %v\n", err)
			return
		}
		stmt := `SELECT * FROM dsmessages;`
		row, err := db.Query(stmt)
		if err != nil {
			Logger().Printf("Error querying statement: %v\n", err)
			return
		}

		fi, err := os.Open("C:\\Users\\Oleksandr Matviienko\\ds-bot\\ds-bot\\commands\\profanity-list.txt")
		if err != nil {
			Logger().Printf("Error opening profanity-list.txt: %v\n", err)
			return
		}

		scanner := bufio.NewScanner(fi)
		profanityList := []string{}
		for scanner.Scan() {
			profanityList = append(profanityList, scanner.Text())
		}

		//statsId := make(map[string]bool)
		antiStatsIds := make(map[string]bool)

		//var ids []string
		var antiIds []string

		//statistics := make(map[string]int)
		antiStatistics := make(map[string]int)

		for row.Next() {
			var id int
			var content string
			var authorId string
			var authorUsername string
			var timestamp time.Time
			var msgId string
			if err = row.Scan(&id, &content, &authorId, &authorUsername, &timestamp, &msgId); err != nil {
				Logger().Printf("Error scanning rows: %v\n", err)
				return
			}

			if profanityCheck(content, profanityList) {
				fmt.Println("PASSED")
				if !antiStatsIds[authorId] {
					antiStatistics[authorId]++
					antiStatsIds[authorId] = true
					antiIds = append(antiIds, authorId)
					fmt.Println("not-okok")
				} else {
					antiStatistics[authorId]++
				}
			} else {
				fmt.Println("NOT PASSED")
			}
		}

		// statistic for polite users
		//id, score := findMax(statistics, ids)
		//user, err := s.User(id)
		//if err != nil {
		//	Logger().Printf("Error retrieving user from session: %v", err)
		//	return
		//}
		//
		//s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("The most polite person on this server..", fmt.Sprintf(`
		//		%v with the score of %d`, user.Mention(), score), 3).Build())

		// statistics for "bad" users
		id, score := findMax(antiStatistics, antiIds)
		user, err := s.User(id)
		if err != nil {
			Logger().Printf("Error retrieving user from session: %v", err)
			return
		}

		s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("The most motherfucking user on this motherfucking server...",
			fmt.Sprintf(`
				%v with the profanity score of %d
			`, user.Mention(), score), 3).Build())

	}
}

//func elementInArray(arr []string, elem string) bool {
//	for i := 0; i < len(arr); i++ {
//		if arr[i] == elem {
//			return true
//		}
//	}
//	return false
//}

func findMax(stats map[string]int, ids []string) (string, int) {
	max := stats[ids[0]]
	id := ids[0]
	for i := 0; i < len(ids); i++ {
		if stats[ids[i]] > max {
			max = stats[ids[i]]
			id = ids[i]
		}
	}
	return id, max
}

// profanityCheck returns bool, if m.Content contains profanity
func profanityCheck(s string, profanityArray []string) bool {
	for i := 0; i <= len(profanityArray)-1; i++ {
		if strings.Contains(s, profanityArray[i]) {
			return true
		}
	}
	return false
}
