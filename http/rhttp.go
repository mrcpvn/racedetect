package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	c  *int
	hf = func(w http.ResponseWriter, r *http.Request) {
		c = updateCounter(c)
		fmt.Printf("&c = %v, c = %v\n", c, *c)
		fmt.Fprintf(w, `{"counter":%v}`, *c)
	}

	rf = func(w http.ResponseWriter, r *http.Request) {
		if c == nil {
			zero := 0
			c = &zero
		}
		fmt.Fprintf(w, `{"counter":%v}`, *c)
	}
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/counter", hf).Methods("POST")
	r.HandleFunc("/counter", rf).Methods("GET")

	http.Handle("/", handlers.LoggingHandler(os.Stdout, r))

	http.ListenAndServe(":8080", nil)
}

func updateCounter(i *int) *int {
	if i != nil {
		*i = *i + 1
		return i
	}
	zero := 0
	fmt.Printf("i = %v\n", i)
	i = &zero
	return i
}
