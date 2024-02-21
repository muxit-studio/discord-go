package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

const VERSION = "v0.1.0"

func App() *cli.App {
	return &cli.App{
		Name:                 "discord-go",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name: "version",
				Action: func(c *cli.Context) error {
					fmt.Println(VERSION)
					return nil
				},
			},
			{
				Name:   "send",
				Usage:  "Send a message or file to a Discord webhook",
				Action: sendMessage,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "webhook-url", Required: true, Usage: "Discord webhook URL"},
					&cli.StringFlag{Name: "text", Usage: "Body text of the message to send"},
					&cli.BoolFlag{Name: "tts", Usage: "Send message with text-to-speech enabled"},
					&cli.StringFlag{Name: "username", Usage: "Set username to display"},
					&cli.StringFlag{Name: "avatar", Usage: "Set avatar URL"},
					&cli.StringFlag{Name: "file", Usage: "Path to the file to upload"},
					&cli.StringFlag{Name: "title", Usage: "Embed title"},
					&cli.StringFlag{Name: "description", Usage: "Embed description"},
					&cli.StringFlag{Name: "url", Usage: "Embed URL"},
					&cli.StringFlag{Name: "color", Usage: "Embed color as a decimal"},
					&cli.StringFlag{Name: "thumbnail", Usage: "Thumbnail URL for the embed"},
					// Author Flags
					&cli.StringFlag{Name: "author", Usage: "Author name"},
					&cli.StringFlag{Name: "author-icon", Usage: "Author icon URL"},
					&cli.StringFlag{Name: "author-url", Usage: "URL to be associated with the author"},
					// Footer Flags
					&cli.StringFlag{Name: "footer", Usage: "Footer text"},
					&cli.StringFlag{Name: "footer-icon", Usage: "Footer icon URL"},
					&cli.BoolFlag{Name: "timestamp", Usage: "Include timestamp in the footer"},
					// Image and Thumbnail Flags
					&cli.StringFlag{Name: "image", Usage: "Image URL for the embed"},
					// Fields Flags
					&cli.StringSliceFlag{Name: "field", Usage: "Add field to embed, format: name;value;inline"},
				},
			},
		},
	}
}

func sendMessage(c *cli.Context) error {
	fmt.Println("discord-go: sending message...")

	webhookURL := c.String("webhook-url")
	message := DiscordMessage{
		Content:   c.String("text"),
		Username:  c.String("username"),
		AvatarURL: c.String("avatar"),
		TTS:       c.Bool("tts"),
	}

	// Construct Embed if any embed-related flags are set
	var embed Embed
	fields := parseFields(c.StringSlice("field"))
	if len(fields) > 0 {
		embed.Fields = fields
	}
	if c.IsSet("title") {
		embed.Title = c.String("title")
	}
	if c.IsSet("description") {
		embed.Description = c.String("description")
	}
	if c.IsSet("url") {
		embed.URL = c.String("url")
	}
	if c.IsSet("color") {
		embed.Color = c.Int("color")
	}
	if c.IsSet("thumbnail") {
		embed.Thumbnail = map[string]string{"url": c.String("thumbnail")}
	}
	if c.IsSet("image") {
		embed.Image = map[string]string{"url": c.String("image")}
	}
	if c.IsSet("author") || c.IsSet("author-icon") || c.IsSet("author-url") {
		embed.Author = &EmbedAuthor{
			Name:    c.String("author"),
			IconURL: c.String("author-icon"),
			URL:     c.String("author-url"),
		}
	}
	if c.IsSet("footer") || c.IsSet("footer-icon") {
		embed.Footer = &EmbedFooter{
			Text:    c.String("footer"),
			IconURL: c.String("footer-icon"),
		}
	}
	if embed.Title != "" || embed.Description != "" {
		message.Embeds = append(message.Embeds, embed)
	}

	// Enforce embed limits
	embed.EnforceLimits()

	payloadBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	filePath := c.String("file")

	formData := FormData()

	if err := formData.Set("payload_json", string(payloadBytes)); err != nil {
		return fmt.Errorf("failed to add payload_json to form data: %v", err)
	}

	if filePath != "" {
		if err := formData.SetFile("file", filePath); err != nil {
			return fmt.Errorf("failed to add file to form data: %v", err)
		}
	}

	req, err := http.NewRequest("POST", webhookURL, formData.Reader())
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", formData.ContentType())

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("discord-go: received non-success status code: %d, body: %s", resp.StatusCode, string(body))
	}

	fmt.Println("discord-go: Message sent successfully")

	return nil
}

func parseFields(fieldStrings []string) []EmbedField {
	var fields []EmbedField
	for _, f := range fieldStrings {
		parts := strings.Split(f, ";")
		if len(parts) >= 3 {
			inline, err := strconv.ParseBool(parts[2])
			if err != nil {
				inline = false
			}
			fields = append(fields, EmbedField{
				Name:   parts[0],
				Value:  parts[1],
				Inline: inline,
			})
		}
	}
	return fields
}
