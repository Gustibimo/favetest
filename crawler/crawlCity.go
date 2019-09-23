package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Gustibimo/favetest/crawler/extract"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// const (
// 	host     = "localhost"
// 	port     = 3306
// 	user     = "root"
// 	password = "0341"
// 	dbname   = "fave_merchant"
// 	charset  = "utf8"
// 	table    = "city"
// )

func main() {
	db, err := sqlx.Connect("mysql", user+":"+password+"@/"+dbname+"?charset="+charset)
	if err != nil {
		panic(err)
	}

	fmt.Printf("connected")
	defer db.Close()

	urlCity := "https://myfave.com/api/mobile/cities"

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

	var cities []extract.Cities

	json.Unmarshal([]byte(body), &cities)
	for _, city := range cities {
		sqlStatement := `
		INSERT INTO city (country, currency, city_id, lat, lon, city_name, slug)
					VALUES (?, ?, ?, ?, ?, ?, ?)`

		db.MustExec(sqlStatement, city.Country, city.Currency, city.ID, city.Lat, city.Lng, city.Name, city.Slug)
		if err != nil {
			panic(err)
		}
	}

}
