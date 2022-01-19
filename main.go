package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	defer os.Exit(0)

	// parse command line for both settings and csv
	csvPath := flag.String("csv", "",
		"The path of the CSV file in which are defined the custom Fujifilm simulation recipes")
	settingsPath := flag.String("s", "",
		"The path of your settings in YAML in which your camera name is set")
	flag.Parse()

	if len(*csvPath) == 0 || len(*settingsPath) == 0 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	fmt.Printf("Settings:     %s\n", *settingsPath)
	fmt.Printf("CSV to parse: %s\n", *csvPath)

	// load the user settings
	settings := loadUserSettings(settingsPath)

	// load the film simulation recipes from the CSV
	recipes := loadCSV(csvPath)

	generateXMLSimulations(&settings, recipes)
}

