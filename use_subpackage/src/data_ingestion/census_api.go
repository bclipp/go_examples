//This module is used for interacting with the census rest api

package data_ingestion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// used for adjusting the timeout
var httpClient = &http.Client{Timeout: 10 * time.Second}

// Inner is the second layer of the json document
type inner struct {
	BlockId   int    `json:"block_id"`
	StateCode string `json:"state_code"`
	StateFips int    `json:"state_fips"`
	BlockPop  int    `json:"block_pop_2015"`
}

//Outer is the first layer of the json document
type outer struct {
	Results []inner `json:"results"`
}

// census_api is used for interacting indirectly with he census api
// Params:
//       latitude: number to be added to y
//       longitude: number to be added to x
//return:
//       Jason return document
//       rest http response code
//       the error
func censusApi(latitude, longitude float64) (outer, int, error) {
	url := "https://geo.fcc.gov/api/census/area?lat=" + FloatToString(latitude) + "0&lon=" + FloatToString(longitude) + "&format=json"
	response, err := httpClient.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var census outer
	err = json.Unmarshal(body, &census)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("%+v\n", census)
	return census, response.StatusCode, err
}
