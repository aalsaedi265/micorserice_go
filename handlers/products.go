package handlers

import (
	// "encoding/json"
	"log"
	"microserices_go/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct{
	l*log.Logger
}

func NewProducts(l*log.Logger)*Products{
	return &Products{l}
}

func (p*Products) GetProducts(rw http.ResponseWriter, h*http.Request){
	lp:= data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "unable to marshal the json", http.StatusInternalServerError)
	}
}

func (p*Products) AddProduct(rw http.ResponseWriter, r*http.Request){
	p.l.Println("Handle Get Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil{
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p Products) UpdateProducts( rw http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}