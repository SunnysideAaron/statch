package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/plush/v5"
	"github.com/hjson/hjson-go/v4"
)

func loadHJSON(configFile string) (map[string]any, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("reading hjson file: %w", err)
	}

	var config map[string]any
	if err := hjson.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parsing hjson: %w", err)
	}

	return config, nil
}

func loadTemplate(file string) string {
	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func main() {
	// Load the configuration
	configFile := flag.String("config", "config.hjson", "config file of templates to render")
	flag.StringVar(configFile, "c", "config.hjson", "config file of templates to render (shorthand)")
	flag.Parse()

	config, err := loadHJSON(*configFile)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	ctx := plush.NewContext()
	ctx.Set("config", config)

	output := config["output"].(map[string]any)
	t := loadTemplate(output["templateFile"].(string))

	s, err := plush.Render(t, ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(output["file"].(string), []byte(s), 0644)
	if err != nil {

		log.Fatal(err)
	}

}
