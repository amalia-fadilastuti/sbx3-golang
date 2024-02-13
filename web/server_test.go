package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumHandler(t *testing.T) {
	// Prepare the HTTP Request
	req, err := http.NewRequest("GET", "http://localhost:8080/sum?a=1&b=1", nil)

	require.NoErrorf(t, err, "failed to create request: %v", err)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Record the HTTP Response
	rec := httptest.NewRecorder()

	// Testing the sumHandler.
	sumHandler(rec, req)

	resp := rec.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("response status code is not 200 OK: %v", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	defer resp.Body.Close()

	s, err := strconv.Atoi(string(b))
	if err != nil {
		t.Fatalf("failed to convert body to integer: %v", err)
	}

	if s != 2 {
		t.Errorf("one plus one should be 2, got: %v", s)
	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	resp, err := http.Get(fmt.Sprintf("%s/sum?a=1&b=1", srv.URL))
	if err != nil {
		t.Fatalf("failed to send request to /sum: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("response status code is not 200 OK: %v", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	defer resp.Body.Close()

	s, err := strconv.Atoi(string(b))
	if err != nil {
		t.Fatalf("failed to convert body to integer: %v", err)
	}

	if s != 2 {
		t.Errorf("one plus one should be 2, got: %v", s)
	}

}
