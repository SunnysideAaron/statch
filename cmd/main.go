package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ajitpratap0/GoSQLX/pkg/sql/ast"
	"github.com/ajitpratap0/GoSQLX/pkg/sql/parser"
	"github.com/ajitpratap0/GoSQLX/pkg/sql/token"
	"github.com/ajitpratap0/GoSQLX/pkg/sql/tokenizer"

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

func loadAST() {
	// 1. Get a tokenizer from the pool
	tkz := tokenizer.GetTokenizer()
	defer tokenizer.PutTokenizer(tkz) // Return to pool when done

	// 2. Tokenize the SQL query
	sql := []byte("SELECT id, name FROM users WHERE age > 18")
	tokens, err := tkz.Tokenize(sql)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tokenized SQL into %d tokens\n", len(tokens))

	// 3. Convert TokenWithSpan to token.Token for the parser
	tokenSlice := make([]token.Token, len(tokens))
	for i, t := range tokens {
		// Debug the token types
		fmt.Printf("Token %d: Type=%v, Value=%s\n", i, t.Token.Type, t.Token.Value)

		// Map token types correctly
		var tokenType token.Type
		switch t.Token.Type {
		case 43: // SELECT
			tokenType = token.SELECT
		case 1: // IDENTIFIER
			tokenType = token.IDENT
		case 124: // COMMA
			tokenType = token.COMMA
		case 59: // FROM
			tokenType = token.FROM
		case 51: // WHERE
			tokenType = token.WHERE
		case 20: // GREATER THAN
			tokenType = token.GT
		case 2: // NUMBER
			tokenType = token.INT
		case 0: // EOF
			tokenType = token.EOF
		default:
			tokenType = token.ILLEGAL
		}

		tokenSlice[i] = token.Token{
			Type:    tokenType,
			Literal: t.Token.Value,
		}
	}

	// 4. Create a parser
	p := parser.NewParser()
	defer p.Release() // Clean up resources

	// 5. Parse tokens into an AST
	result, err := p.Parse(tokenSlice)
	if err != nil {
		fmt.Printf("Error parsing tokens: %v\n", err)
		return
	}

	// 6. Work with the AST
	fmt.Printf("Parsed %d statements\n", len(result.Statements))

	// 7. Return AST to the pool when done
	ast.ReleaseAST(result)
}

func loadSchema(file string) (map[string]any, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("reading schema file: %w", err)
	}

	// Parse the schema file by splitting on semicolons
	content := string(data)
	rawStatements := strings.Split(content, ";")

	// Clean up statements and remove empty ones
	var statements []string
	for _, stmt := range rawStatements {
		trimmed := strings.TrimSpace(stmt)
		if trimmed != "" {
			statements = append(statements, trimmed)
		}
	}

	loadAST()

	return map[string]any{"statements": statements}, nil
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

		outputFile, ok := output["outputFile"].(string)
		if !ok {
			log.Fatal("output missing outputFile")
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

		err = os.WriteFile(outputFile, []byte(s), 0644)
		if err != nil {
			log.Fatalf("error writing to file %s: %v", outputFile, err)
		}
	}
}
