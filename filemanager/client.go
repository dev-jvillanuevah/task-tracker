package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	filename string
}

func NewClient(filename string) *Client {
	return &Client{filename: filename}
}

func (f *Client) WriteFile(tasks any) error {
	file, err := os.Create(f.filename)
	if err != nil {
		return fmt.Errorf("error creating tasks.json: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(tasks); err != nil {
		return fmt.Errorf("error writing tasks.json: %w", err)
	}
	return nil
}

func (f *Client) ReadFile(tasks any) error {
	file, err := os.Open(f.filename)
	if err != nil {
		return fmt.Errorf("error opening tasks.json: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tasks); err != nil {
		return fmt.Errorf("error reading tasks.json: %w", err)
	}
	return nil
}
