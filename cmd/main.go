package main

import (
	"log"
	"os"
)

//import "fmt"

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("server has failed to start %s", err)
		os.Exit(1)
	}

}
