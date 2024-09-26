package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const serverPort = 4200

var random *rand.Rand

func main() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))

	initMetrics()

	http.Handle("/", instrumentHandlerDuration(handleHome))
	http.Handle("/hello", instrumentHandlerDuration(handleHello))
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Running server on port %d...\n", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	millis := sleep(50)
	fmt.Fprintf(w, "Home %d!", millis)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	millis := sleep(150)
	fmt.Fprintf(w, "Hello after %d milliseconds!", millis)
}

func sleep(maxMillis int) int {
	millis := random.Intn(maxMillis + 1)
	time.Sleep(time.Millisecond * time.Duration(millis))
	return millis
}
