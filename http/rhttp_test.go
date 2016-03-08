package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(hf))
	defer ts.Close()

	rch := make(chan (int))
	fmt.Printf("test url = %v\n", ts.URL)
	expNum := 10
	var wg sync.WaitGroup
	for i := 0; i < expNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := http.DefaultClient.Post(ts.URL, "", nil)
			if err != nil {
				fmt.Printf("error calling test server: %v\n", err)
			}
			defer r.Body.Close()
			b, _ := ioutil.ReadAll(r.Body)
			fmt.Printf("%s\n", b)
			rch <- 0
		}()
	}
	result := 0

	for range rch {
		result++
		if result == expNum {
			close(rch)
		}
	}

	wg.Wait()
	fmt.Printf("result = %v\n", result)
	//check value
	tc := httptest.NewServer(http.HandlerFunc(rf))
	defer tc.Close()

	r, err := http.DefaultClient.Get(tc.URL)
	if err != nil {
		fmt.Printf("error calling check server: %v\n", err)
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s\n", b)
	expected := fmt.Sprintf(`{"counter":%v}`, expNum)
	if string(b) != expected {
		t.Errorf("Error concurrent count: expected %v, got %s", expected, b)
	}
}
