package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

// ListLocations -
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	// check if the data is in the cache.
	val, ok := c.cache.Get(url)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}

	// cache miss 
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	// add the data to the cache.
	c.cache.Add(url, dat)

	return locationResp, nil
}



