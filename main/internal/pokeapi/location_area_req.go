package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullurl := baseUrl + endpoint

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

	return locationAreasResp, nil
}
