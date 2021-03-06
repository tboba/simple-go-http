package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const version = "1.0.0"

// Config specifies the available keys in the configuration file.
type Config struct {
	Port string
}

// LoadConfiguration loads the basic configuration from the "config.yml" resource to the Config struct.
// It uses go-yaml library to parse the YAML file.
func LoadConfiguration() Config {
	config := Config{}
	fileConfig, err := ioutil.ReadFile("config.yml")

	if err != nil {
		fmt.Println("Cannot open the configuration file!")
		panic(err.Error())
	}

	// Parse the configuration file.
	err = yaml.Unmarshal(fileConfig, &config)

	if err != nil {
		fmt.Println("Cannot parse the configuration file!")
		panic(err.Error())
	}

	return config
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough parameters have been provided!")
		fmt.Println("Type \"help\" for the command list.")
		return
	}

	switch os.Args[1] {
	case "help":
		HandleHelpCommand()
	case "version":
		HandleVersionCommand()
	case "run":
		HandleRunCommand()
	default:
		HandleUnknownCommand()
	}
}
