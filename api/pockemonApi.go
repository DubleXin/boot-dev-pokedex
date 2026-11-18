package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const pokemonResourceLink = "https://pokeapi.co/api/v2/pokemon/"

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type NamedResource struct {
	Name string `json:"name"`
}

type PokemonStat struct {
	BaseStat int           `json:"base_stat"`
	Stat     NamedResource `json:"stat"`
}

type PokemonType struct {
	Type NamedResource `json:"type"`
}

func (m *PokeApi) GetPokemon(name string) (Pokemon, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Pokemon{}, fmt.Errorf("cannot fetch pokemon: name is empty")
	}

	resourceURL := pokemonResourceLink + url.PathEscape(name)

	var data []byte
	var fromCache bool

	if m.Cache != nil {
		data, fromCache = m.Cache.Get(resourceURL)
	}

	if !fromCache {
		res, err := http.Get(resourceURL)
		if err != nil {
			return Pokemon{}, fmt.Errorf(
				"failed to request pokemon %q: %w",
				name, err,
			)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return Pokemon{}, fmt.Errorf(
				"failed to fetch pokemon %q: unexpected HTTP status %s",
				name, res.Status,
			)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf(
				"failed to read response body for pokemon %q: %w",
				name, err,
			)
		}
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, fmt.Errorf(
			"failed to decode pokemon %q: %w",
			name, err,
		)
	}

	if !fromCache && m.Cache != nil {
		m.Cache.Add(resourceURL, data)
	}

	return pokemon, nil
}
