package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}

const baseURL = "https://pokeapi.co/api/v2/location-area/"

var commands map[string]cliCommand

func main() {
	commands = make(map[string]cliCommand)
	config := &config{
		next:     "",
		previous: "",
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Display 20 map locations",
		callback:    commandMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Display the previous 20 map locations if they exist",
		callback:    commandMapb,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		inputArray := cleanInput(text)
		//fmt.Printf("Your command was: %v\n", inputArray[0])

		if len(inputArray) != 0 {
			command, ok := commands[inputArray[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := command.callback(config)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}
}

func cleanInput(text string) []string {
	parts := strings.Fields(strings.ToLower(strings.Trim(text, " ")))
	return parts
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(cfg *config) error {
	if len(cfg.next) == 0 {
		cfg.next = baseURL
	}

	locations := pokeapi.GetLocations(cfg.next)

	cfg.next = locations.Next

	if locations.Previous != nil {
		cfg.previous = *locations.Previous
	} else {
		cfg.previous = ""
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if len(cfg.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	locations := pokeapi.GetLocations(cfg.previous)

	cfg.next = locations.Next

	if locations.Previous != nil {
		cfg.previous = *locations.Previous
	} else {
		cfg.previous = ""
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
