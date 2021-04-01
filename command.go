package main

import (
	"flag"
	"fmt"
	"os"
)

var commandMap = make(map[string]string)

// HandleUnknownCommand sends the message to the user if the command cannot be recognized by the CLI.
func HandleUnknownCommand() {
	fmt.Println("Cannot recognize provided command! Type \"help\" to get available commands.")
}

// HandleHelpCommand sends the message to the user with the values from the commandMap.
// It displays all available messages.
func HandleHelpCommand() {
	// If the command map is empty, register all commands to the command map.
	if len(commandMap) == 0 {
		initializeCommandMap()
	}

	fmt.Println("> Available commands:")
	for command, description := range commandMap {
		fmt.Println("  " + command + " - " + description)
	}
}

// HandleVersionCommand sends the message to the user with the current version of the CLI.
func HandleVersionCommand() {
	fmt.Println("Current version: " + version)
}

// HandleRunCommand launches the HTTP server on the local machine with the provided port from the configuration.
func HandleRunCommand() {
	configuration := LoadConfiguration()

	// Create a new flag set and register the 'file' flag.
	fileFlag := flag.NewFlagSet("run", flag.ExitOnError)
	fileName := fileFlag.String("file", "",
		"Correct path to the file with adequate extension.")

	// Parse all registered flags.
	_ = fileFlag.Parse(os.Args[2:])

	// Launch HTTP server if all parameters are correct.
	if IsRunCommandCorrect(os.Args, *fileName) {
		ServeHtmlServer(configuration, *fileName)
	}
}

// initializeCommandMap puts all available commands into the commandMap variable.
func initializeCommandMap() {
	commandMap["help"] = "Displays all commands."
	commandMap["version"] = "Shows the current version of the application."
	commandMap["run --file <file>"] = "Launches basic HTTP server with the provided HTML file."
}
