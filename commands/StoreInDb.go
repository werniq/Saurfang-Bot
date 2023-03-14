package commands

import (
	"context"
	"database/sql"
	"github.com/bwmarrin/discordgo"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type Message struct {
	ID             string    `json:"id"`
	Content        string    `json:"content"`
	AuthorID       string    `json:"author_id"`
	AuthorUsername string    `json:"author_username"`
	Timestamp      time.Time `json:"timestamp"`
}

func OpenDb() (*sql.DB, error) {
	dsn := os.Getenv("CONN_STR")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return conn, nil
}

func SaveMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := OpenDb()
	if err != nil {
		log.Fatal(err)
	}

	stmt := `INSERT INTO dsmessages(content, author_id, author_username, timestamp, message_id) VALUES($1, $2, $3, $4, $5)`

	row := db.QueryRowContext(ctx, stmt, m.Content, m.Author.ID, m.Author.Username, time.Now(), m.ID)
	if row.Err() != nil {
		log.Fatal(row.Err())
	}
}
