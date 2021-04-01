package main

import (
	"flag"
	"fmt"
	"os"
)

var commandMap = make(map[string]string)

func HandleUnknownCommand() {
	fmt.Println("Cannot recognize provided command! Type \"help\" to get available commands.")
}

func HandleHelpCommand() {
	if len(commandMap) == 0 {
		initializeCommandMap()
	}

	fmt.Println("> Available commands:")
	for command, description := range commandMap {
		fmt.Println("  " + command + " - " + description)
	}
}

func HandleVersionCommand() {
	fmt.Println("Current version: " + version)
}

func HandleRunCommand() {
	configuration := LoadConfiguration()

	fileFlag := flag.NewFlagSet("run", flag.ExitOnError)
	fileName := fileFlag.String("file", "",
		"Correct path to the file with adequate extension.")

	_ = fileFlag.Parse(os.Args[2:])

	if IsRunCommandCorrect(os.Args, *fileName) {
		ServeHtmlServer(configuration, *fileName)
	}
}

func initializeCommandMap() {
	commandMap["help"] = "Displays all commands."
	commandMap["version"] = "Shows the current version of the application."
	commandMap["run --file <file>"] = "Launches basic HTTP server with the provided HTML file."
}
