package web_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amalia-fadilastuti/sbx3-golang-level4/db"

	"github.com/amalia-fadilastuti/sbx3-golang-level4/web"
)

// func TestSumHandler(t *testing.T) {
// 	// Prepare the HTTP Request
// 	req, err := http.NewRequest("GET", "http://localhost:8080/sum?a=1&b=1", nil)

// 	require.NoErrorf(t, err, "failed to create request: %v", err)
// 	if err != nil {
// 		t.Fatalf("failed to create request: %v", err)
// 	}

// 	// Record the HTTP Response
// 	rec := httptest.NewRecorder()

// 	// Testing the sumHandler.
// 	sumHandler(rec, req)

// 	resp := rec.Result()
// 	if resp.StatusCode != http.StatusOK {
// 		t.Fatalf("response status code is not 200 OK: %v", resp.StatusCode)
// 	}

// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("failed to read body: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	s, err := strconv.Atoi(string(b))
// 	if err != nil {
// 		t.Fatalf("failed to convert body to integer: %v", err)
// 	}

// 	if s != 2 {
// 		t.Errorf("one plus one should be 2, got: %v", s)
// 	}
// }

// func TestRouting(t *testing.T) {
// 	srv := httptest.NewServer(handler())
// 	defer srv.Close()

// 	resp, err := http.Get(fmt.Sprintf("%s/sum?a=1&b=1", srv.URL))
// 	if err != nil {
// 		t.Fatalf("failed to send request to /sum: %v", err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Fatalf("response status code is not 200 OK: %v", resp.StatusCode)
// 	}

// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("failed to read body: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	s, err := strconv.Atoi(string(b))
// 	if err != nil {
// 		t.Fatalf("failed to convert body to integer: %v", err)
// 	}

// 	if s != 2 {
// 		t.Errorf("one plus one should be 2, got: %v", s)
// 	}

// }

func TestRouting(t *testing.T) {
	conn, err := db.CreateConnection()
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	srv := httptest.NewServer(web.Handler(conn))
	resp, err := http.Get(srv.URL + "/department")
	if err != nil {
		t.Fatalf("failed to get department: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("failed to receive status OK, got: %v", resp.StatusCode)
	}

	var department []db.Department
	if err := json.NewDecoder(resp.Body).Decode(&department); err != nil {
		t.Fatalf("failed to decode department: %v", err)
	}

	if department[0].DepartmentId != 1 {
		t.Fatalf("got wrong data for first row: %d", department[0].DepartmentId)
	}

}
