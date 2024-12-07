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
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// Read the file into a buffer
	stat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info for %s: %w", filePath, err)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s into buffer: %w", filePath, err)
	}

	if stat.Size() != int64(buf.Len()) {
		return nil, fmt.Errorf("file size mismatch for %s: expected %d bytes, got %d bytes", filePath, stat.Size(), buf.Len())
	}
	fmt.Println("mock zip created")
	return buf, nil
}

func main() {}
