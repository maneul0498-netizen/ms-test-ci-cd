package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	s := Message("manuel")

	if s != "notification sent to manuel" {
		t.Fatal()
	}

}

func TestNotifyEndpoint(t *testing.T) {

	body := `{"name":"manuel"}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/notify",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(notify)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf(
			"expected status 200 but got %d",
			rr.Code,
		)
	}

	expected := "notification sent to manuel"

	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf(
			"expected body %s but got %s",
			expected,
			rr.Body.String(),
		)
	}
}
