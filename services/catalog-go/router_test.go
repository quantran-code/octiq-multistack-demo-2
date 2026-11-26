package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
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
		{"/info", http.StatusOK, `"service":"catalog-go"`},
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

func TestInfo(t *testing.T) {
	rec := httptest.NewRecorder()
	newRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/info", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("/info: status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got ServiceInfo
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("/info: decode body: %v", err)
	}
	want := ServiceInfo{
		Service:   "catalog-go",
		Runtime:   "go/net-http",
		Endpoints: []string{"/health", "/products", "/products/{id}", "/info"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("/info: body = %+v, want %+v", got, want)
	}
}
