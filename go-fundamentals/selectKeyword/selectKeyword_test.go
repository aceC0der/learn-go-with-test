package selectKeyword

import (
	"learn-go-with-tests/message-handler"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speed of servers, returning the url of the fastest one", func(t *testing.T){
		slowServer := makeDelayedServer(20 * time.Millisecond)
		slowURL := slowServer.URL
		defer slowServer.Close()
	
		fastServer := makeDelayedServer(0 * time.Millisecond)
		fastURL := fastServer.URL
		defer fastServer.Close()
	
		want := fastURL
		got, _ := Racer(slowURL, fastURL)
	
		if got != want {
			message.PrintError(t, &got, &want);
		}
	})

	t.Run("returns an error if a server doesn't respond within 20 seconds", func(t *testing.T){
		serverA := makeDelayedServer(25 * time.Second)
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 20 * time.Second)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}