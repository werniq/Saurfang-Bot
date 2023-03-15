package commands

import (
	"crypto/hmac"
	"crypto/sha256"
	"ds-bot/tmp"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	BYBIT_API_KEY    = "s4jHbNdr0bQqZYHZfv"
	BYBIT_API_SECRET = "ljA9MXXFcohtLasjIa6BzszHnkLHpdTVlMY0"
	endpoint         = "/spot/v3/public/quote/ticker/price?symbol="
	BuyingPrice      = map[string]float32{}

	MaticLogo = "https://th.bing.com/th/id/R.7ad2ca3d46e1b51160f446aab3910938?rik=EwVa4B4hnM%2fqpw&pid=ImgRaw&r=0"
)

type ByBitResponse struct {
	CurrentTimestamp time.Time `json:"t"`
	TradingPairName  string    `json:"s"`
	LastTradedPrice  string    `json:"lp"`
	LowestPrice      string    `json:"l"`
	HighestPrice     string    `json:"h"`
	TradingVolume    string    `json:"v"`
}

func TokenPrice(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "" {
		return
	}

	args := strings.Split(strings.TrimPrefix(m.Content, "."), " ")
	command := args[0]
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = nil
	}

	if args != nil && command == "token" {
		if args[0] == "MATICUSDT" {
			BuyingPrice[strings.ToLower(args[0])] = 117.7
			endpoint = fmt.Sprintf(endpoint + args[0])
			fmt.Println(endpoint)
			req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bybit.com%s", endpoint), nil)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Failed to create new request: "+err.Error())
				return
			}
			signature := generateSignature(endpoint, BYBIT_API_KEY, BYBIT_API_SECRET)

			req.Header.Add("api-key", BYBIT_API_KEY)
			req.Header.Add("api-signature-method", "HmacSHA256")
			req.Header.Add("api-signature-version", "2")
			req.Header.Add("api-signature", signature)
			req.Header.Add("Content-Type", "application/json")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Failed executing request: "+err.Error())
				return
			}

			type Result struct {
				Price  string `json:"price"`
				Symbol string `json:"symbol"`
			}

			type Response struct {
				Result     Result                 `json:"result"`
				RetCode    int                    `json:"retCode"`
				RetExtInfo map[string]interface{} `json:"retExtInfo"`
				RetMsg     string                 `json:"retMsg"`
				Time       float64                `json:"time"`
			}

			var r Response
			err = json.NewDecoder(res.Body).Decode(&r)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v", err))
				return
			}

			profitPerc, err := strconv.ParseFloat(r.Result.Price, 32)
			if err != nil {
				fmt.Println(err)
			}
			tokenPrice := BuyingPrice[strings.ToLower(args[0])]
			gain := tokenPrice * float32(profitPerc)

			s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("ByBit Response for "+args[0], fmt.Sprintf(`
				Token name: %s
				Latest Price: %s
				USD amount after selling your MATIC Balance: %f
				Purchase Price: %s
			`,
				r.Result.Symbol,
				r.Result.Price,
				gain,
				"127.5",
			), 0).Build())
		} else {
			endpoint = fmt.Sprintf(endpoint + args[0])
			req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bybit.com%s", endpoint), nil)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Failed to create new request: "+err.Error())
				return
			}
			signature := generateSignature(endpoint, BYBIT_API_KEY, BYBIT_API_SECRET)

			req.Header.Add("api-key", BYBIT_API_KEY)
			req.Header.Add("api-signature-method", "HmacSHA256")
			req.Header.Add("api-signature-version", "2")
			req.Header.Add("api-signature", signature)
			req.Header.Add("Content-Type", "application/json")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Failed executing request: "+err.Error())
				return
			}

			type Result struct {
				Price  string `json:"price"`
				Symbol string `json:"symbol"`
			}

			type Response struct {
				Result     Result                 `json:"result"`
				RetCode    int                    `json:"retCode"`
				RetExtInfo map[string]interface{} `json:"retExtInfo"`
				RetMsg     string                 `json:"retMsg"`
				Time       float64                `json:"time"`
			}
			var r Response
			err = json.NewDecoder(res.Body).Decode(&r)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v", err))
				return
			}

			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(fmt.Sprintf("Price for %s", args[0]), fmt.Sprintf(`
				Symbol: %s
				Price:  %s
			`,
				r.Result.Symbol,
				r.Result.Price), 4).Build())
		}

		//s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage(fmt.Sprintf("ByBit Response for %s", args[0]), fmt.Sprintf(`
		//		Trading Pair Name: %s
		//		Last Traded Price: %s
		//		Highest Price: 	   %s
		//		Lowest Price: 	   %s
		//		Trading Volume:    %s
		//		TimeStamp:		   %s
		//	`,
		//	exchResp.TradingPairName,
		//	exchResp.LastTradedPrice,
		//	exchResp.HighestPrice,
		//	exchResp.LowestPrice,
		//	exchResp.TradingVolume,
		//	exchResp.CurrentTimestamp), 3).Build())
	}
	_, err := getWalletBalance()
	if err != nil {
		log.Printf("Error :%v", err)
	}
}

func getWalletBalance() (interface{}, error) {
	type Result struct {
		Balances []struct {
			Coin   string `json:"coin"`
			CoinId string `json:"coinId"`
			Total  string `json:"total"`
		} `json:"balances"`
	}
	var Response struct {
		Result     Result                 `json:"result"`
		RetCode    int                    `json:"retCode"`
		RetExtInfo map[string]interface{} `json:"retExtInfo"`
		RetMsg     string                 `json:"retMsg"`
		Time       float64                `json:"time"`
	}

	uri := "https://api.bybit.com"
	end := "/spot/v3/private/account"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", uri, end), nil)
	if err != nil {
		log.Println(err)
		return Response, err
	}
	signature := generateSignature(end, BYBIT_API_KEY, BYBIT_API_SECRET)

	req.Header.Add("api-key", BYBIT_API_KEY)
	req.Header.Add("api-signature-method", "HmacSHA256")
	req.Header.Add("api-signature-version", "2")
	req.Header.Add("api-signature", signature)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response, err
	}
	err = json.NewDecoder(res.Body).Decode(&Response)
	if err != nil {
		log.Printf("Here 3")
		return Response, err
	}
	fmt.Println(Response)
	return Response, nil
}

func generateSignature(endpoint, apiKey, secretKey string) string {
	message := fmt.Sprintf("GET%s%d%s", endpoint, time.Now().UnixMilli(), apiKey)

	hmac := hmac.New(sha256.New, []byte(secretKey))
	hmac.Write([]byte(message))

	signature := hex.EncodeToString(hmac.Sum(nil))
	return signature
}
