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

    dat, err := io.ReadAll(res.Body)
    if err != nil {
        return ResLocationArea{}, err
    }

    resLocationArea := ResLocationArea{}
    err = json.Unmarshal(dat, &resLocationArea)
    if err != nil {
        return ResLocationArea{}, err
    }

    return resLocationArea, nil
}

