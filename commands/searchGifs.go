package commands

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type GifSearch struct {
	Data struct {
		Type             string `json:"type"`
		Id               string `json:"id"`
		Url              string `json:"url"`
		Slug             string `json:"slug"`
		BitlyGifUrl      string `json:"bitly_gif_url"`
		BitlyUrl         string `json:"bitly_url"`
		EmbedUrl         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Title            string `json:"title"`
		Rating           string `json:"rating"`
		ContentUrl       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostUrl    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		Images           struct {
			FixedWidthStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_still"`
			PreviewGif struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_gif"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				Url      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_downsampled"`
			Preview struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"preview"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				Url      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_small"`
			Downsized struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				Url      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_downsampled"`
			FixedWidth struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				Url      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width"`
			DownsizedStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_still"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_medium"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"original_mp4"`
			DownsizedLarge struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_large"`
			PreviewWebp struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_webp"`
			Original struct {
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				Url      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"original"`
			OriginalStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"original_still"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_small_still"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				Url      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_small"`
			Looping struct {
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
			} `json:"looping"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"downsized_small"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_small_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_still"`
			FixedHeight struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				Url      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height"`
			WStill struct {
				Url    string `json:"url"`
				Width  string `json:"width"`
				Height string `json:"height"`
			} `json:"480w_still"`
		} `json:"images"`
	} `json:"data"`
	Meta struct {
		Msg        string `json:"msg"`
		Status     int    `json:"status"`
		ResponseId string `json:"response_id"`
	} `json:"meta"`
}

func SearchGifs(s *discordgo.Session, mes *discordgo.MessageCreate) {
	if mes.Author.Bot {
		return
	}

	args := strings.Split(strings.TrimPrefix(mes.Content, botPrefix), " ")

	command := args[0]

	if len(args) > 1 {
		args = args[1:]
	} else {
		args = nil
	}
	if command == "search" && len(command) > 1 {
		err := godotenv.Load(".env")

		giphyToken := os.Getenv("GIPHY_TOKEN")

		if err != nil {
			log.Fatal(err)
		}

		url := "https://api.giphy.com/v1/gifs/random"
		var result GifSearch

		gifKeyword := strings.Join(args[1:], " ")

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			fmt.Println("Error in making a new Request", err)
		}

		query := req.URL.Query()
		query.Add("api_key", giphyToken)
		query.Add("tag", gifKeyword)
		req.URL.RawQuery = query.Encode()
		client := http.Client{}
		res, err := client.Do(req)

		if err != nil {
			fmt.Println("Error in getting a response, ", err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshall JSON", err)
		}

		s.ChannelMessageSend(mes.ChannelID, result.Data.EmbedUrl)
		res.Body.Close()
	}
}
