package commands

import (
	"crypto/hmac"
	"crypto/sha256"
	"ds-bot/tmp"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
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
	EthLogo
)

type ByBitResponse struct {
	CurrentTimestamp time.Time `json:"t"`
	TradingPairName  string    `json:"s"`
	LastTradedPrice  string    `json:"lp"`
	LowestPrice      string    `json:"l"`
	HighestPrice     string    `json:"h"`
	TradingVolume    string    `json:"v"`
}

type Resp struct {
	result []struct {
		Name  string `json:"symbol"`
		Price string `json:"price"`
	} `json:"result"`
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
		BuyingPrice["matic"] = 127.5
		endpoint = fmt.Sprintf(endpoint + args[0])
		fmt.Println(endpoint)
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bybit.com%s", endpoint), nil)
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Failed to create new request: "+err.Error())
			return
		}
		signature := generateSignature(endpoint, BYBIT_API_KEY, BYBIT_API_SECRET, time.Now().UnixMilli())

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
		profitPerc, _ := strconv.Atoi(r.Result.Price)
		var gain float32
		tokenPrice := BuyingPrice[strings.ToLower(args[0])]

		gain = tokenPrice * float32(profitPerc)
		s.ChannelMessageSendEmbed(m.ChannelID, tmp.CreateEmbedMessage("ByBit Response for "+args[0], fmt.Sprintf(`
			Token name: %s
			Latest Price: %s
			Approximate gain: %f
		`,
			r.Result.Symbol,
			r.Result.Price,
			gain,
		), 0).Build())
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
}

func generateSignature(endpoint, apiKey, secretKey string, timestamp int64) string {
	message := fmt.Sprintf("GET%s%d%s", endpoint, timestamp, apiKey)

	hmac := hmac.New(sha256.New, []byte(secretKey))
	hmac.Write([]byte(message))

	signature := hex.EncodeToString(hmac.Sum(nil))
	return signature
}
