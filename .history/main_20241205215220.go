package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"path"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)



func init() {
    spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Received request")

        fmt.Println("Reading zip file...")
        fileList, fileMap := ReadZip(zipReader)

        fmt.Println("Extracted files:")
        response := map[string]interface{}{
            "status": "success",
            "files": fileList,
            "fileMap": fileMap,
            "fileCount": len(fileList),
        }

        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, "Failed to encode response", http.StatusInternalServerError)
            return
        }
    })
}
func ReadZip(zipReader *zip.Reader) ([]string, map[string][]byte) {
    
        fileMap := make(map[string][]byte)
        var fileList []string
        for _, file := range zipReader.File {
            if file.FileInfo().IsDir() {
                continue
            }

            parts := strings.Split(file.Name, "/")
            if len(parts) > 1 {
                parts = parts[1:]
            }
            relativePath := path.Join(parts...)
            fmt.Printf("File: %s, Method: %d\n", file.Name, file.Method)

            if file.Method != zip.Deflate && file.Method != zip.Store {
                fmt.Printf("Warning: Unsupported compression method %d for file %s\n", file.Method, file.Name)
                continue
            }
            rc, err := file.Open()
            if err != nil {
                fmt.Printf("Warning: Could not open file %s: %v\n", file.Name, err)
                continue
            }

            content, err := io.ReadAll(rc)
            rc.Close()
            
            if err != nil {
                fmt.Printf("Warning: Could not read file %s: %v\n", file.Name, err)
                continue
            }

            fileMap[relativePath] = content
            fileList = append(fileList, relativePath)
        }
        return fileList, fileMap;

}
func ProcessZip(zipResp *http.Response, w http.ResponseWriter, r *http.Request) (*zip.Reader, error) {
    fmt.Println("Processing zip file...")
     zipBytes, err := io.ReadAll(zipResp.Body)
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to read zip file: %v", err), http.StatusInternalServerError)
        
        }

        if len(zipBytes) == 0 {
            http.Error(w, "Received empty zip file", http.StatusInternalServerError)
           
        }

        zipReader, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to read zip archive: %v", err), http.StatusInternalServerError)
          
        }
        return zipReader, nil
}

func createMockZip() *bytes.Buffer {
	// Create an in-memory ZIP archive
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Add a mock file to the ZIP archive
	file, _ := zipWriter.Create("README.md")
	file.Write([]byte("Don't actually read me"))

	// Close the ZIP writer to finalize the archive
	zipWriter.Close()

	return buf
}


func main() {}
