package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	// r for read
	http.HandleFunc("/dio", func(http.ResponseWriter, r*http.Request){
		log.Println("the world")
		data, err := ioutil.ReadAll(r.Body)
		log.Printf(data %s, data)
	})
	http.ListenAndServe(":9090", nil)

}