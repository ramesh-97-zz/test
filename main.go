package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		log.Println("Hello Ramesh")
		d, err := ioutil.ReadAll(r.Body)
		//log.Printf("Data %s\n", d)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			/*rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Oopps"))*/
			return
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {

		log.Println("Hello Rams")
	})

	http.ListenAndServe(":9090", nil)
}
