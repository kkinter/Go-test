package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"wepapp/pkg/data"
)

func Test_app_enableCORS(t *testing.T) {
	nextHandeler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	var tests = []struct {
		name       string
		method     string
		wantHeader bool
	}{
		{"preflight", "OPTIONS", true},
		{"get", "GET", false},
	}

	for _, e := range tests {
		handlerToTest := app.enableCORS(nextHandeler)

		req, _ := http.NewRequest(e.method, "http://testing.com", nil)
		rr := httptest.NewRecorder()

		handlerToTest.ServeHTTP(rr, req)

		if e.wantHeader && rr.Header().Get("Access-Control-Allow-Credentials") == "" {
			t.Errorf("%s: want header, but empty", e.name)
		}

		if !e.wantHeader && rr.Header().Get("Access-Control-Allow-Credentials") != "" {
			t.Errorf("%s: want empty hedaer but got %s", e.name, rr.Header().Get("Access-Control-Allow-Credentials"))
		}
	}
}

func Test_app_authRequired(t *testing.T) {
	nextHandeler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}

	tokens, _ := app.generateTokenPair(&testUser)

	var tests = []struct {
		name           string
		token          string
		wantAuthorized bool
		setHeader      bool
	}{
		{name: "valid token", token: fmt.Sprintf("Bearer %s", tokens.Token), wantAuthorized: true, setHeader: true},
		{name: "no token", token: "", wantAuthorized: false, setHeader: false},
		{name: "invalid token", token: fmt.Sprintf("Bearer %s", expiredToken), wantAuthorized: false, setHeader: true},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		if e.setHeader {
			req.Header.Set("Authorization", e.token)
		}

		rr := httptest.NewRecorder()
		handelrToTest := app.authRequired(nextHandeler)
		handelrToTest.ServeHTTP(rr, req)

		if e.wantAuthorized && rr.Code == http.StatusUnauthorized {
			t.Errorf("%s: want authorization but got %d", e.name, rr.Code)
		}

		if !e.wantAuthorized && rr.Code != http.StatusUnauthorized {
			t.Errorf("%s: want unauthorization but got %d", e.name, rr.Code)
		}
	}
}
