package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const mapResourceLink = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

type MapResponse struct {
	NextUrl     string        `json:"next"`
	PreviousUrl string        `json:"previous"`
	Results     []MapLocation `json:"results"`
}

type MapLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (p *PokeApi) getMap(url string) (*MapResponse, error) {
	if url == "" {
		return nil, fmt.Errorf("cannot fetch map resource: URL is empty")
	}

	if p.Cache != nil {
		if data, ok := p.Cache.Get(url); ok {
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

	if p.Cache != nil {
		p.Cache.Add(url, data)
	}

	return &mapRes, nil
}

func (p *PokeApi) getDefaultLocation() []MapLocation {
	res, err := p.getMap(mapResourceLink)
	if err != nil {
		fmt.Println(err)
		if p.lastMapResponse != nil {
			return p.lastMapResponse.Results
		}
		return nil
	}

	p.lastMapResponse = res
	return res.Results
}

func (p *PokeApi) GetNextLocation() []MapLocation {
	if p.lastMapResponse == nil {
		return p.getDefaultLocation()
	}

	if p.lastMapResponse.NextUrl == "" {
		return p.lastMapResponse.Results
	}

	res, err := p.getMap(p.lastMapResponse.NextUrl)
	if err != nil {
		fmt.Println(err)
		return p.lastMapResponse.Results
	}

	p.lastMapResponse = res
	return res.Results
}

func (p *PokeApi) GetPreviousLocation() []MapLocation {
	if p.lastMapResponse == nil {
		return p.getDefaultLocation()
	}

	if p.lastMapResponse.PreviousUrl == "" {
		return p.lastMapResponse.Results
	}

	res, err := p.getMap(p.lastMapResponse.PreviousUrl)
	if err != nil {
		fmt.Println(err)
		return p.lastMapResponse.Results
	}

	p.lastMapResponse = res
	return res.Results
}
