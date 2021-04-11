package main

import (
	"github.com/carlosstrand/manystagings/app"
)

func main() {
	a := app.NewApp(app.Options{})
	a.Start()
}
