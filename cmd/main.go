package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fi3te/hello-go/pkg/playground"
	"github.com/fi3te/hello-go/pkg/printer"
	"github.com/fi3te/hello-go/pkg/weather"
)

const serverPort = 4200

func main() {
	playground.RegisterHttpHandler()
	weather.RegisterHttpHandler()
	printer.RegisterHttpHandler()
	err := http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err.Error())
		os.Exit(1)
	}
}
