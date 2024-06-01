package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullurl := baseUrl + endpoint
	if pageUrl != nil {
		fullurl = *pageUrl
	}

	data, ok := c.cache.Get(fullurl)
	if !ok {
		//make http request
		req, err := http.NewRequest("GET", fullurl, nil)
		if err != nil {
			return LocationAreaResp{}, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResp{}, err
		}

		if resp.StatusCode > 399 {
			return LocationAreaResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResp{}, fmt.Errorf("error retreiving data : %s", err)
		}
		// take the data from the response, and a pointer to instance of the Location struct
		// and set the data to the struct
		locationAreasResp := LocationAreaResp{}
		err = json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		c.cache.Add(fullurl, data)

		return locationAreasResp, nil

	}

	// take the data from the response, and a pointer to instance of the Location struct
	// and set the data to the struct
	locationAreasResp := LocationAreaResp{}
	err := json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullurl := baseUrl + endpoint

	data, ok := c.cache.Get(fullurl)
	if !ok {
		//make http request
		req, err := http.NewRequest("GET", fullurl, nil)
		if err != nil {
			return LocationArea{}, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}

		if resp.StatusCode > 399 {
			return LocationArea{}, fmt.Errorf("bad status code %v", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error retreiving data : %s", err)
		}
		// take the data from the response, and a pointer to instance of the Location struct
		// and set the data to the struct
		locationArea := LocationArea{}
		err = json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		c.cache.Add(fullurl, data)

		return locationArea, nil

	}

	// take the data from the response, and a pointer to instance of the Location struct
	// and set the data to the struct
	locationArea := LocationArea{}
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}

func (c *Client) CatchPokemon(pokemon string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemon
	fullurl := baseUrl + endpoint

	data, ok := c.cache.Get(fullurl)
	if !ok {
		//make http request
		req, err := http.NewRequest("GET", fullurl, nil)
		if err != nil {
			return Pokemon{}, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		if resp.StatusCode > 399 {
			return Pokemon{}, fmt.Errorf("bad status code %v", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error retreiving data : %s", err)
		}
		// take the data from the response, and a pointer to instance of the Location struct
		// and set the data to the struct
		pokemonData := Pokemon{}
		err = json.Unmarshal(data, &pokemonData)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(fullurl, data)

		return pokemonData, nil

	}

	// take the data from the response, and a pointer to instance of the Location struct
	// and set the data to the struct
	pokemonData := Pokemon{}
	err := json.Unmarshal(data, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonData, nil
}
