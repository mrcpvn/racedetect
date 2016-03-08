package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

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
)

func main() {
	s := run()
	fmt.Printf("concurrent count = %v\n", s)

	fmt.Printf("c = %v\n", c)
	r := mux.NewRouter()

	r.HandleFunc("/users/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "{Hello %s!}", vars["name"])
	}).Methods("GET")

	r.HandleFunc("/counter", hf).Methods("GET")

	http.Handle("/", handlers.LoggingHandler(os.Stdout, r))

	http.ListenAndServe(":8080", nil)
}

func run() int {
	var mux sync.Mutex
	var wg sync.WaitGroup
	race := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mux.Lock()
			race++
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return race
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
