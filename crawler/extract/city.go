package extract

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Cities struct is data structure of result
type Cities struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Slug    string `json:"slug"`
}

// GetCity is func to extract city
func GetCity(urlCity string) []string {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, urlCity, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "fave-testcoding")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var cities []Cities

	var cityList []string
	json.Unmarshal([]byte(body), &cities)
	for _, city := range cities {
		cityList = append(cityList, city.Slug)
	}
	return cityList
}
