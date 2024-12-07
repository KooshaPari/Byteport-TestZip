package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"path"
	"strings"

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
func ReadZip(zipBytes []byte) ([]string, map[string][]byte) {
        


        zipReader, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
        if err != nil {
            fmt.Sprintf("Failed to read zip archive: %v", err)
        }
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
