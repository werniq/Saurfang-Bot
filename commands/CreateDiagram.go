package commands

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"strings"
)

type URIData struct {
	DiagramType string   `json:"type"`
	Options     *Options `json:"options"`
	Data        []*Data  `json:"data"`
}

type Data struct {
	Labels   []any     `json:"labels"`
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	BorderColor string `json:"borderColor"`
	BorderWidth string `json:"borderWidth"`
	Fill        bool   `json:"fill"`
	Data        []any  `json:"data"`
}

type Options struct {
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	BackgroundColor string `json:"backgroundColor"`
	Format          string `json:"format"`
}

func CreateDiagram(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if m.Content == "" {
		return
	}

	args := strings.Split(m.Content, " ")
	comm := args[0]
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = nil
	}

	if comm == "diagram" && args != nil {
		if len(m.Attachments) != 0 {
			for _, attachment := range m.Attachments {
				resp, err := s.Client.Get(attachment.URL)
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}

				var UriData URIData
				err = json.Unmarshal(body, &UriData)
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}
				uri, err := createDiagramUrl(UriData)
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}

				s.ChannelMessageSend(m.ChannelID, uri)
			}
		}
	}
}

func createDiagramUrl(Data URIData) (string, error) {
	uri := "https://quickchart.io/chart?chart={type:"
	uri += "'" + Data.DiagramType + "',"
	uri += "data:{labels:["

	for i := 0; i < len(Data.Data); i++ {
		v := Data.Data[i]
		uri += fmt.Sprintf("%d,", v.Labels[i])
	}
	uri += "],datasets:[{"

	for i, v := range Data.Data {
		uri += "'" + fmt.Sprintf("%v", v.Datasets[i].Label) + "',data:["
		for _, d := range v.Datasets[i].Data {
			uri += fmt.Sprintf("%v,", d)
		}
		uri += "]}"
		if Data.Data[i+1] != nil {
			uri += ","
		}
	}

	if Data.Options.Height != 0 && Data.Options.Width != 0 {
		uri += "?height=" + fmt.Sprintf("%d", Data.Options.Height)
		uri += "?width=" + fmt.Sprintf("%d", Data.Options.Width)
	}

	if Data.Options.BackgroundColor != "" {
		uri += "?backgroundColor=" + Data.Options.BackgroundColor
	}

	if Data.Options.Format != "" {
		uri += "?format=" + Data.Options.Format
	}

	return uri, nil
}
