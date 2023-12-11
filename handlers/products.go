package handlers

import (
	"encoding/json"
	"log"
	"microserices_go/data"
	"net/http"
)

type Product struct{
	l*log.Logger
}

func NewProducts(l*log.Logger)*Product{
	return &Product{l}
}

func(p*Product) ServeHTTP(rw http.ResponseWriter, h*http.Request){
	lp:= data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil{
		http.Error(rw, "unable to marshal the json", http.StatusInternalServerError)
	}
	rw.Write(d)
}