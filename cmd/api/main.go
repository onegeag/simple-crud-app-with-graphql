package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env string
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 6060, "Graphql API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stag|prod)")

	flag.Parse()

	app := application{
		config: cfg,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
		Handler: mux,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on port %d\n", app.config.port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
