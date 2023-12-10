package main

import (
	// "fmt"
	// "io/ioutil"
	"context"
	"log"
	"microserices_go/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120* time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	
	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c //blocks and waits for a signal to be sent to the channe
	log.Println("Got signal: ", sig)
	// http.ListenAndServe(":9090", sm)
	tc,_ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}