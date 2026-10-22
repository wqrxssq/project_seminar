package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	rec := httptest.NewRecorder()
	NewMux("passport").ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/ping", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Fatalf("content-type %q", ct)
	}

	var got map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["service"] != "passport" || got["status"] != "ok" || len(got) != 2 {
		t.Fatalf("body %s", rec.Body.Bytes())
	}
}

func TestNoOtherRoutes(t *testing.T) {
	h := NewMux("media")
	tests := []struct {
		method string
		path   string
		want   int
	}{
		{http.MethodPost, "/ping", http.StatusMethodNotAllowed},
		{http.MethodPut, "/ping", http.StatusMethodNotAllowed},
		{http.MethodGet, "/v1/apartments", http.StatusNotFound},
		{http.MethodPost, "/v1/auth/sessions", http.StatusNotFound},
		{http.MethodGet, "/v1/stories", http.StatusNotFound},
	}
	for _, tt := range tests {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, nil))
		if rec.Code != tt.want {
			t.Errorf("%s %s: got %d want %d", tt.method, tt.path, rec.Code, tt.want)
		}
	}
}

func TestListenAddr(t *testing.T) {
	t.Setenv("LISTEN_ADDR", "")
	if ListenAddr() != ":8080" {
		t.Fatalf("default %q", ListenAddr())
	}
	t.Setenv("LISTEN_ADDR", "127.0.0.1:9090")
	if ListenAddr() != "127.0.0.1:9090" {
		t.Fatalf("from env %q", ListenAddr())
	}
}
