package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int       `json:"chance"`
				ConditionValues []*string `json:"condition_values"`
				MaxLevel        int       `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type PokemonDetails struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string  `json:"back_default"`
		BackFemale       *string `json:"back_female"`
		BackShiny        string  `json:"back_shiny"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontDefault     string  `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       string  `json:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string  `json:"front_default"`
				FrontFemale  *string `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string  `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       string  `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string  `json:"back_default"`
				BackFemale       *string `json:"back_female"`
				BackShiny        string  `json:"back_shiny"`
				BackShinyFemale  *string `json:"back_shiny_female"`
				FrontDefault     string  `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       string  `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string  `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        string  `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string  `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        string  `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string  `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        string  `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string  `json:"back_default"`
						BackFemale       *string `json:"back_female"`
						BackShiny        string  `json:"back_shiny"`
						BackShinyFemale  *string `json:"back_shiny_female"`
						FrontDefault     string  `json:"front_default"`
						FrontFemale      *string `json:"front_female"`
						FrontShiny       string  `json:"front_shiny"`
						FrontShinyFemale *string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string  `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        string  `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string  `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string  `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
}

func GetLocations(url string) Locations {
	cache := pokecache.NewCache(5 * time.Millisecond)

	if val, ok := cache.Get(url); ok {
		locations := Locations{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			log.Fatal(err)
		}

		return locations
	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		cache.Add(url, body)

		locations := Locations{}
		err = json.Unmarshal(body, &locations)
		if err != nil {
			log.Fatal(err)
		}

		return locations
	}
}

func GetLocationDetails(url string) LocationDetails {
	cache := pokecache.NewCache(5 * time.Millisecond)

	if val, ok := cache.Get(url); ok {
		locationDetails := LocationDetails{}
		err := json.Unmarshal(val, &locationDetails)
		if err != nil {
			log.Fatal(err)
		}

		return locationDetails
	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		cache.Add(url, body)

		locationDetails := LocationDetails{}
		err = json.Unmarshal(body, &locationDetails)
		if err != nil {
			log.Fatal(err)
		}

		return locationDetails
	}
}

func GetPokemonDetails(url string) PokemonDetails {
	cache := pokecache.NewCache(5 * time.Millisecond)

	if val, ok := cache.Get(url); ok {
		pokemonDetails := PokemonDetails{}
		err := json.Unmarshal(val, &pokemonDetails)
		if err != nil {
			log.Fatal(err)
		}

		return pokemonDetails
	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		cache.Add(url, body)

		pokemonDetails := PokemonDetails{}
		err = json.Unmarshal(body, &pokemonDetails)
		if err != nil {
			log.Fatal(err)
		}

		return pokemonDetails
	}
}
