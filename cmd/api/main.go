package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// application version number
const version = "1.0.0"

// config struct holds all the configuration settings
type config struct {
	port 		int
	env 		string
}

// application struct holds the dependencies for HTTP handlers, helpers, and middlewares
type application struct {
	config 		config
	infoLog		*log.Logger
	errorLog	*log.Logger
}

func main () {
	// cfg instance of the config struct
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()


	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags | log.Lshortfile)


	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	infoLog.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()

	if err != nil {
		errorLog.Fatal(err)
	}
}