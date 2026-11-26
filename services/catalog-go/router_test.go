package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	cases := []struct {
		path   string
		status int
		body   string
	}{
		{"/health", http.StatusOK, `"status":"ok"`},
		{"/products", http.StatusOK, `"name":"Backpack"`},
		{"/products/p-2", http.StatusOK, `"name":"Pen"`},
		{"/products/missing", http.StatusNotFound, "product not found"},
	}
	router := newRouter()
	for _, tc := range cases {
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tc.path, nil))
		if rec.Code != tc.status {
			t.Errorf("%s: status = %d, want %d", tc.path, rec.Code, tc.status)
		}
		if !strings.Contains(rec.Body.String(), tc.body) {
			t.Errorf("%s: body = %s, want it to contain %s", tc.path, rec.Body.String(), tc.body)
		}
	}
}
