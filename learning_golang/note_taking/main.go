package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type Outputable interface {
	Saver // Embed Saver interface
	Display() string
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println("Error saving the todo: ", err)
		return
	}

	outputData(userNote)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // single quotes are used for single characters

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n") // remove the newline character
	text = strings.TrimSuffix(text, "\r") // remove the carriage return character

	return text
}

func saveToFile(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving the data: ", err)
		return err
	}

	fmt.Println("Data saved successfully!")
	return nil
}

func outputData(data Outputable) error {
	fmt.Println(data.Display())
	return saveToFile(data)
}
