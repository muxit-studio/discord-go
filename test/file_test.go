package discord_test

import "testing"

// Test sending an empty title, which should fail
func TestValidFileUpload(t *testing.T) {
	err := run("--text", "Test.txt file", "--file", "text.txt")
	if err != nil {
		t.Error("error uploading file:", err)
	}
}
