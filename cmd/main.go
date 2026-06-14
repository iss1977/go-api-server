package main

import "net/http"

//import "fmt"

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	app := &application{}
	handler := app.mount()
	http.ListenAndServe(":3000", handler)

}
