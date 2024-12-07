package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)



func init() {
    spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
       /* fmt.Println("Received request")


        zipResp, err := DownloadRepo(w, r)
        if err != nil {
            fmt.Println("Error downloading zip file: ", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        } 
        fmt.Println("Processing zip file...")
        zipReader, err := ProcessZip(zipResp, w,r)
        if err != nil {
            fmt.Println("Error processing zip file: ", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

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
        }*/
    })
}
func createMockZip(filepath string) (*bytes.Buffer, error) {
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

	return fileList, fileMap, nil}

/*
func DownloadRepo(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
    // Use Spin's HTTP client for the initial request
        req,err := http.NewRequest("GET", archiveURL, nil)
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to create request: %v", err), http.StatusInternalServerError)
         
        }
        req.Header.Add("User-Agent", "Byteport") // Set user agent  
        //req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken)) // Add token to headers
        req.Header.Add("Accept", "application/vnd.github+json") // Set accept header to GitHub API
        //fmt.Println("HEAD: ", req.Header)
        // print out full request
        //fmt.Println("Request: ", req)
        resp, err := spinhttp.Send(req)
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to get archive: %v", err), http.StatusInternalServerError)
        
        }
        defer resp.Body.Close()
        fmt.Println("Checking for Redirect...")
        fmt.Println("Status: ", resp.StatusCode)
        fmt.Println("Response: ", resp)
        // Check for redirect
        var zipResp *http.Response
        if resp.StatusCode != http.StatusFound && resp.StatusCode != http.StatusOK {
            http.Error(w, fmt.Sprintf("Expected redirect or file, got: %s", resp.Status), resp.StatusCode)
           
        }
        if(resp.StatusCode == http.StatusFound){
            fmt.Println("Found, Redirecting...")
            // Get redirect URL from Location header
            redirectURL := resp.Header.Get("Location")
            if redirectURL == "" {
                http.Error(w, "No redirect URL provided", http.StatusInternalServerError)
              
            }

            fmt.Println("Redirect URL: ", redirectURL)
            redirReq, err := http.NewRequest("GET", redirectURL, nil)
            zipResp, err := spinhttp.Send(redirReq)
            if err != nil {
                http.Error(w, fmt.Sprintf("Failed to download zip: %v", err), http.StatusInternalServerError)
             
            }
            defer zipResp.Body.Close()

            if zipResp.StatusCode != http.StatusOK {
                http.Error(w, fmt.Sprintf("Failed to download zip, status: %s", zipResp.Status), zipResp.StatusCode)
               
            }
        }else{
            zipResp = resp
        }
        return zipResp, nil
}*/
func main() {}
