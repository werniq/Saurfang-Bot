package tmp

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/image/colornames"
)

const (
	//ColorCyan      = colornames.Cyan
	TypeInfo Type = iota
	TypeSuccess
	TypeWarning
	TypeError
)

type Type uint8

type Builder struct {
	em *discordgo.MessageEmbed
}

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

func colorFromType(t Type) color.RGBA {
	var c color.RGBA
	switch t {
	case TypeInfo:
		c = colornames.Crimson
	case TypeSuccess:
		c = colornames.Lemonchiffon
	case TypeError:
		c = colornames.Peachpuff
	case TypeWarning:
		c = colornames.Lightgoldenrodyellow
	case 4:
		c = colornames.Goldenrod
	case 5:
		c = colornames.Blueviolet
	}
	return c
}

func (b *Builder) Build() *discordgo.MessageEmbed {
	return b.em
}

func (b *Builder) AddFields(fs ...*discordgo.MessageEmbedField) *Builder {
	for _, f := range fs {
		b.em.Fields = append(b.em.Fields, f)
	}
	return b
}

func (b *Builder) Fields(fs ...*discordgo.MessageEmbedField) *Builder {
	b.em.Fields = fs
	return b
}

func (b *Builder) Title(t string) *Builder {
	b.em.Title = t
	return b
}

func (b *Builder) Description(d string) *Builder {
	b.em.Description = d
	return b
}

func (b *Builder) Color(c color.RGBA) *Builder {
	b.em.Color = colorToInt(c)
	return b
}

func (b *Builder) Timestamp(t time.Time) *Builder {
	b.em.Timestamp = t.Format(time.RFC3339)
	return b
}

func (b *Builder) URL(u string) *Builder {
	b.em.URL = u
	return b
}

func (b *Builder) Type(t discordgo.EmbedType) *Builder {
	b.em.Type = t
	return b
}

type FooterBuilder struct {
	Builder
}

func (f *FooterBuilder) Text(n string) *FooterBuilder {
	f.em.Footer.Text = n
	return f
}

func (f *FooterBuilder) Icon(url string) *FooterBuilder {
	f.em.Footer.IconURL = url
	return f
}

func (f *FooterBuilder) ProxyURL(url string) *FooterBuilder {
	f.em.Image.ProxyURL = url
	return f
}

func (b *Builder) Footer() *FooterBuilder {
	b.em.Footer = &discordgo.MessageEmbedFooter{}
	return &FooterBuilder{*b}
}

type ImageBuilder struct {
	Builder
}

func (i *ImageBuilder) URL(url string) *ImageBuilder {
	i.em.Image.URL = url
	return i
}

func (i *ImageBuilder) ProxyURL(url string) *ImageBuilder {
	i.em.Image.ProxyURL = url
	return i
}

func (i *ImageBuilder) Width(width int) *ImageBuilder {
	i.em.Image.Width = width
	return i
}

func (i *ImageBuilder) Height(height int) *ImageBuilder {
	i.em.Image.Height = height
	return i
}

func (i *ImageBuilder) Set(img *discordgo.MessageEmbedImage) *ImageBuilder {
	i.em.Image = img
	return i
}

func (b *Builder) Image() *ImageBuilder {
	b.em.Image = &discordgo.MessageEmbedImage{}
	return &ImageBuilder{*b}
}

type ThumbnailBuilder struct {
	Builder
}

func (t *ThumbnailBuilder) URL(url string) *ThumbnailBuilder {
	t.em.Thumbnail.URL = url
	return t
}

func (t *ThumbnailBuilder) ProxyURL(url string) *ThumbnailBuilder {
	t.em.Thumbnail.ProxyURL = url
	return t
}

func (t *ThumbnailBuilder) Width(width int) *ThumbnailBuilder {
	t.em.Thumbnail.Width = width
	return t
}

func (t *ThumbnailBuilder) Height(height int) *ThumbnailBuilder {
	t.em.Thumbnail.Height = height
	return t
}

func (b *Builder) Thumbnail() *ThumbnailBuilder {
	b.em.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	return &ThumbnailBuilder{*b}
}

type VideoBuilder struct {
	Builder
}

func (v *VideoBuilder) URL(s string) *VideoBuilder {
	v.em.Video.URL = s
	return v
}

func (v *VideoBuilder) Width(width int) *VideoBuilder {
	v.em.Video.Width = width
	return v
}

func (v *VideoBuilder) Height(height int) *VideoBuilder {
	v.em.Video.Height = height
	return v
}

func (b *Builder) Video() *VideoBuilder {
	b.em.Video = &discordgo.MessageEmbedVideo{}
	return &VideoBuilder{*b}
}

type ProviderBuilder struct {
	Builder
}

func (p *ProviderBuilder) Name(s string) *ProviderBuilder {
	p.em.Provider.Name = s
	return p
}

func (p *ProviderBuilder) URL(url string) *ProviderBuilder {
	p.em.Provider.URL = url
	return p
}

func (b *Builder) Provider() *ProviderBuilder {
	b.em.Provider = &discordgo.MessageEmbedProvider{}
	return &ProviderBuilder{*b}
}

type AuthorBuilder struct {
	Builder
}

func (a *AuthorBuilder) Name(n string) *AuthorBuilder {
	a.em.Author.Name = n
	return a
}

func (a *AuthorBuilder) Icon(url string) *AuthorBuilder {
	a.em.Author.IconURL = url
	return a
}

func (a *AuthorBuilder) ProxyIcon(url string) *AuthorBuilder {
	a.em.Author.ProxyIconURL = url
	return a
}

func (a *AuthorBuilder) URL(url string) *AuthorBuilder {
	a.em.Author.URL = url
	return a
}

func (b *Builder) Author() *AuthorBuilder {
	b.em.Author = &discordgo.MessageEmbedAuthor{}
	return &AuthorBuilder{*b}
}

type FieldBuilder struct {
	fieldNum int
	Builder
}

func (f *FieldBuilder) Value(v string) *FieldBuilder {
	f.em.Fields[f.fieldNum].Value = v
	return f
}

func (f *FieldBuilder) Name(n string) *FieldBuilder {
	f.em.Fields[f.fieldNum].Name = n
	return f
}

func (f *FieldBuilder) Inline(b bool) *FieldBuilder {
	f.em.Fields[f.fieldNum].Inline = b
	return f
}

func (b *Builder) WithField() *FieldBuilder {
	n := len(b.em.Fields)
	b.em.Fields = append(b.em.Fields, &discordgo.MessageEmbedField{})
	return &FieldBuilder{n, *b}
}

func NewBuilder() *Builder {
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
	return &Builder{em: em}
}

func CreateEmbedMessage(title string, description string, t Type) *Builder {
	return NewBuilder().Title(title).Description(description).Thumbnail().URL("https://th.bing.com/th/id/R.2bae4a2e154d7f42ab34a1cf5cea45db?rik=9ePfQ2UWbcQaow&pid=ImgRaw&r=0").Height(50).Width(50).Footer().Text("Bot Saurfang v1.0.0").Timestamp(time.Now()).Color(colorFromType(t))
}

func Max(arr map[int]int, members []*discordgo.Member) (int, int) {
	id, _ := strconv.Atoi(members[0].User.ID)
	max := arr[id]
	userID := 0
	returnUserID := 0
	for i := 0; i < len(members)-1; i++ {
		userID, _ = strconv.Atoi(members[i].User.ID)
		if arr[userID] > max {
			max = arr[userID]
			returnUserID = userID
		}
	}
	return max, returnUserID
}
