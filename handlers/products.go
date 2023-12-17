package handlers

import (
	// "encoding/json"
	"log"
	"microserices_go/data"
	"net/http"
	"regexp"
	"strconv"
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
	if r.Method == http.MethodPost{
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut{
		p.l.Println("PUT", r.URL.Path)
		
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) !=1{
			p.l.Println("invalid uri ", http.StatusBadRequest)
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(g[0]) !=2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil{
			p.l.Println("Invalid URI unable to convert to number", idString)
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, r)
	
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

func (p*Products) addProduct(rw http.ResponseWriter, r*http.Request){
	p.l.Println("Handle Get Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil{
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p Products) updateProducts(id int, rw http.ResponseWriter, r*http.Request) {
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