package handlers

import (
	// "encoding/json"
	"log"
	"microserices_go/data"
	"net/http"
)

type Products struct{
	l*log.Logger
}

func NewProducts(l*log.Logger)*Products{
	return &Products{l}
}

func(p*Products) ServeHTTP(rw http.ResponseWriter, r*http.Request){
	if r.Method == http.MethodGet{
		p.GetProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*Products) GetProducts(rw http.ResponseWriter, h*http.Request){
	lp:= data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "unable to marshal the json", http.StatusInternalServerError)
	}
}