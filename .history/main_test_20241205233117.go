package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestZipReader(t *testing.T){
	fmt.Println("starting test...")
	// set as local file ./odin-dash.zip
	

	resp, err := spinhttp.Get("http://localhost:8080/")
	// Process the HTTP response
	fileList, fileMap, err := processZipResponse(resp)
	if err != nil {
		t.Fatalf("Failed to process ZIP from HTTP response: %v", err)
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
	


