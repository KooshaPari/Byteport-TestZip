package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)


func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getting files...")
		testFile,err := createMockZip("./odin-dash.zip")
		if err != nil {
			http.Error(w, "Failed to create mock ZIP file", http.StatusInternalServerError)
			return
		}
		fmt.Println("got files")
		w.Header().Set("Content-Type", "application/zip")
        w.WriteHeader(http.StatusOK)
		w.Write(testFile.Bytes())
        
	})
}
func createMockZip(filePath string) (*bytes.Buffer, error) {
	// Create an in-memory ZIP archive
	fmt.Println("creating mock zip")
	// check if file exists
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Add a mock file to the ZIP archive
	file, _ := zipWriter.Create("README.md")
	file.Write([]byte("Don't actually read me"))

	// Close the ZIP writer to finalize the archive
	zipWriter.Close()

	return buf
}
}

func main() {}
