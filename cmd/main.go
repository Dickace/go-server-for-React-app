package main

import (
	"awesomeProject/pkg/transport"
	"net/http"
)

func main() {
	r := transport.Router()
	http.ListenAndServe("8000", r)

}
