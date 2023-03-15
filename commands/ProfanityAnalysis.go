package commands

import (
	"ds-bot/tmp"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
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
	command := args[0]
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = nil
	}

	if args != nil && command == botPrefix+"profanity" {
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

		fi, err := os.ReadFile("profanity-list.txt")
		if err != nil {
			Logger().Printf("Error opening profanity-list.txt: %v\n", err)
			return
		}

		statsIds := []string{}
		antiStatsIds := []string{}

		statistics := make(map[string]int)
		antiStatistics := make(map[string]int)

		for row.Next() {
			var id int
			var content string
			var authorId string
			if err = row.Scan(&id, &content, &authorId); err != nil {
				Logger().Printf("Error scanning rows: %v\n", err)
				return
			}

			if profanityCheck(content, fi) {
				if elementInArray(antiStatsIds, authorId) {
					antiStatistics[authorId]++
					antiStatsIds = append(antiStatsIds, authorId)
				}
			} else {
				if elementInArray(statsIds, authorId) {
					statistics[authorId]++
					statsIds = append(statsIds, authorId)
				}
			}
		}

		// statistic for polite users
		id, score := findMax(statistics, statsIds)
		user, err := s.User(id)
		if err != nil {
			Logger().Printf("Error retrieving user from session: %v", err)
			return
		}

		s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("The most polite person on this server..", fmt.Sprintf(`
				%v with the score of %d`, user.Mention(), score), 3).Build())

		// statistics for "bad" users
		id, score = findMax(antiStatistics, antiStatsIds)
		user, err = s.User(id)
		if err != nil {
			Logger().Printf("Error retrieving user from session: %v", err)
			return
		}

		s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("The most motherfucking user on this motherfucking server...",
			fmt.Sprintf(`
				%v with the score of %d
			`, user.Mention(), score), 3).Build())

	}
}

func elementInArray(arr []string, elem string) bool {
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == elem {
			return true
		}
	}
	return false
}

func findMax(stats map[string]int, ids []string) (string, int) {
	max := stats[ids[0]]
	id := ids[0]
	for i := 0; i <= len(ids)-1; i++ {
		if stats[ids[i]] > max {
			max = stats[ids[i]]
			id = ids[i]
		}
	}
	return id, max
}

// profanityCheck returns bool, if m.Content contains profanity
func profanityCheck(s string, file []byte) bool {
	fi := string(file)
	arr := strings.Split(string(fi), " ")
	for i := 0; i < len(arr)-1; i++ {
		if strings.Contains(s, arr[i]) {
			return false
		}
	}
	return true
}
