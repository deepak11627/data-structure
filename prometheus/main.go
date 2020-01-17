package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"time"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "request_counter",
			Help:      "This is my counter",
		}, []string{"code", "method", "path"})
)

func main() {

	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(requestCounter)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print(fmt.Sprintf("Recieved request on %s %s", r.URL.Path, time.Now()))
	if rand.Intn(100)%2 == 0 {
		requestCounter.WithLabelValues("200", "GET", r.URL.Path).Add(1)
		fmt.Fprint(w, fmt.Sprintf("You requested page %s", r.URL.Path))
	} else {
		requestCounter.WithLabelValues("500", "GET", r.URL.Path).Add(1)
		w.WriteHeader(500)
	}

}
