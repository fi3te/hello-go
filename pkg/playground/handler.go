package playground

import (
	"fmt"
	"net/http"
)

func RegisterHttpHandler() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hi")

	values := []string{}
	for i := 0; i < 10000000; i++ {
		values = append(values, "Increase memory consumption")
	}
}
