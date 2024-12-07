package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestZipReader(t *testing.T){
	fmt.Println("starting test...")
	// set as local file ./odin-dash.zip
	

	/*resp, err := spinhttp.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Failed to get HTTP response: %v", err)
	}
	
	// Process the HTTP response
	fileList, fileMap, err := processZipResponse(resp)
	if err != nil {
		t.Fatalf("Failed to process ZIP from HTTP response: %v", err)
	}*/
	testFile,err := createMockZip("./odin-dash.zip")
	if err != nil {
		t.Errorf("failed to create mock zip file: %w", err)
	}
	fileList, fileMap, err := ReadZip(testFile)
	if err != nil {
		t.Errorf("failed to process zip file: %w", err)
	}
	if len(fileList) == 0 {
		t.Errorf("expected non-empty file list, got %v", fileList)
	}
	if len(fileMap) == 0 {
		t.Errorf("expected non-empty file map, got %v", fileMap)
	}
	
	rootDir,err := getRootDir(fileMap)
	if err != nil {
		t.Errorf("failed to get root directory: %w", err)
	}

	
	t.Logf("Read zip file successfully with %d files", len(fileList))
	//t.Logf("File map: %v", fileMap)
	expectedValue := "Don't actually read me"
	actualValue := strings.TrimSpace(string(fileMap[rootDir+"README.md"]))
	if actualValue != expectedValue {
		t.Errorf("expected %s, got `%s`", expectedValue, actualValue)
	}
	t.Logf("Readme Content correct: `%s`", actualValue)
	t.Logf("Full process completed successfully")
}
	


