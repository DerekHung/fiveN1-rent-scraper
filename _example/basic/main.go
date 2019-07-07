package main

import (
	"log"

	rent "github.com/neighborhood999/fiveN1-rent-scraper"
)

func main() {
	o := rent.NewOptions()

	o.Region = 3
	o.MRTCoods = `4213,4212,4211,4215,4214`
	o.RentPrice = `15000, 27000`
	o.Area = `20,70`
	o.Shape = 2
	o.Other = `cook`
	o.Order = `posttime`
	o.OrderType = `desc`


	url, err := rent.GenerateURL(o)
	if err != nil {
		log.Fatalf("\x1b[91;1m%s\x1b[0m", err)
	}

	f := rent.NewFiveN1(url)
	if err := f.Scrape(1); err != nil {
		log.Fatal(err)
	}

	json := rent.ConvertToJSON(f.RentList)
	log.Println(string(json))
	// rent.ExportJSON(json)
}
