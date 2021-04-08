package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Hi World!")
	})

	http.ListenAndServe(":"+port, nil)

}
