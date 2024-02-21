package discord_test

import "testing"

// Test sending an empty title, which should fail
func TestEmbedEmptyTitle(t *testing.T) {
	err := run("--title", "")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestEmbedWithTitle(t *testing.T) {
	err := run("--title", "Hello, World!")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

// Test sending an embed with a title and description
func TestEmbedWithTitleAndDescription(t *testing.T) {
	err := run("--title", "Hello, World!", "--description", "This is a test")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

// TestEmbedWithColor tests sending an embed with a title, description, and color.
func TestEmbedWithColor(t *testing.T) {
	err := run("--title", "Title with Color", "--description", "This embed has a color.", "--color", "15258703")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

// TestEmbedWithThumbnail tests sending an embed with a title, description, and thumbnail URL.
func TestEmbedWithThumbnail(t *testing.T) {
	err := run("--title", "Title with Thumbnail", "--description", "This embed has a thumbnail.", "--thumbnail", "https://i.imgur.com/o96JZ1Y.png")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

// TestEmbedWithImage tests sending an embed with a title, description, and image URL.
func TestEmbedWithImage(t *testing.T) {
	err := run("--title", "Title with Image", "--description", "This embed has an image.", "--image", "https://i.imgur.com/o96JZ1Y.png")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

// TestEmbedWithAuthor tests sending an embed with a title, description, and author information.
func TestEmbedWithAuthor(t *testing.T) {
	err := run("--title", "Title with Author Icon and Url", "--description", "This embed has an author.", "--author", "Author Name", "--author-url", "https://example.com/author", "--author-icon", "https://i.imgur.com/o96JZ1Y.png")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

// TestEmbedWithFooter tests sending an embed with a title, description, and footer information.
func TestEmbedWithFooter(t *testing.T) {
	err := run("--title", "Title with Footer Icon", "--description", "This embed has a footer.", "--footer", "Footer Text", "--footer-icon", "https://i.imgur.com/o96JZ1Y.png")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

// TestEmbedWithFields tests sending an embed with a title, description, and fields.
func TestEmbedWithFields(t *testing.T) {
	err := run(
		"--title", "Title with Fields",
		"--description", "This embed has fields.",
		"--field", "Name1;Value1;false",
		"--field", "Name2;Value2;true",
	)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

// TestEmbedFullDetails tests sending an embed with full details including title, description, color, URL, author, image, thumbnail, footer, and fields.
func TestEmbedFullDetails(t *testing.T) {
	err := run(
		"--title", "Full Detail Test", "--description", "This embed has full details.", "--color", "15258703",
		"--url", "https://example.com",
		"--author", "Author Name", "--author-url", "https://i.imgur.com/o96JZ1Y.png", "--author-icon", "https://i.imgur.com/o96JZ1Y.png",
		"--image", "https://i.imgur.com/o96JZ1Y.png",
		"--thumbnail", "https://i.imgur.com/o96JZ1Y.png",
		"--footer", "Footer Text", "--footer-icon", "https://i.imgur.com/o96JZ1Y.png",
		"--field", "Name1;Value1;false", "--field", "Name2;Value2;true",
	)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
