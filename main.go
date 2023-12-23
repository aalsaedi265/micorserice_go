package main

import (
	// "fmt"
	// "io/ioutil"
	"context"
	"log"
	"microserices_go/data"
	"microserices_go/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/tree/master/middleware"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "Go_Microservices", log.LstdFlags)
	v := data.NewValidation()

	ph := handlers.NewProducts(l, v)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", ph.ListAll)
	getR.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/products", ph.Update)
	putR.Use(ph.MiddlewareValidateProduct)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/products", ph.Create)
	postR.Use(ph.MiddlewareValidateProduct)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
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