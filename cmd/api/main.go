package main

import (
	"github.com/kRhodes-interfolio/social/internal/env"
	"log"
)

func main() {
	app := &application{
		config{
			addr: env.GetString("ADDR", ":8080"),
		},
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
