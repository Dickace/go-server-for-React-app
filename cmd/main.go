package main

import (
	"awesomeProject/pkg/transport"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const serverUrl = ":8000"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.WithFields(log.Fields{"url": serverUrl}).Info("server start")
	r := transport.Router()
	fmt.Println(http.ListenAndServe(serverUrl, r))

}
