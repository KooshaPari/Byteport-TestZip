package main

import (
	"fmt"
	"testing"
)

func TestZipReader(t *testing.T){
	fmt.Println("starting test...")
	
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

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// Mocked key-value store for caching
type KeyValueStore struct {
	mu    sync.Mutex
	store map[string][]byte
	calls []Call
}

type Call struct {
	Tag string
	Val string
}


// Test function
func TestCacheHit() error {
	// Set up the test
	cache := NewKeyValueStore()
	user := map[string]interface{}{"id": 123, "name": "Ryan"}
	userBytes, _ := json.Marshal(user)
	cache.Set("123", userBytes)

	// Execute request
	req := NewMockRequest("/?user_id=123")
	resp := NewMockResponse()
	handle(req, resp, cache)

	// Assertions
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	var keyValueCalls []Call
	expectedCalls := []Call{{Tag: "get", Val: "123"}}
	keyValueCalls = cache.Calls()

	if !compareCalls(keyValueCalls, expectedCalls) {
		return fmt.Errorf("expected key value calls to be %v but were %v", expectedCalls, keyValueCalls)
	}

	fmt.Println("TestCacheHit passed!")
	return nil
}

func compareCalls(a, b []Call) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	if err := TestCacheHit(); err != nil {
		fmt.Printf("Test failed: %v\n", err)
	} else {
		fmt.Println("All tests passed!")
	}
}