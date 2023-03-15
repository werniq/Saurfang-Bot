package commands

import (
	"ds-bot/tmp"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"strings"
)

func FindLyrics(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	if command == ".lyrics" && args != nil {
		uri := generateUri(args)
		req, _ := http.NewRequest("GET", uri, nil)

		req.Header.Add("X-RapidAPI-Key", "0e16774897msh0d7f6465454053ap11a57ejsn3d7990cdb478")
		req.Header.Add("X-RapidAPI-Host", "genius-song-lyrics1.p.rapidapi.com")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Error executing request: %v", err)
			return
		}
		type Result struct {
			Type         string `json:"_type"`
			Artist       string `json:"artist_names"`
			Title        string `json:"full_title"`
			ThumbnailImg string `json:"header_image_thumbnail_url"`
			Url          string `json:"url"`
		}

		type Hits struct {
			Result Result `json:"result"`
		}

		type Response struct {
			Hits []Hits `json:"hits"`
		}

		var response Response
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			log.Printf("Error decoding response body: %v", err)
			return
		}
		for i := 0; i < len(response.Hits); i++ {
			s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(response.Hits[i].Result.Title, fmt.Sprintf(`
				Song Author: %s,
				%s
				%s
		`,
				response.Hits[i].Result.Artist,
				response.Hits[i].Result.ThumbnailImg,
				response.Hits[i].Result.Url,
			), 5).Build())
		}
	}
}

func generateUri(arr []string) string {
	uri := "https://genius-song-lyrics1.p.rapidapi.com/search/?q="
	for i := 0; i <= len(arr)-1; i++ {
		uri += arr[i] + "%20"
	}
	uri += "&per_page=5&page=1"
	return uri
}
