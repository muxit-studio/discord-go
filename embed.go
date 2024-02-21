package discord

import "fmt"

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type EmbedAuthor struct {
	Name    string `json:"name,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
	URL     string `json:"url,omitempty"`
}

type EmbedFooter struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Embed struct {
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	URL         string            `json:"url,omitempty"`
	Color       int               `json:"color,omitempty"`
	Fields      []EmbedField      `json:"fields,omitempty"`
	Author      *EmbedAuthor      `json:"author,omitempty"`
	Footer      *EmbedFooter      `json:"footer,omitempty"`
	Image       map[string]string `json:"image,omitempty"`
	Thumbnail   map[string]string `json:"thumbnail,omitempty"`
	Timestamp   string            `json:"timestamp,omitempty"`
}

func (e *Embed) EnforceLimits() {
	if len(e.Title) > 256 {
		e.Title = e.Title[:256]
		fmt.Println("warning: embed title limited to 256 characters")
	}
	if len(e.Description) > 2048 {
		e.Description = e.Description[:2048]
		fmt.Println("warning: embed description limited to 2048 characters")
	}
	if e.Footer != nil && len(e.Footer.Text) > 2048 {
		e.Footer.Text = e.Footer.Text[:2048]
		fmt.Println("warning: embed footer text limited to 2048 characters")
	}
	if e.Author != nil && len(e.Author.Name) > 256 {
		e.Author.Name = e.Author.Name[:256]
		fmt.Println("warning: embed author name limited to 256 characters")
	}
}
