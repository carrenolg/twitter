package main

import (
	"log"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Not connection to DB")
		return
	}
	handlers.Handler()
}
