package main

import (
	"ds-bot/tmp"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	botPrefix       = "$"
	ownerChatId     = "1058658554325770320"
	forTheHordeLink = "https://media.giphy.com/media/xThtatRttFzLD9oEtG/giphy.gif"
	message         = `
	Hey!
		My name is Saurfang, and I am extremely useful bot!
		I have such commands: 
			$search {keyword} 	  	- Send gif, which represents keyword
			$help 			      	- List of commands, which this bot wields
			$ban	{user} {reason} - Kick out user from server 
			$listrules 		  	    - Prints all rules, in this server
			$unban 					- Unbans user(need to reply to message)
	`
)

func colorToInt(color color.RGBA) int {
	return 256*256*int(color.R) + 256*int(color.G) + int(color.B)
}

var (
	rules = `
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
		Use the platform in any way that may be abusive or fraudulent`
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	Token := os.Getenv("DISCORD_TOKEN")

	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating a discord Session, ", err)
	}

	bot.AddHandler(ban)
	bot.AddHandler(help)
	bot.AddHandler(changeStatus)
	bot.AddHandler(searchGifs)
	bot.AddHandler(greetNewMember)
	bot.AddHandler(listRules)
	bot.AddHandler(unban)

	bot.Identify.Intents = discordgo.IntentsGuildMessages
	bot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	err = bot.Open()
	if err != nil {
		fmt.Println("Error opening Discord Session, ", err)
	}
	fmt.Println("Bot is currently running. CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func searchGifs(s *discordgo.Session, mes *discordgo.MessageCreate) {
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

func changeStatus(s *discordgo.Session, event *discordgo.Event) {
	s.UpdateListeningStatus("Horde theme song")
}

func unban(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	args := strings.Split(strings.TrimPrefix(m.Content, botPrefix), " ")
	command := args[0]
	if len(args) <= 2 {
		args = args[1:]
	} else {
		args = nil
	}

	if command == "unban" {
		if tmp.HasPerm(s, m.Author, m.ChannelID, discordgo.PermissionBanMembers) {
			var u *discordgo.User
			if m.ReferencedMessage != nil {
				u = m.ReferencedMessage.Author
			}

			if u == nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "That user was never in the server.")
				return
			}
			err := s.GuildBanDelete(m.GuildID, u.ID)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error unbanning user: %s", err.Error()))
				return
			}
			title := "User successfully unbanned"
			desc := fmt.Sprintf("User %v has been unbanned! Welcome back!", u.Mention())

			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(title, desc).Return())
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You have no permission")
		}
	}
}

func help(s *discordgo.Session, ms *discordgo.MessageCreate) {
	if ms.Author.Bot {
		return
	}
	if ms.Content == botPrefix+"help" {
		s.ChannelMessageSend(ms.ChannelID, message)
	}
}

func ban(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	if command == "ban" {
		if tmp.HasPerm(s, m.Author, m.ChannelID, discordgo.PermissionBanMembers) {
			if len(args) < 2 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "You should provide reasons for banning: $ban {user} {reason}")
				return
			}
			u := tmp.FindUser(s, m.Mentions, args[0])
			if u == nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "That user is not in the server.")
				return
			}
			err := s.GuildBanCreateWithReason(m.GuildID, u.ID, strings.Join(args[1:], " "), 1)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error banning user: %s", err.Error()))
				return
			}
			title := "User successfully banned"
			desc := fmt.Sprintf("User %v has been banned. Do not repeat his mistakes.", u.Mention())

			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(title, desc).Return())
			fmt.Println("success")
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You have no permission")
		}
	}
}

// listRules shows reasons, which can be used to ban the user
func listRules(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	args := strings.Split(strings.TrimPrefix(m.Content, botPrefix), " ")
	command := args[0]

	if command == "listrules" {
		s.ChannelMessageSend(m.ChannelID, rules)
	}
}

// greetNewMember basically greets new member
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
