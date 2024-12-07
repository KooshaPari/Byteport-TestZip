package main

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/tools/go/expect"
)

func TestZipReader(t *testing.T){
	fmt.Println("starting test...")
	testFile := createMockZip()
	t.Logf("Created mock zip file: %s", testFile)
	fileList, fileMap, err := ReadZip(bytes.NewReader(testFile.Bytes()), int64(testFile.Len()))
	if err != nil {
		t.Errorf("failed to read zip file: %w", err)
	}
	if(len(fileList) != 1) {
		t.Errorf("expected 1 file, got %d", len(fileList))
	}
	t.Logf("Read zip file successfully with %d files", len(fileList))
	t.Logf("File map: %v", fileMap)
	expect.Equal(t, fileMap["README.md"], []byte("Don't actually read me"))
	
	t.Logf("Full process completed successfully")
}
	/*req, err := http.NewRequest("GET","127.0.0.1:3000/",nil)
	if err != nil {
		 t.Errorf("failed to create request: %w", err)
	}
	resp, err := spinhttp.Send(req)
	if err != nil {
		 t.Errorf("failed to fetch archive: %w", err)
	}
	response := map[string]interface{}{
            "status": "success",
            "files": []string{},
            "fileMap": nil,
            "fileCount": 0,
        }
	// list out 
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		 t.Errorf("failed to read response body: %w", err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		 t.Errorf("failed to unmarshal response: %w", err)
	}*/
	t.Logf("Full process completed successfully with response: %d", /*(response["fileCount"].(float64))*/ 2)
}

