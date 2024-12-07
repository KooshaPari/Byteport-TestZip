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
		w.Header().Set("Content-Type", "application/zip")
        w.WriteHeader(http.StatusOK)
		w.Write(testFile.Bytes())
        
	})
}
/*************  ✨ Codeium Command ⭐  *************/
// createMockZip reads a file into an in-memory ZIP archive. It returns a Buffer containing
// the ZIP archive and a nil error if successful. If there is an error reading the file,
// getting the file info, or reading from the file, it returns a nil Buffer and an error.
/******  8456d9cd-5e68-4779-b2de-ec56d67f63c9  *******/
func createMockZip(filePath string) (*bytes.Buffer, error) {
	// Create an in-memory ZIP archive
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

	return buf, nil
}

func main() {}
