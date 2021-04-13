package transport

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/valuateHistory", helloWorld).Methods(http.MethodGet)
	return logMiddleware(r)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	requS, err:= ioutil.ReadAll(r.Body)
	log.Print(requS)
	if err != nil{
		log.Fatal(err)
	}
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})

}
