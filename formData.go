package discord

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

// formData provides a simplified interface for creating multipart form data.
type formData struct {
	buffer bytes.Buffer
	writer *multipart.Writer
}

// NewFormData creates a new FormData instance.
func FormData() *formData {
	fd := &formData{}
	fd.writer = multipart.NewWriter(&fd.buffer)
	return fd
}

// SetField adds a text field to the form data.
func (fd *formData) Set(fieldName, value string) error {
	return fd.writer.WriteField(fieldName, value)
}

// SetFile adds a file field to the form data.
func (fd *formData) SetFile(fieldName, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := fd.writer.CreateFormFile(fieldName, filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	return err
}

// ContentType returns the Content-Type of the FormData, including the boundary.
func (fd *formData) ContentType() string {
	return fd.writer.FormDataContentType()
}

// Bytes returns the raw bytes of the constructed form data.
func (fd *formData) Bytes() []byte {
	// Ensure multipart writer is closed before returning bytes
	fd.writer.Close()
	return fd.buffer.Bytes()
}

// Reader returns a bytes.Reader for the form data.
func (fd *formData) Reader() *bytes.Reader {
	return bytes.NewReader(fd.Bytes())
}
