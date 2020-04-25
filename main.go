package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello\n")
}

func header(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/headers", header)
	http.ListenAndServe(":8080", nil)
}
