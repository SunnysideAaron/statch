package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/knadh/koanf/parsers/hjson"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

func loadConfig(configFile string) error {
	if err := k.Load(file.Provider(configFile), hjson.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	return nil
}

func main() {
	// Define command line flags
	configFile := flag.String("config", "config.hjson", "config file of templates to render")
	flag.StringVar(configFile, "c", "config.hjson", "config file of templates to render (shorthand)")
	flag.Parse()

	// Load the configuration
	if err := loadConfig(*configFile); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	println(k.String("template.file"))

	// ctx := plush.NewContext()

	// // Read statch.json and add it to context
	// statchData, err := readStatchJSON("statch.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx.Set("statch", statchData)

	// s, err := plush.Render(template(), ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = os.WriteFile("example.go", []byte(s), 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
