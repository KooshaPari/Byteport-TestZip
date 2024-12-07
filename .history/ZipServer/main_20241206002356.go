package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"net/http"

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
    fmt.Println("creating mock zip")
    buf := new(bytes.Buffer)
    zipWriter := zip.NewWriter(buf)
    fmt.Println("created mock zip")

    // Add a mock file to the ZIP archive
    file, err := zipWriter.Create("README.md")
    if err != nil {
        return nil, err
    }

    _, err = file.Write([]byte("Don't actually read me"))
    if err != nil {
		fmt.Println("failed to write file")
        return nil, err
    }
    fmt.Println("added mock file")

    // Close the ZIP writer to finalize the archive
    err = zipWriter.Close()
    if err != nil {
        return nil, err
    }

    return buf, nil
}

func main() {}
