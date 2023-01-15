package tmp

import (
	"fmt"
	"image/color"
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/image/colornames"
)

var (
	ColorCyan = colornames.Cyan
)

func colorToInt(color color.RGBA) int {
	return 256*256*int(color.R) + 256*int(color.G) + int(color.B)
}

func HasPerm(session *discordgo.Session, user *discordgo.User, channelID string, perm int64) bool {
	perms, err := session.State.UserChannelPermissions(user.ID, channelID)
	if err != nil {
		_, _ = session.ChannelMessageSend(channelID, fmt.Sprintf("Failed to retrieve perms: %s", err.Error()))
		return false
	}
	return perms&perm != 0
}

func FindUser(session *discordgo.Session, mentions []*discordgo.User, arg string) *discordgo.User {
	if len(mentions) > 0 {
		return mentions[0]
	}
	user, _ := session.User(arg)
	return user
}

func CreateEmbedMessage(title string, description string) *EmbMessage {
	return NewMessageEmbed().SetTitle(title).SetDescription(description).SetFooter().SetFooterText("Bot Saurfang ").SetTimestamp(time.Now()).SetColor(colornames.Cyan)
}

type EmbMessage struct {
	em *discordgo.MessageEmbed
}

func (e *EmbMessage) Return() *discordgo.MessageEmbed {
	return e.em
}

func (e *EmbMessage) SetTitle(t string) *EmbMessage {
	e.em.Title = t
	return e
}

func (e *EmbMessage) SetDescription(d string) *EmbMessage {
	e.em.Description = d
	return e
}

func (e *EmbMessage) SetColor(c color.RGBA) *EmbMessage {
	e.em.Color = colorToInt(c)
	return e
}

func (e *EmbMessage) SetTimestamp(t time.Time) *EmbMessage {
	e.em.Timestamp = t.Format(time.RFC3339)
	return e
}

type Footer struct {
	EmbMessage
}

func (f *Footer) SetFooterText(n string) *Footer {
	f.em.Footer.Text = n
	return f
}

func (e *EmbMessage) SetFooter() *Footer {
	e.em.Footer = &discordgo.MessageEmbedFooter{}
	return &Footer{*e}
}

func NewMessageEmbed() *EmbMessage {
	em := &discordgo.MessageEmbed{
		URL:         "",
		Type:        "",
		Title:       "",
		Description: "",
		Timestamp:   "",
		Color:       0,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
		Fields:      nil,
	}
	return &EmbMessage{em: em}
}
