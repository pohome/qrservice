package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
	"os"
)

var getMap map[string]http.HandlerFunc = map[string]http.HandlerFunc{
	"/qr": getQR}

func main() {
	mux := routes.New()
	for k, v := range getMap {
		mux.Get(k, v)
	}

	http.Handle("/", mux)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	os.Exit(0)
}
