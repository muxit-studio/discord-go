package discord_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	cmd "github.com/muxit-studio/discord-go"
)

// helper function to run the CLI command
func run(args ...string) error {
	webhookURL := os.Getenv("DISCORD_WEBHOOK")
	if webhookURL == "" {
		log.Fatal("DISCORD_WEBHOOK environment variable is not set")
	}

	if webhookURL != "" {
		args = append([]string{"discord-go", "send", "--webhook-url", webhookURL}, args...)
	}

	err := cmd.App().Run(args)

	time.Sleep(1 * time.Second)

	return err
}

func TestMain(m *testing.M) {
	// Load environment variables from .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Run tests
	m.Run()
}
