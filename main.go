package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
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

const (
	botPrefix = "$"
	ownerChatId 	 = "1058658554325770320"
	forTheHordeLink = "https://media.giphy.com/media/xThtatRttFzLD9oEtG/giphy.gif"
	message = `
	Hey!
	My name is Saurfang #or Super Bebra # , and I am extremely useful bot!
	I have such commands: 
		$search {keyword} 	  - Send gif, which represents keyword
		$help 			      - List of commands, which this bot wields
		$ban --replyToMessage - Bans author of replied message
		$listReasons		  - Prints all rules, in this server

	`
)

var (
	reasons = map[int]string{
		1:		"Harass, bully, or threat marginalized or vulnerable groups of people",
		2:		"Engage in content manipulation (spamming, subscriber fraud, vote manipulation, or ban evasion)",
		3:		"Post or threaten to post intimate or sexually explicit photos or videos of another person without their consent",
		4:		"Impersonate someone in a misleading way",
		5:		"Label the content and communities improperly (especially graphic content)",
		6: 		"Post suggestive or sexual content that involves minors",
		7: 		"Post illegal content",
		8: 		"Do anything that stops the normal use of this server",
		9: 		"Transmit, distribute, or upload any viruses, worms, or other malware intended to interfere with this servers service",
		10: 	"Use the platform to violate the law or infringe on intellectual and other property right",
		11: 	"Engage in actions that could disrupt, disable, overburden, or impair this servers service",
		12: 	"Attempt to gain access to another user`s account",
		13: 	"Access, search, or collect data from Reddit",
		14: 	"Use the platform in any way that may be abusive or fraudulent",
	}	
)


func main() {
	// 1
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	Token := os.Getenv("DISCORD_TOKEN")

	// 2
	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating a discord Session, ", err)
	}

	// set status
	bot.AddHandler(changeStatus)
	// list all commands
	bot.AddHandler(help)
	// send gif, which represents keyword
	bot.AddHandler(searchGifs)
	// bans user
	bot.AddHandler(ban)
	// say greetengs to new member
	bot.AddHandler(greetNewMember)
	// lists all reasons, using which user can be banned
	bot.AddHandler(listRules)
	// bot.AddHandler(thisServerIsAlwaysAlive)

	err = bot.Open()
	if err != nil {
		fmt.Println("Error opening Discord Session, ", err)
	}
	fmt.Println("The bot is now running. Press CTRL-C to exit.")

	// 5
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func changeStatus(s *discordgo.Session, event *discordgo.Event) {
	fmt.Println("ready function running")
	s.UpdateGameStatus(1, "I Love Hot Milfs")
	// s.UpdateGameStatus()
	
}

func searchGifs(s *discordgo.Session, message *discordgo.MessageCreate) {
	fmt.Println("Gif function running")
	// 1
	err := godotenv.Load(".env")
	giphyToken := os.Getenv("GIPHY_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	

	command := strings.Split(message.Content, " ")
	
	if command[0] == botPrefix + "search" && len(command) > 1 {
		url := "https://api.giphy.com/v1/gifs/random"
		var result GifSearch
		
		gifKeyword := strings.Join(command[1:], " ")
		
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
		
		s.ChannelMessageSend(message.ChannelID, result.Data.EmbedUrl)
		res.Body.Close()
	}
}


// func say(s *discordgo.Session, mes *discordgo.MessageCreate) {
// 	if mes.Author.ID == s.State.User.ID {
// 		return
// 	}
// 	command := strings.Split(mes.Content, " ")
	
// 	if command[0] == botPrefix + "say" && len(mes) > 1 {
// 		// messageToSend := strings.Join()
// 		s.ChannelMessageSend(mes.ChannelID, )
// 	}asdfadf
// }
func help(s *discordgo.Session, mes *discordgo.MessageCreate) {
	// command := strings.Split(m.Content, " ")
	if mes.Content == botPrefix + "help" {
		s.ChannelMessageSend(mes.ChannelID, message)
	}
}

func ban(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Ban function running")
	
	refMes := m.ReferencedMessage
	command := strings.Split(m.Content, " ")
	// botPRefix + ban + "1-14"
	if command[0] == botPrefix + "ban" && len(command) > 1 {	
		if m.Content == botPrefix + "ban " + command[1] {
			banReason, _ := strconv.Atoi(command[1])
			err := s.GuildBanCreateWithReason(m.GuildID, refMes.Author.ID, reasons[banReason], 7)
			
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Something went wrong.. Reported to owner")
				mes :=  fmt.Sprintf("Error banning user %s | in chat %s.", ownerChatId, m.ChannelID)
				s.ChannelMessageSend(ownerChatId, mes)
			}
			s.ChannelMessageSend(m.ChannelID, "User is banned")
			return
		} 
		s.ChannelMessageSend(m.ChannelID, "Please, provide a reason.")	 
	}	
}

func listRules(s *discordgo.Session, m *discordgo.MessageCreate) {
	var reason string
	// command := strings.Split(m.Content, " ")
	if m.Content == botPrefix + "listReasons"  {
		for i := 1; i <= 14; i++ {
			// i = strconv.Atoi(i)
			reason = reasons[i]
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Reason %d: ----> %s", i, reason))
		}
	}
}
// func unban(s *discordgo.Session, m *discordgo.MessageCreate) {

// }

func greetNewMember(s *discordgo.Session, mes *discordgo.MessageCreate) {
	if mes.Type == 7 {
		s.ChannelMessageSend(mes.ChannelID, `
		I'm glad to see you here, recruit!
		Here is some rules, you have to follow, to become a High Overload, Leader, or Warchief:
			Do NOT:
				Harass, bully, or threat marginalized or vulnerable groups of people,
				Engage in content manipulation (spamming, subscriber fraud, vote manipulation, or ban evasion),
				Post or threaten to post intimate or sexually explicit photos or videos of another person without their consent,
				Impersonate someone in a misleading way,
				Label the content and communities improperly (especially graphic content),
		 		Post suggestive or sexual content that involves minors,
		 		Post illegal content,
		 		Do anything that stops the normal use of this server,
		 		Transmit, distribute, or upload any viruses, worms, or other malware intended to interfere with this servers service,
		     	Use the platform to violate the law or infringe on intellectual and other property right,
		     	Engage in actions that could disrupt, disable, overburden, or impair this servers service,
		     	Attempt to gain access to another user's account,
		     	Access, search, or collect data from this server,
		     	Use the platform in any way that may be abusive or fraudulent 
		For The Horde!	
	`)
	s.ChannelMessageSend(mes.ChannelID, fmt.Sprint(forTheHordeLink))
	}
}