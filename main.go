package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type config struct {
	PORT string
	l    log.Logger
}

var cfg config

func main() {

	s := http.Server{
		Addr: fmt.Sprintf(":%s", cfg.PORT),
	}

	go func() {
		cfg.l.Println(fmt.Sprintf("Server Listening on port %s", cfg.PORT))
		err := s.ListenAndServe()
		if err != nil {
			cfg.l.Fatal(err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Kill)
	signal.Notify(c, os.Interrupt)

	sig := <-c
	cfg.l.Println("Got Signal ", sig)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	s.Shutdown(ctx)
}
