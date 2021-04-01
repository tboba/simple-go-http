package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func FileIsDirectory(filename string) bool {
	properties, _ := os.Stat(filename)

	if properties.IsDir() {
		return true
	}

	return false
}

func IsRunCommandCorrect(args []string, filename string) bool {
	if len(args) < 4 {
		fmt.Println("> Cannot run 'run' command!")
		fmt.Println("Command doesn't have all required arguments!")
		fmt.Println("Run CLI with help parameter to show all available command arguments!")
		return false
	}

	if !strings.Contains(args[2], "--file") {
		fmt.Println("> Cannot run 'run' command!")
		fmt.Println("Command doesn't contain --file flag.")
		return false
	}

	if !FileExists(filename) {
		fmt.Println("> Cannot run 'run' command!")
		fmt.Println("File " + filename + " doesn't exist!")
		return false
	}

	if FileIsDirectory(filename) {
		fmt.Println("> Cannot run 'run' command!")
		fmt.Println("File " + filename + " is a directory!")
		return false
	}

	return true
}

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
