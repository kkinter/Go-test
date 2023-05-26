package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_authenticate(t *testing.T) {
	var tests = []struct {
		name           string
		requestBody    string
		wantStatusCode int
	}{
		{"valid user", `{"email":"admin@example.com", "password":"secret"}`, http.StatusOK},
		{"not json", `not json`, http.StatusUnauthorized},
		{"empty json", `{}`, http.StatusUnauthorized},
		{"empty email", `{"email":"", "password":"secret"}`, http.StatusUnauthorized},
		{"empty password", `{"email":"admin@example.com", "password":""}`, http.StatusUnauthorized},
		{"invalid user", `{"email":"admin@some.com", "password":"secret"}`, http.StatusUnauthorized},
	}

	for _, e := range tests {
		var reader io.Reader = strings.NewReader(e.requestBody)

		req, _ := http.NewRequest("POST", "/auth", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.authenticate)

		handler.ServeHTTP(rr, req)

		if e.wantStatusCode != rr.Code {
			t.Errorf("%s: returned wrong status code want %d but got %d", e.name, e.wantStatusCode, rr.Code)
		}
	}
}
