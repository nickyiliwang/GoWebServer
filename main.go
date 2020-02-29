package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// reminder: you need to go build to test changes
	// route
	http.HandleFunc("/welcome", welcome)

	// static assets
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the jungle!")
}
