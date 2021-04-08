package main

import (
	"awesomeProject/pkg/transport"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("port")

	r := transport.Router()
	http.ListenAndServe(":"+ port, r)

}
