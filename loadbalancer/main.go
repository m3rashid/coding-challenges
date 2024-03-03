package main

import (
	"fmt"
	"net/http"
)

var onlineServers = []string{}

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request URL:", r.URL)
		w.Write([]byte(""))

	})

	http.ListenAndServe(":4000", nil)
}
