package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/plush/v5"
	"github.com/hjson/hjson-go/v4"

	pg_query "github.com/pganalyze/pg_query_go/v6"
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

// func parseSchema(schema string) ([]pg_query.ParseResult, error) {
// 	return pg_query.Parse(schema)
// }

// func parseSchema() {
// 	tree, err := pg_query.ParseToJSON("SELECT 1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", tree)
// }

func loadSchema(file string) (map[string]any, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("reading schema file: %w", err)
	}

	// Parse the schema file by splitting on semicolons
	content := string(data)
	rawStatements := strings.Split(content, ";")

	// Clean up statements and remove empty ones
	var statementsData []map[string]any
	for _, stmt := range rawStatements {
		trimmed := strings.TrimSpace(stmt)
		if trimmed == "" {
			continue
		}

		// Parse each statement with pg_query
		parseResult, err := pg_query.Parse(trimmed)
		if err != nil {
			// Skip statements that can't be parsed
			continue
		}

		statementsData = append(statementsData, map[string]any{
			"sql":          trimmed,
			"parse_result": parseResult,
		})
	}

	return map[string]any{"statements": statementsData}, nil
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

		generatedFile, ok := output["generatedFile"].(string)
		if !ok {
			log.Fatal("output missing generatedFile")
		}

		ctx := plush.NewContext()
		ctx.Set("config", output)

		// Process sources if they exist
		if sources, ok := output["sources"].([]any); ok {
			for _, src := range sources {
				source, ok := src.(map[string]any)
				if !ok {
					log.Fatal("source is not a valid object")
				}

				sourceFile, ok := source["sourceFile"].(string)
				if !ok {
					log.Fatal("source missing sourceFile")
				}

				sourceType, ok := source["type"].(string)
				if !ok {
					log.Fatal("source missing type")
				}

				if sourceType == "schema" {
					schema, err := loadSchema(sourceFile)
					if err != nil {
						log.Fatalf("error loading schema %s: %v", sourceFile, err)
					}
					ctx.Set("schema", schema)
				}
				// Other source types can be handled here
			}
		}

		t := loadTemplate(templateFile)
		s, err := plush.Render(t, ctx)
		if err != nil {
			log.Fatalf("error rendering template %s: %v", templateFile, err)
		}

		err = os.WriteFile(generatedFile, []byte(s), 0644)
		if err != nil {
			log.Fatalf("error writing to file %s: %v", generatedFile, err)
		}
	}
}
