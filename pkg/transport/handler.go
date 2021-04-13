package transport

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/schema"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/valuateHistory", helloWorld).Methods(http.MethodGet)
	return logMiddleware(r)
}
type ValuteStruct struct {
	Valute string
}

var decoder = schema.NewDecoder()

func helloWorld(w http.ResponseWriter, r *http.Request) {
	var valuteStruct ValuteStruct
	var valute = r.URL.Query()
	err := decoder.Decode(&valuteStruct, r.URL.Query())
	if err != nil{
		log.Fatal(err)
	} else {
		log.Println(valuteStruct)
		log.Println(valute)
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
