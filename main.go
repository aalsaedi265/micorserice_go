package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// r for read
	http.HandleFunc("/dio", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("the world")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(rw, "wwarryyy", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "data %s\n", data)
	})
	http.ListenAndServe(":9090", nil)

}