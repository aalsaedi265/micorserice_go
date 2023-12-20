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

	"github.com/gorilla/mux"
)

func main() {
	
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)

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
	//import to graceful shut down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c //blocks and waits for a signal to be sent to the channe
	log.Println("Got signal: ", sig)
	// http.ListenAndServe(":9090", sm)
	tc,_ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}