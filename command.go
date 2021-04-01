package main

import "fmt"

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
	// TODO: Serve an HTTP server.
}

func initializeCommandMap() {
	commandMap["help"] = "Displays all commands."
	commandMap["version"] = "Shows the current version of the application."
	commandMap["run --file <file>"] = "Launches basic HTTP server with the provided HTML file."
}
