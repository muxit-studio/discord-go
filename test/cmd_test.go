package discord_test

import (
	"testing"
)

// TestTextNoUsername tests sending a text message without a username.
func TestTextNoUsername(t *testing.T) {
	err := run("--text", "text, no username, --gnu-style=<>")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

// TestTextUsername tests sending a text message with a username.
func TestTextUsername(t *testing.T) {
	err := run("--username", "username test", "--text", "text, username, --gnu-style varied")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

// TestTextUsernameAvatar tests sending a text message with a username and an avatar URL.
func TestTextUsernameAvatar(t *testing.T) {
	err := run("--username", "avatar test", "--text", "avatar test", "--avatar", "https://i.imgur.com/o96JZ1Y.png")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

// TestTextInvalidWebhookURL tests sending a message with an invalid webhook URL.
func TestTextInvalidWebhookURL(t *testing.T) {
	err := run("--text", "it should be error", "--webhook-url", "https://discordapp.com/api/webhooks/invalid/webhook")
	if err == nil {
		t.Errorf("Expected an error due to invalid webhook URL, but got none")
	}
}
