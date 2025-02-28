package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	Filename string
}

func NewClient(filename string) *Client {
	return &Client{Filename: filename}
}

func (f *Client) WriteFile(tasks any) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return fmt.Errorf("error creating tasks.json: %w", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(tasks); err != nil {
		return fmt.Errorf("error writing tasks.json: %w", err)
	}
	return nil
}

func (f *Client) ReadFile(tasks any) error {
	file, err := os.Open("tasks.json")
	if err != nil {
		return fmt.Errorf("error opening tasks.json: %w", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tasks); err != nil {
		return fmt.Errorf("error reading tasks.json: %w", err)
	}
	return nil
}
