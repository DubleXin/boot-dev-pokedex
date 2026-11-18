package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DubleXin/pokedexcli/pokecache"
)

const resourceLink = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

type MapResponse struct {
	NextUrl     string        `json:"next"`
	PreviousUrl string        `json:"previous"`
	Results     []MapLocation `json:"results"`
}

type MapLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeApi struct {
	lastMapResponse *MapResponse
	Cache           *pokecache.Cache
}

func (m *PokeApi) getMapResource(url string) (*MapResponse, error) {
	if url == "" {
		return nil, fmt.Errorf("cannot fetch map resource: URL is empty")
	}

	if m.Cache != nil {
		if data, ok := m.Cache.Get(url); ok {
			var mapRes MapResponse
			if err := json.Unmarshal(data, &mapRes); err != nil {
				return nil, fmt.Errorf(
					"failed to decode cached map response for %q: %w",
					url, err,
				)
			}
			return &mapRes, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to request map resource %q: %w",
			url, err,
		)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"failed to fetch map resource %q: unexpected HTTP status %s",
			url, res.Status,
		)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read map response body for %q: %w",
			url, err,
		)
	}

	var mapRes MapResponse
	if err := json.Unmarshal(data, &mapRes); err != nil {
		return nil, fmt.Errorf(
			"failed to decode map response for %q: %w",
			url, err,
		)
	}

	if m.Cache != nil {
		m.Cache.Add(url, data)
	}

	return &mapRes, nil
}

func (m *PokeApi) getDefault() []MapLocation {
	res, err := m.getMapResource(resourceLink)
	if err != nil {
		fmt.Println(err)
		if m.lastMapResponse != nil {
			return m.lastMapResponse.Results
		}
		return nil
	}

	m.lastMapResponse = res
	return res.Results
}

func (m *PokeApi) GetNext() []MapLocation {
	if m.lastMapResponse == nil {
		return m.getDefault()
	}

	if m.lastMapResponse.NextUrl == "" {
		return m.lastMapResponse.Results
	}

	res, err := m.getMapResource(m.lastMapResponse.NextUrl)
	if err != nil {
		fmt.Println(err)
		return m.lastMapResponse.Results
	}

	m.lastMapResponse = res
	return res.Results
}

func (m *PokeApi) GetPrevious() []MapLocation {
	if m.lastMapResponse == nil {
		return m.getDefault()
	}

	if m.lastMapResponse.PreviousUrl == "" {
		return m.lastMapResponse.Results
	}

	res, err := m.getMapResource(m.lastMapResponse.PreviousUrl)
	if err != nil {
		fmt.Println(err)
		return m.lastMapResponse.Results
	}

	m.lastMapResponse = res
	return res.Results
}
