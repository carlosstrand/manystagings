package main

import (
	"github.com/carlosstrand/manystagings/app"
	"github.com/carlosstrand/manystagings/seeds"
)

func main() {
	db, err := app.CreateDB()
	if err != nil {
		panic(err)
	}
	err = seeds.RunSeeds(db, seeds.DEFAULT)
	if err != nil {
		panic(err)
	}
}
