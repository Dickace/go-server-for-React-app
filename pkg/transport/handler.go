package transport

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/schema"
	"time"
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
	err := decoder.Decode(&valuteStruct, r.URL.Query())
	if err != nil{
		log.Fatal(err)
	} else {
		MakeRequest(valuteStruct.Valute)
	}
}
func MakeRequest(nameValute string) {
	var date = time.Now()
	var prevDate = time.Date(date.Year(),date.Month(),date.Day() - 20, date.Hour(),date.Minute(),date.Second(),date.Nanosecond(),date.Location())
	var url = "http://www.cbr.ru/scripts/XML_dynamic.asp?date_req1=&"+ prevDate.Format("2/1/2006")+"&date_req2=" + date.Format("2/1/2006") + "&VAL_NM_RQ="+ nameValute
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(url)
	log.Println(body)
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
