package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Todo struct {
	// struct tags are used to specify how the fields should be serialized
	Text string `json:"text"`
}

func (todo Todo) Display() string {
	return todo.Text
}

func (todo Todo) Save() error {
	// Save the note to a file
	fileName := "todo.json"

	// Save the note to a JSON file
	jsonData, err := json.Marshal(todo) // Convert the note content to a JSON string

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644) // 0644 is the file permission
}

func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("content is required")
	}

	return &Todo{
		Text: content,
	}, nil
}
