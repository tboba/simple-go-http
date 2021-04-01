package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// FileExists checks whether the given path refers to the file that exists.
// Returns boolean, where true value means that file exists in the given path.
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

// FileIsDirectory checks whether the given path refers to the file that is a directory.
// Returns boolean, where true value means that file is a directory.
func FileIsDirectory(filename string) bool {
	properties, _ := os.Stat(filename)

	if properties.IsDir() {
		return true
	}

	return false
}

// IsRunCommandCorrect checks whether the function 'run' can be correctly executed.
// Returns boolean, where true value means that command can be performed correctly.
func IsRunCommandCorrect(args []string, filename string) bool {
	// Check whether the command has all required parameters
	if len(args) < 4 {
		fmt.Println("> Cannot run 'run' command!")
		fmt.Println("Command doesn't have all required arguments!")
		fmt.Println("Run CLI with help parameter to show all available command arguments!")
		return false
	}

	// Check whether the command has a --file flag.
	if !strings.Contains(args[2], "--file") {
		fmt.Println("> Cannot run 'run' command!")
		fmt.Println("Command doesn't contain --file flag.")
		return false
	}

	// Check whether the file exists.
	if !FileExists(filename) {
		fmt.Println("> Cannot run 'run' command!")
		fmt.Println("File " + filename + " doesn't exist!")
		return false
	}

	// Check whether the file is a directory.
	if FileIsDirectory(filename) {
		fmt.Println("> Cannot run 'run' command!")
		fmt.Println("File " + filename + " is a directory!")
		return false
	}

	return true
}

// ServeHtmlServer allows to serve the HTML server, using the provided HTML file.
// It uses the configuration to launch the server with the provided port.
func ServeHtmlServer(configuration Config, filename string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, filename)
	})

	fmt.Println("Serving the HTTP server on http://localhost:" + configuration.Port + "/ ...")
	err := http.ListenAndServe(":"+configuration.Port, nil)
	if err != nil {
		fmt.Println("An error occurred during the HTTP server launch!")
		fmt.Println(err)
	}
}
