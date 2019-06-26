// This package allows for simpler creation of MessageEmbeds,
// similar to how you would do it using JavaScript/Node.js
package embed

import "github.com/bwmarrin/discordgo"

// Embed type for simpler embeds
type Embed struct {
	*discordgo.MessageEmbed
}

// NewEmbed creates a new Embed for use
func NewEmbed() *Embed {
	return &Embed{&discordgo.MessageEmbed{}}
}

// SetTitle sets title of the Embed
func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

// SetColor sets color of the embed
func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

// SetURL sets the URL for the embed
func (e *Embed) SetURL(url string) *Embed {
	e.URL = url
	return e
}

// AddField adds a field to an embed
func (e *Embed) AddField(name, value string, inline bool) *Embed {
	newField := &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	}

	e.Fields = append(e.Fields, newField)

	return e
}

// SetThumbnail sets the thumbnail for the embed
func (e *Embed) SetThumbnail(url string) *Embed {
	thumbnail := &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
	e.Thumbnail = thumbnail

	return e
}
