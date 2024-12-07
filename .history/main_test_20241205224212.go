package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestZipReader(t *testing.T){
	fmt.Println("starting test...")
	// set as local file ./odin-dash.zip
	testFile,err := createMockZip("./odin-dash.zip")
	if err != nil {
		t.Errorf("failed to create mock zip file: %w", err)
	}
	//t.Logf("Created mock zip file: %s", testFile)
	fileList, fileMap, err := ReadZip(bytes.NewReader(testFile.Bytes()), int64(testFile.Len()))
	if err != nil {
		t.Errorf("failed to read zip file: %w", err)
	}
	if(len(fileList) != 50) {
		t.Errorf("expected 1 file, got %d", len(fileList))
	}
	t.Logf("Read zip file successfully with %d files", len(fileList))
	t.Logf("File map: %v", fileMap)
	expectedValue := "Don't actually read me"
	actualValue :=
	if string(fileMap["README.md"]) != expectedValue {
		t.Errorf("expected %s, got `%s`", expectedValue, fileMap["KooshaPari-odin-dash-e628e86/README.md"])
	}
	t.Logf("Readme Content correct: `%s`", fileMap["README.md"])
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
	


