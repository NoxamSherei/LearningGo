package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "web services are easy with go") //more general function to print use Writer
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})
	http.ListenAndServe(":3000", nil) //nil == null
}
