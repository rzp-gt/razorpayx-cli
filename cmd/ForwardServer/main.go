package main

import (
	"fmt"
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"io/ioutil"
	"net/http"
	"os"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// create response binary data
	data, err := ioutil.ReadAll(req.Body)

	// write `data` to response
	result := ansi.ColorizeJSON(string(data), false, os.Stdout)
	fmt.Println(result)

	if err != nil {
		msg := "error"
		res.Write([]byte(msg))
		res.WriteHeader(http.StatusInternalServerError)
	}

	msg := "success"
	res.Write([]byte(msg))

}

func main() {
	fmt.Printf("Listening Now \n")

	// create a new handler
	handler := HttpHandler{}
	http.HandleFunc("/webhooks", handler.ServeHTTP)

	// listen and serve
	http.ListenAndServe(":8000", handler)
}
