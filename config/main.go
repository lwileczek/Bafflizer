package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/lwileczek/Bafflizer/models"
)

var (
	file   *string
	output *string
	format *string
	letter *string
	lang   *string
)

func init() {
	file = flag.String("file", "", "Path to the input file to change")
	output = flag.String("output", "", "Path to save the output")
	format = flag.String("format", "Hash", "Convert variables to Hash or Letter")
	lang = flag.String("lang", "python", "Programming language of input file")
	letter = flag.String("letter", "a", "Letter to use as variable names")
}

//SetupEnv Check to see if global values were supplied via environment variables, cli, or config file
func SetupEnv() models.Config {
	flag.Parse()
	DefaultConfig := models.Config{
		Input:  *file,
		Output: *output,
		Format: *format,
	}
	if err := DefaultConfig.SetLang(*lang); err != nil {
		fmt.Println("Error creating default config", err)
	}
	if err := DefaultConfig.SetLetter(*letter); err != nil {
		fmt.Println("Error creating default config", err)
	}

	return DefaultConfig
}

//getEnv get key environment variable if exist otherwise return fallback
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
