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
	ID        string
	Content   string
	Author    Author
	Timestamp time.Time
}

type Author struct {
	ID       int    `json:"id"`
	UserID   string `json:"userID"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
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

	author := Author{
		UserID:   m.Author.ID,
		Email:    m.Author.Email,
		Username: m.Author.Username,
		Avatar:   m.Author.Avatar,
	}

	stmt := `SELECT (email, username, avatar, user_id) FROM author WHERE username = $1`
	row := db.QueryRowContext(ctx, stmt, author.Username)
	if row.Err() != sql.ErrNoRows {
		stmt = `INSERT INTO author(user_id, email, username, avatar) VALUES($1, $2, $3, $4)`

		row := db.QueryRowContext(ctx, stmt, author.UserID, author.Email, author.Username, author.Avatar)
		if row.Err() != nil {
			log.Fatal(row.Err())
		}
	}

	mes := Message{
		ID:      m.ID,
		Content: m.Content,
		Author: Author{
			UserID:   author.UserID,
			Email:    author.Email,
			Username: author.Username,
			Avatar:   author.Avatar,
		},
		Timestamp: time.Now(),
	}

	stmt = `INSERT INTO message(content,  timestamp, author_id, mess_id) VALUES($1, $2, $3, $4)`

	row = db.QueryRowContext(ctx, stmt, mes.Content, time.Now(), author.ID, mes.ID)
	if row.Err() != nil {
		log.Fatal(row.Err())
	}

}
