package commands

import (
	"bytes"
	"ds-bot/tmp"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ConnectEthereumWallet(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	if command == "get-my-balance" {
		address := args[0]
		if len(address) != 42 {
			s.ChannelMessageSend(m.ChannelID, "Ethereum address should be of length 42 characters")
			return
		} else if !strings.Contains(address[:2], "0x") {
			s.ChannelMessageSend(m.ChannelID, "Please, enter valid wallet address.")
			return
		}
		url := fmt.Sprintf("https://deep-index.moralis.io/api/v2/%s/balance?chain=eth", address)

		bytesObj := []byte(fmt.Sprintf(`{"address":%s}`, address))
		payload := bytes.NewBuffer(bytesObj)

		req, err := http.NewRequest("GET", url, payload)
		if err != nil {
			log.Fatal(err)
			return
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Add("X-API-KEY", "cxFW6hUQTIbBA30UB3CKaKrHBwqqC26G1806jRNhOXwiAJDSYanuYN2AHMCH5uKw")

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))
		s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("ETH BALANCE", fmt.Sprintf("Balance of %s is %d WEI", address, body), 0).Build())
	}
}
