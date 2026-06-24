package main
import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(PingHandler)
	handler.ServeHTTP(rr, req)
	
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Incorrect status: получили %v, expected %v,", status, http.StatusOK)
	}
	if rr.Body.String() != "pong" {
		t.Errorf("Incorrect answer: получили %v, expected %v,", rr.Body.String(), "pong")
	}
}