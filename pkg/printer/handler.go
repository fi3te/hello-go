package printer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func RegisterHttpHandler() {
	http.HandleFunc("/printer", handle)
}

func handle(w http.ResponseWriter, req *http.Request) {
	b, err := httputil.DumpRequest(req, true)
	if err == nil {
		fmt.Println(string(b))
	}
	w.WriteHeader(http.StatusBadRequest)
}
