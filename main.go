package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}
}
