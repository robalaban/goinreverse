package heartbeat

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

//TODO: After decoupling the logic from heartbeat package, testing multiple routes is made simple

func externalServerStub() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/healthCheck":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}))
}

func TestPingServer(t *testing.T) {
	server := externalServerStub()
	defer server.Close()

	healthCheckURL, _ := url.Parse(server.URL)
	status := PingServer(healthCheckURL)
	if status != true {
		t.Errorf("Got false when expected true, 200 != 200")
	}
}
