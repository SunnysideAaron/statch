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

	// Process multiple outputs
	outputs, ok := config["outputs"].([]any)
	if !ok {
		log.Fatal("config file missing 'outputs' array")
	}

	for _, out := range outputs {
		output, ok := out.(map[string]any)
		if !ok {
			log.Fatal("output is not a valid object")
		}

		templateFile, ok := output["templateFile"].(string)
		if !ok {
			log.Fatal("output missing templateFile")
		}

		outputFile, ok := output["outputFile"].(string)
		if !ok {
			log.Fatal("output missing outputFile")
		}

		// Create a new context for each output
		outCtx := plush.NewContext()
		outCtx.Set("config", output)

		t := loadTemplate(templateFile)
		s, err := plush.Render(t, outCtx)
		if err != nil {
			log.Fatalf("error rendering template %s: %v", templateFile, err)
		}

		err = os.WriteFile(outputFile, []byte(s), 0644)
		if err != nil {
			log.Fatalf("error writing to file %s: %v", outputFile, err)
		}
	}
}
