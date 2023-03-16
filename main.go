package main

import (
	"ds-bot/commands"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func colorToInt(color color.RGBA) int {
	return 256*256*int(color.R) + 256*int(color.G) + int(color.B)
}

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

	bot.AddHandler(commands.SaveMessage)
	bot.AddHandler(commands.GetWeather)
	bot.AddHandler(commands.Help)
	bot.AddHandler(commands.ChangeStatus)
	bot.AddHandler(commands.SearchGifs)
	bot.AddHandler(commands.FindLyrics)
	bot.AddHandler(commands.ProfanityAnalysis)
	bot.AddHandler(commands.GreetNewMember)
	bot.AddHandler(commands.ListRules)
	bot.AddHandler(commands.Ban)
	bot.AddHandler(commands.Unban)
	bot.AddHandler(commands.FindLyrics)
	bot.AddHandler(commands.TokenPrice)
	//bot.AddHandler(commands.Stats)
	bot.AddHandler(commands.WordStats)

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
