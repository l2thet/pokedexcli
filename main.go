package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	name     string
}

type Pokemon struct {
	Name           string
	BaseExperience int
}

const locationAreaBaseURL = "https://pokeapi.co/api/v2/location-area/"
const pokemonBaseURL = "https://pokeapi.co/api/v2/pokemon/"

var commands map[string]cliCommand
var pokedex map[string]Pokemon

func main() {
	commands = make(map[string]cliCommand)
	config := &config{
		next:     "",
		previous: "",
	}
	pokedex = make(map[string]Pokemon)

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

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "See all Pokemon in a location",
		callback:    commandExplore,
	}

	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Catch a Pokemon",
		callback:    commandCatch,
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

			if len(inputArray) > 1 && (inputArray[0] == "explore" || inputArray[0] == "catch") {
				config.name = inputArray[1]
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
		cfg.next = locationAreaBaseURL
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

func commandExplore(cfg *config) error {
	if len(cfg.name) == 0 {
		fmt.Println("Please enter a location name")
		return nil
	}

	fmt.Printf("Exploring %s...\n", cfg.name)

	locationDetails := pokeapi.GetLocationDetails(locationAreaBaseURL + cfg.name)

	if locationDetails.PokemonEncounters == nil {
		fmt.Println("No Pokemon found in this location")
		return nil
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range locationDetails.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config) error {
	if len(cfg.name) == 0 {
		fmt.Println("Please enter a Pokemon name")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.name)

	pokemonDetails := pokeapi.GetPokemonDetails(pokemonBaseURL + cfg.name)

	if pokemonDetails.Name == "" {
		fmt.Println("Pokemon not found")
		return nil
	}

	difficutlyChance := int(float64(pokemonDetails.BaseExperience) * 0.5)

	chance := rand.Intn(pokemonDetails.BaseExperience)

	if chance > difficutlyChance {
		fmt.Printf("%s was caught!\n", pokemonDetails.Name)
		pokedex[pokemonDetails.Name] = Pokemon{
			Name:           pokemonDetails.Name,
			BaseExperience: pokemonDetails.BaseExperience,
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)
	}

	return nil
}
