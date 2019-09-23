package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	extract "github.com/Gustibimo/fave"
	"github.com/Gustibimo/fave/api/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "0341"
	dbname   = "fave_merchant"
	charset  = "utf8"
)

func main() {
	db, err := sqlx.Connect("mysql", user+":"+password+"@/"+dbname+"?charset="+charset)
	if err != nil {
		panic(err)
	}

	fmt.Printf("connected")
	defer db.Close()

	urlCity := "https://myfave.com/api/mobile/cities"
	var cities []string
	cities = extract.GetCity(urlCity)

	// var d model.Details
	urlAPI := "https://myfave.com/api/mobile/search/outlets?&limit=100&city="

	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
	}

	for _, city := range cities {
		fmt.Println(urlAPI + city)
		req, err := http.NewRequest(http.MethodGet, urlAPI+city, nil)
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

		var data model.Data
		json.Unmarshal([]byte(body), &data)
		for _, m := range data.Outlet {
			partnerName := strings.TrimLeft(m.Partner.Name, " ")
			// d = model.Details{m.ID, partnerName, m.Address, m.Partner.AvgRating, m.FavePayCnt, city, m.Partner.logo}

			sqlStatement := `
			INSERT INTO merchants (merchant_id, address, name, city, rating, fave_paycnt, category, logo)
						VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

			db.MustExec(sqlStatement, m.ID, m.Address, partnerName, city, m.Partner.AvgRating, m.FavePayCnt, m.Categories[0], m.Partner.Logo)
			// id := 0
			// err = db.QueryRow(sqlStatement, m.ID, m.Address, partnerName, city, m.Partner.AvgRating, m.FavePayCnt, m.Categories[0], m.Partner.Logo).Scan(&id)
			if err != nil {
				panic(err)
			}
			fmt.Println("new record for merchant ID")
		}

	}

	// 			sqlStatement := `
	// INSERT INTO merchants (merchant_id, address, name, city, rating, fave_paycnt, category, logo)
	// 			VALUES ($1, $2	, $3, $4, $5, $6, $7, $8)
	// 			RETURNING id`
	// 	id := 0
	// 	err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("New record ID is:", id)
}
