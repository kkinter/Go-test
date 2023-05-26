package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	// create a dummy handler that we'll use to check the contex
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that the value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "존재하지 않습니다")
		}

		// make sure we got a string back
		ip, ok := val.(string)
		if !ok {
			t.Error("string 타입이 아닙니다.")
		}
		t.Log(ip)
	})

	for _, test := range tests {
		// create the handler to test
		handlerToTest := app.addIPToContext(nextHandler)

		// mock request
		req := httptest.NewRequest("GET", "http://testing", nil)

		if test.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(test.headerName) > 0 {
			req.Header.Add(test.headerName, test.headerValue)
		}

		if len(test.addr) > 0 {
			req.RemoteAddr = test.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {

	// get a context and put something in the context
	ctx := context.WithValue(context.Background(), contextUserKey, "whatever")

	// call the function
	ip := app.ipFromContext(ctx)

	// perfom the test
	if !strings.EqualFold("whatever", ip) {
		t.Error("잘못된 값이 컨텍스트로 반환되었습니다.")
	}
}
