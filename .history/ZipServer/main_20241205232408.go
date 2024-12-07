package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/zip")
        w.WriteHeader(http.StatusOK)
        w.Write(fn"Hi")
	})
}

func main() {}
