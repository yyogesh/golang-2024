package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Application_Handler(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/abc", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	// create a test server
	ts := httptest.NewServer(routes)
	defer ts.Close()

	for _, tt := range theTests {
		resp, err := ts.Client().Get(ts.URL + tt.url)

		if err != nil {
			t.Fatalf("%s: Failed to get response: %v", tt.name, err)
		}

		if resp.StatusCode != tt.expectedStatusCode {
			t.Errorf("%s: Expected status code %d, got %d", tt.name, tt.expectedStatusCode, resp.StatusCode) // compare the actual status code with the expected status code
		}
	}

}
