package main

import (
	"fmt"
	"log"
	"net/http"
)

// Try starting the program and execute an HTTP request using curl: curl localhost:3000/home
func main() {
	fmt.Println("Start HTTP server")

	http.HandleFunc("/home", trackUserAgent(homepageHandler))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Define a function returning a closure, which prints the user agent of the request.
// Hint: both the input parameter and return type of the function should adhere to the http Handler function which
// looks as follows: func(http.ResponseWriter, *http.Request)

func trackUserAgent(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Useragent: ", request.UserAgent())
		fmt.Fprintln(writer, request.UserAgent())
		handler(writer, request)
	}
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from server")
	fmt.Fprintln(w, "I am learning about closures")
}
