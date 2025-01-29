package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	// struct tags are used to specify how the fields should be serialized
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() string {
	return note.Title + "\n" + note.Content + "\n" + note.CreatedAt.String()
}

func (note Note) Save() error {
	// Save the note to a file
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// Save the note to a JSON file
	jsonData, err := json.Marshal(note) // Convert the note content to a JSON string

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644) // 0644 is the file permission
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("title and content are required")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
