package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type doConnection interface {
	Do(req *http.Request) (*http.Response, error)
}

type noConnection struct{}

func (c noConnection) Do(req *http.Request) (*http.Response, error) {
	response := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("Hello World"))),
	}
	return &response, nil
}

func main() {
	// Refactor the following code, such that we can call the getTheDocs function
	// without making an actual http request. Hint: can we define an interface for
	// the behaviour of the httpClient and pass it to the function instead of the http client struct?
	// Write a unit test for the getTheDocs function, where we implement the interface with a
	// struct that doesn't actually makes an http call.
	// httpClient := http.DefaultClient
	httpClient := &noConnection{}
	goDocsRespBody := getTheDocs(httpClient)
	fmt.Println(goDocsRespBody)
}

func getTheDocs(client doConnection) string {
	req, err := http.NewRequest(http.MethodGet, "https://go.dev/doc/", nil)
	if err != nil {
		// we did not learn about errors yet, so ignore for now
		return err.Error()
	}
	resp, err := client.Do(req)
	if err != nil {
		// we did not learn about errors yet, so ignore for now
		return err.Error()
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		// we did not learn about errors yet, so ignore for now
		return err.Error()
	}
	return string(b)
}
