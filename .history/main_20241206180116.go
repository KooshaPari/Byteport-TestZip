package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)



func init() {
    spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("starting test...")
	resp, err := spinhttp.Get("/home")
	if err != nil {
		fmt.Printf("Failed to get HTTP response: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Got response from server")
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("unexpected status code: %d", resp.StatusCode), http.StatusInternalServerError)
	fmt.Println("Bad Status Code")
	return}

	fileList, fileMap, err := processZipResponse(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Got file list and file map")
	rootDir,err := getRootDir(fileMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)}
	fmt.Println("Read zip file successfully with %d files", len(fileList))
	//fmt.Println("File map: %v", fileMap)
	expectedValue := "Don't actually read me"
	actualValue := strings.TrimSpace(string(fileMap[rootDir+"README.md"]))
	if actualValue != expectedValue {
		http.Error(w, fmt.Sprintf("expected %s, got `%s`", expectedValue, actualValue), http.StatusInternalServerError)
	}
	fmt.Println("Readme Content correct: `%s`", actualValue)
	fmt.Println("Full process completed successfully")
    })
}
func getRootDir(fileMap map[string][]byte) (string, error) {
	for key := range fileMap {
		// Split the key into parts using "/" as the delimiter
		parts := strings.Split(key, "/")
		if len(parts) > 1 {
			// Return the first part (root directory)
			return parts[0] + "/", nil
		}
	}
	return "", fmt.Errorf("no valid root directory found")
}

func ReadZip(reader io.ReaderAt, size int64) ([]string, map[string][]byte, error) {
	zipReader, err := zip.NewReader(reader, size)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read ZIP archive: %w", err)
	}

	fileList := []string{}
	fileMap := make(map[string][]byte)

	for _, file := range zipReader.File {
		fileList = append(fileList, file.Name)
		f, err := file.Open()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file %s: %w", file.Name, err)
		}

		content, err := io.ReadAll(f)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read file %s: %w", file.Name, err)
		}
		fileMap[file.Name] = content
		f.Close()
	}

	return fileList, fileMap, nil
}

func processZipResponse(resp *http.Response) ([]string, map[string][]byte, error) {
	defer resp.Body.Close()

	// Read the response body into memory
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Process the ZIP file
	reader := bytes.NewReader(body)
	return ReadZip(reader, int64(len(body)))
}
func createMockZip(filePath string) (*bytes.Buffer, error) {
    fmt.Println("creating mock zip")
    buf := new(bytes.Buffer)
    zipWriter := zip.NewWriter(buf)
    fmt.Println("created mock zip")

    // Add a mock file to the ZIP archive
    file, err := zipWriter.Create("README.md")
    if err != nil {
		fmt.Println("failed to create file: ", err)
        return nil, err
    }

    _, err = file.Write([]byte("Don't actually read me"))
    if err != nil {
		fmt.Println("failed to write file: ", err)
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