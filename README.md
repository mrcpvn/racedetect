# racedetect
race detector examples

### detect race condition at runtime

  cd http
  
  build -race rhttp.go
  
  ./rhttp
  
  ab -c 10 -n 10 -m POST http://127.0.0.1:8080/counter
