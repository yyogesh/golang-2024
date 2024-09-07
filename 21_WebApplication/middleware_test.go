package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_addIpToContext(t *testing.T) {
	tests := []struct {
		headerName   string
		headerValue  string
		address      string
		emptyAddress bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	var app application

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)

		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}

		t.Log(ip)
	})

	for _, e := range tests {
		handleToTest := app.addIpToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddress {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.address) > 0 {
			req.RemoteAddr = e.address
		}

		handleToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}
