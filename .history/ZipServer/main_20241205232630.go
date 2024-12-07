package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		testFile,err := Byteport.createMockZip("./odin-dash.zip")
	if err != nil {
		t.Errorf("failed to create mock zip file: %w", err)
	}
		w.Header().Set("Content-Type", "application/zip")
        w.WriteHeader(http.StatusOK)
        
	})
}

func main() {}
