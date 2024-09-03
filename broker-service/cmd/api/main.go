package main

import (
	"fmt"
	"log"
	"net/http"
)

const wepPort = "80"

type Config struct {
}

func main() {
	app := Config{}
	log.Printf("Starting broker service on port %s\n", wepPort)

	//define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", wepPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}
