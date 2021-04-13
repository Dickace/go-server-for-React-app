package transport

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/valuateHistory", helloWorld).Methods("GET", "OPTIONS")
	return logMiddleware(r)
}
type ValuteStruct struct {
	Valute string
}

var decoder = schema.NewDecoder()

func helloWorld(w http.ResponseWriter, r *http.Request) {
	var valuteStruct ValuteStruct
	err := decoder.Decode(&valuteStruct, r.URL.Query())
	if err != nil{
		log.Fatal(err)
	} else {
		var body = MakeRequest(valuteStruct.Valute)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, string(body))
	}
}
func MakeRequest(nameValute string) []byte {
	var date = time.Now()
	var prevDate = time.Date(date.Year(),date.Month(),date.Day() - 20, date.Hour(),date.Minute(),date.Second(),date.Nanosecond(),date.Location())
	var url = "http://www.cbr.ru/scripts/XML_dynamic.asp?date_req1="+ prevDate.Format("02/01/2006")+"&date_req2=" + date.Format("02/01/2006") + "&VAL_NM_RQ="+ nameValute
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(url)
	return body
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
