package main

import (
	"net/http"
)

func main() {
	http.Handle("/sr/", http.StripPrefix("/sr/", http.FileServer(http.Dir("/home/eranl/go/src/github.com/elerer/goserve"))))
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
