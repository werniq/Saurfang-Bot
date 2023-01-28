package commands

import (
	"ds-bot/tmp"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	forecastURI = "https://th.bing.com/th/id/R.67e4e22a4f6ea615af55373263d9bb08?rik=RB4AVfq3N1jzNQ&pid=ImgRaw&r=0"
)

type Weather struct {
	Longitude int
	Latitude  int
	Timezone  string
	Currently struct {
		Time                 int
		Summary              string
		Icons                string
		NearestStormDistance int
		NearestStormBearing  int
		PrecipType           string
		Ozone                int
	}
	Daily struct {
		Summary     string
		Icon        string
		SunriseTime int
		SunsetTime  int
	}
	Hourly struct {
		Summary string
		Icon    string
	}
}

func GetWeather(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	if command == "weather" {
		if len(args) == 2 {
			longitude, _ := strconv.Atoi(args[0])
			latitude, _ := strconv.Atoi(args[1])
			url := fmt.Sprintf("https://dark-sky.p.rapidapi.com/%d,%d?units=auto&lang=en", longitude, latitude)

			req, _ := http.NewRequest("GET", url, nil)

			req.Header.Add("X-RapidAPI-Key", "0e16774897msh0d7f6465454053ap11a57ejsn3d7990cdb478")
			req.Header.Add("X-RapidAPI-Host", "dark-sky.p.rapidapi.com")

			res, _ := http.DefaultClient.Do(req)

			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			var weather Weather
			json.Unmarshal(body, &weather)

			s.ChannelMessageSend(m.ChannelID, forecastURI)
			s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(fmt.Sprintf("Weather forecast"), fmt.Sprintf(
				`
						
						Current weather info: %v
						
						Hourly weather forecats: %v
						
						Timezone: %s
						
						`, weather.Currently, weather.Hourly, weather.Timezone), 4).Build())
		} else {
			s.ChannelMessageSend(m.ChannelID, ".weather [longitude] [latitude]")
		}
	}
}
