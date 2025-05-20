package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/plush/v5"
)

func main() {
	ctx := plush.NewContext()

	// Read statch.json and add it to context
	statchData, err := readStatchJSON("statch.json")
	if err != nil {
		log.Fatal(err)
	}
	ctx.Set("statch", statchData)

	s, err := plush.Render(template(), ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("example.go", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func template() string {
	b, err := os.ReadFile("example.plush")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func readStatchJSON(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading statch.json: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("parsing statch.json: %w", err)
	}

	return result, nil
}
