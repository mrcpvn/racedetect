package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRace(t *testing.T) {
	res := run()
	if res != 10 {
		t.Errorf("concurrent count wrong: expected 10, got %v", res)
	}
}

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(hf))
	defer ts.Close()

	rch := make(chan (int))
	fmt.Printf("test url = %v\n", ts.URL)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			r, err := http.DefaultClient.Get(ts.URL)
			if err != nil {
				fmt.Printf("error calling test server: %v\n", err)
			}
			defer r.Body.Close()
			b, _ := ioutil.ReadAll(r.Body)
			fmt.Printf("%s\n", b)
			wg.Done()
			rch <- 0
		}()
	}

	result := 0

	for range rch {
		result++
		if result == 10 {
			close(rch)
		}
	}

	wg.Wait()
	fmt.Printf("result = %v\n", result)
}
