package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const exploreResourceLink = "https://pokeapi.co/api/v2/location-area/"

type ExploreResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonStr `json:"pokemon"`
}

type PokemonStr struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (p *PokeApi) GetExplorePokemons(location string) ([]PokemonStr, error) {
	location = strings.TrimSpace(location)
	if location == "" {
		return []PokemonStr{}, nil
	}

	resourceURL := exploreResourceLink + url.PathEscape(location)

	var data []byte
	var fromCache bool

	if p.Cache != nil {
		data, fromCache = p.Cache.Get(resourceURL)
	}

	if !fromCache {
		res, err := http.Get(resourceURL)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to request pokemon encounters for location %q: %w",
				location, err,
			)
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusNotFound {
			return []PokemonStr{}, nil
		}

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf(
				"failed to fetch pokemon encounters for location %q: unexpected HTTP status %s",
				location, res.Status,
			)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to read pokemon encounters for location %q: %w",
				location, err,
			)
		}
	}

	var response ExploreResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf(
			"failed to decode pokemon encounters for location %q: %w",
			location, err,
		)
	}

	if !fromCache && p.Cache != nil {
		p.Cache.Add(resourceURL, data)
	}

	pokemons := make([]PokemonStr, len(response.PokemonEncounters))
	for i, encounter := range response.PokemonEncounters {
		pokemons[i] = encounter.Pokemon
	}

	return pokemons, nil
}
