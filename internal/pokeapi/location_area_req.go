package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (ResLocationArea, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache in
		resLocationArea := ResLocationArea{}
		err := json.Unmarshal(dat, &resLocationArea)
		if err != nil {
			return ResLocationArea{}, err
		}

		return resLocationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ResLocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResLocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return ResLocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return ResLocationArea{}, err
	}

	resLocationArea := ResLocationArea{}
	err = json.Unmarshal(dat, &resLocationArea)
	if err != nil {
		return ResLocationArea{}, err
	}

    c.cache.Add(fullURL, dat)

	return resLocationArea, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache in
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

    c.cache.Add(fullURL, dat)

	return locationArea, nil
}
