package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type config struct {
	PORT string
}

var cfg config

func main() {

	flag.StringVar(&cfg.PORT, "port", "8080", "Port where the server listens foi requests")
	flag.Parse()

	s := http.Server{
		Addr: fmt.Sprintf(":%s", cfg.PORT),
	}

	log.Printf("Server Listening on port %s \n", cfg.PORT)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill)
	signal.Notify(c, os.Interrupt)

	sig := <-c
	log.Println("Got Signal: ", sig)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	s.Shutdown(ctx)
}
