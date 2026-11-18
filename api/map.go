package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type MapApi struct {
	lastResponse *MapResponse
}

func getMapResource(url string) (*MapResponse, error) {
	if url == "" {
		return nil, fmt.Errorf("Failed to get resourse. Url is empty")
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to get resourse url: %s", url)
	}
	defer res.Body.Close()
	var mapRes MapResponse

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to get resourse url: %s", url)
	}

	err = json.Unmarshal(data, &mapRes)
	if err != nil {
		return nil, fmt.Errorf("Failed to get resourse url: %s", url)
	}

	return &mapRes, nil
}

func (m *MapApi) getDefault() []MapLocation {
	res, err := getMapResource(resourceLink)
	if err != nil {
		fmt.Println(err.Error())
		return m.lastResponse.Results
	}
	m.lastResponse = res
	return m.lastResponse.Results
}

func (m *MapApi) GetNext() []MapLocation {
	if m.lastResponse != nil {
		if len(m.lastResponse.NextUrl) > 0 {
			res, err := getMapResource(m.lastResponse.NextUrl)
			if err != nil {
				fmt.Println(err.Error())
				return m.lastResponse.Results
			}
			m.lastResponse = res
		}
		return m.lastResponse.Results
	}
	return m.getDefault()
}

func (m *MapApi) GetPrevious() []MapLocation {
	if m.lastResponse != nil {
		if len(m.lastResponse.PreviousUrl) > 0 {
			res, err := getMapResource(m.lastResponse.PreviousUrl)
			if err != nil {
				fmt.Println(err.Error())
				return m.lastResponse.Results
			}
			m.lastResponse = res
		}
		return m.lastResponse.Results
	}
	return m.getDefault()
}
